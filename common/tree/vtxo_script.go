package tree

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/ark-network/ark/common"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

var ErrNoExitLeaf = fmt.Errorf("no exit leaf")

type VtxoScript common.VtxoScript[bitcoinTapTree, Closure]

func ParseVtxoScript(scripts []string) (VtxoScript, error) {
	if len(scripts) == 0 {
		return nil, fmt.Errorf("empty tapscripts array")
	}

	types := []VtxoScript{
		&TapscriptsVtxoScript{},
	}

	for _, v := range types {
		if err := v.Decode(scripts); err == nil {
			return v, nil
		}
	}

	return nil, fmt.Errorf("invalid vtxo scripts: %s", scripts)
}

type TapscriptsVtxoScript struct {
	Closures []Closure
}

func (v *TapscriptsVtxoScript) Encode() ([]string, error) {
	encoded := make([]string, 0)
	for _, closure := range v.Closures {
		script, err := closure.Script()
		if err != nil {
			return nil, err
		}
		encoded = append(encoded, hex.EncodeToString(script))
	}
	return encoded, nil
}

func (v *TapscriptsVtxoScript) Decode(scripts []string) error {
	if len(scripts) == 0 {
		return fmt.Errorf("empty scripts array")
	}

	v.Closures = make([]Closure, 0, len(scripts))
	for _, script := range scripts {
		scriptBytes, err := hex.DecodeString(script)
		if err != nil {
			return err
		}

		closure, err := DecodeClosure(scriptBytes)
		if err != nil {
			return err
		}
		v.Closures = append(v.Closures, closure)
	}

	if len(v.Closures) == 0 {
		return fmt.Errorf("no valid closures found in scripts")
	}

	return nil
}

func (v *TapscriptsVtxoScript) Validate(server *secp256k1.PublicKey, minLocktime common.RelativeLocktime) error {
	serverXonly := schnorr.SerializePubKey(server)
	for _, forfeit := range v.ForfeitClosures() {
		keys := make([]*secp256k1.PublicKey, 0)
		switch c := forfeit.(type) {
		case *MultisigClosure:
			keys = c.PubKeys
		case *CLTVMultisigClosure:
			keys = c.PubKeys
		case *ConditionMultisigClosure:
			keys = c.PubKeys
		}

		if len(keys) == 0 {
			return fmt.Errorf("invalid forfeit closure, expected MultisigClosure, CLTVMultisigClosure or ConditionMultisigClosure")
		}

		// must contain server pubkey
		found := false
		for _, pubkey := range keys {
			if bytes.Equal(schnorr.SerializePubKey(pubkey), serverXonly) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid forfeit closure, server pubkey not found")
		}
	}

	smallestExit, err := v.SmallestExitDelay()
	if err != nil {
		if err == ErrNoExitLeaf {
			return nil
		}
		return err
	}

	if smallestExit.LessThan(minLocktime) {
		return fmt.Errorf("exit delay is too short")
	}

	return nil
}

func (v *TapscriptsVtxoScript) SmallestExitDelay() (*common.RelativeLocktime, error) {
	var smallest *common.RelativeLocktime

	for _, closure := range v.Closures {
		if csvClosure, ok := closure.(*CSVMultisigClosure); ok {
			if smallest == nil || csvClosure.Locktime.LessThan(*smallest) {
				smallest = &csvClosure.Locktime
			}
		}
	}

	if smallest == nil {
		return nil, ErrNoExitLeaf
	}

	return smallest, nil
}

func (v *TapscriptsVtxoScript) ForfeitClosures() []Closure {
	forfeits := make([]Closure, 0)
	for _, closure := range v.Closures {
		switch closure.(type) {
		case *MultisigClosure, *CLTVMultisigClosure, *ConditionMultisigClosure:
			forfeits = append(forfeits, closure)
		}
	}
	return forfeits
}

func (v *TapscriptsVtxoScript) ExitClosures() []Closure {
	exits := make([]Closure, 0)
	for _, closure := range v.Closures {
		switch closure.(type) {
		case *CSVMultisigClosure:
			exits = append(exits, closure)
		}
	}
	return exits
}

func NewDefaultVtxoScript(owner, server *secp256k1.PublicKey, exitDelay common.RelativeLocktime) *TapscriptsVtxoScript {
	return &TapscriptsVtxoScript{
		[]Closure{
			&CSVMultisigClosure{
				MultisigClosure: MultisigClosure{PubKeys: []*secp256k1.PublicKey{owner}},
				Locktime:        exitDelay,
			},
			&MultisigClosure{PubKeys: []*secp256k1.PublicKey{owner, server}},
		},
	}
}

func (v *TapscriptsVtxoScript) TapTree() (*secp256k1.PublicKey, bitcoinTapTree, error) {
	leaves := make([]txscript.TapLeaf, len(v.Closures))
	for i, closure := range v.Closures {
		script, err := closure.Script()
		if err != nil {
			return nil, bitcoinTapTree{}, fmt.Errorf("failed to get script for closure %d: %w", i, err)
		}
		leaves[i] = txscript.NewBaseTapLeaf(script)
	}

	tapTree := txscript.AssembleTaprootScriptTree(leaves...)
	root := tapTree.RootNode.TapHash()
	taprootKey := txscript.ComputeTaprootOutputKey(
		UnspendableKey(),
		root[:],
	)

	return taprootKey, bitcoinTapTree{tapTree}, nil
}

// bitcoinTapTree is a wrapper around txscript.IndexedTapScriptTree to implement the common.TaprootTree interface
type bitcoinTapTree struct {
	*txscript.IndexedTapScriptTree
}

func (b bitcoinTapTree) GetRoot() chainhash.Hash {
	return b.RootNode.TapHash()
}

func (b bitcoinTapTree) GetTaprootMerkleProof(leafhash chainhash.Hash) (*common.TaprootMerkleProof, error) {
	index, ok := b.LeafProofIndex[leafhash]
	if !ok {
		return nil, fmt.Errorf("leaf %s not found in tree", leafhash.String())
	}
	proof := b.LeafMerkleProofs[index]

	controlBlock := proof.ToControlBlock(UnspendableKey())
	controlBlockBytes, err := controlBlock.ToBytes()
	if err != nil {
		return nil, err
	}

	return &common.TaprootMerkleProof{
		ControlBlock: controlBlockBytes,
		Script:       proof.Script,
	}, nil
}

func (b bitcoinTapTree) GetLeaves() []chainhash.Hash {
	leafHashes := make([]chainhash.Hash, 0)
	for hash := range b.LeafProofIndex {
		leafHashes = append(leafHashes, hash)
	}
	return leafHashes
}
