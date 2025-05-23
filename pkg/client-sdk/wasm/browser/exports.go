//go:build js && wasm
// +build js,wasm

package browser

import (
	"context"
	"fmt"
	"syscall/js"

	arksdk "github.com/ark-network/ark/pkg/client-sdk"
	"github.com/ark-network/ark/pkg/client-sdk/types"
	"github.com/ark-network/ark/pkg/client-sdk/wallet"
	singlekeywallet "github.com/ark-network/ark/pkg/client-sdk/wallet/singlekey"
	walletstore "github.com/ark-network/ark/pkg/client-sdk/wallet/singlekey/store"
)

var (
	arkSdkClient arksdk.ArkClient
	store        types.Store
	version      string
)

func init() {
	js.Global().Set("init", InitWrapper())
	js.Global().Set("unlock", UnlockWrapper())
	js.Global().Set("lock", LockWrapper())
	js.Global().Set("locked", IsLockedWrapper())
	js.Global().Set("balance", BalanceWrapper())
	js.Global().Set("receive", ReceiveWrapper())
	js.Global().Set("sendOffChain", SendOffChainWrapper())
	js.Global().Set("settle", SettleWrapper())
	js.Global().Set("collaborativeExit", CollaborativeExitWrapper())
	js.Global().Set("startUnilateralExit", StartUnilateralExitWrapper())
	js.Global().Set("completeUnilateralExit", CompleteUnilateralExitWrapper())
	js.Global().Set("onboardAgainAllExpiredBoardings", OnboardAgainAllExpiredBoardingsWrapper())
	js.Global().Set("withdrawFromAllExpiredBoardings", WithdrawFromAllExpiredBoardingsWrapper())
	js.Global().Set("getTransactionHistory", GetTransactionHistoryWrapper())
	js.Global().Set("log", LogWrapper())
	js.Global().Set("dump", DumpWrapper())
	js.Global().Set("redeemNotes", RedeemNotesWrapper())
	js.Global().Set("recoverSweptVtxos", RecoverSweptVtxosWrapper())
	js.Global().Set("listVtxos", ListVtxosWrapper())
	js.Global().Set("signTransaction", SignTransactionWrapper())
	js.Global().Set("notifyIncomingFunds", NotifyIncomingFundsWrapper())
	js.Global().Set("reset", ResetWrapper())

	js.Global().Set("getServerUrl", GetServerUrlWrapper())
	js.Global().Set("getServerPubkeyHex", GetServerPubkeyWrapper())
	js.Global().Set("getWalletType", GetWalletTypeWrapper())
	js.Global().Set("getClientType", GetClientTypeWrapper())
	js.Global().Set("getNetwork", GetNetworkWrapper())
	js.Global().Set("getVtxoTreeExpiry", GetVtxoTreeExpiryWrapper())
	js.Global().Set("getUnilateralExitDelay", GetUnilateralExitDelayWrapper())
	js.Global().Set("getDust", GetDustWrapper())
	js.Global().Set("getVersion", GetVersionWrapper())
}

func NewArkClient(
	ctx context.Context, storeSvc types.Store, v string,
) error {
	var err error

	data, err := storeSvc.ConfigStore().GetData(ctx)
	if err != nil {
		return err
	}

	if data == nil {
		arkSdkClient, err = arksdk.NewArkClient(storeSvc)
	} else {
		var walletSvc wallet.WalletService
		switch data.WalletType {
		case arksdk.SingleKeyWallet:
			walletSvc, err = getSingleKeyWallet(storeSvc.ConfigStore())
			if err != nil {
				return err
			}
		// TODO: Support HD wallet
		default:
			return fmt.Errorf("unknown wallet type")
		}
		arkSdkClient, err = arksdk.LoadArkClientWithWallet(storeSvc, walletSvc)
	}
	if err != nil {
		js.Global().Get("console").Call("error", err.Error())
		return err
	}
	store = storeSvc
	version = v

	select {}
}

func getWalletStore(storeType string) (walletstore.WalletStore, error) {
	if storeType == LocalStorageStore {
		return NewLocalStorageWalletStore()
	}
	// TODO: Support IndexDB store
	return nil, fmt.Errorf("unknown wallet store type")
}

func getSingleKeyWallet(configStore types.ConfigStore) (wallet.WalletService, error) {
	walletStore, err := getWalletStore(configStore.GetType())
	if err != nil {
		return nil, err
	}
	return singlekeywallet.NewBitcoinWallet(configStore, walletStore)
}
