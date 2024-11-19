// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"database/sql"
)

type Entity struct {
	ID             int64
	NostrRecipient string
}

type EntityVtxo struct {
	EntityID int64
	VtxoTxid string
	VtxoVout int64
}

type EntityVw struct {
	ID             int64
	NostrRecipient string
	VtxoTxid       sql.NullString
	VtxoVout       sql.NullInt64
}

type MarketHour struct {
	ID            int64
	StartTime     int64
	Period        int64
	RoundInterval int64
	UpdatedAt     int64
}

type Note struct {
	ID int64
}

type Payment struct {
	ID      string
	RoundID string
}

type PaymentReceiverVw struct {
	PaymentID      sql.NullString
	Pubkey         sql.NullString
	OnchainAddress sql.NullString
	Amount         sql.NullInt64
}

type PaymentVtxoVw struct {
	Txid      sql.NullString
	Vout      sql.NullInt64
	Pubkey    sql.NullString
	Amount    sql.NullInt64
	PoolTx    sql.NullString
	SpentBy   sql.NullString
	Spent     sql.NullBool
	Redeemed  sql.NullBool
	Swept     sql.NullBool
	ExpireAt  sql.NullInt64
	CreatedAt sql.NullInt64
	PaymentID sql.NullString
	RedeemTx  sql.NullString
}

type Receiver struct {
	PaymentID      string
	Pubkey         sql.NullString
	OnchainAddress sql.NullString
	Amount         int64
}

type Round struct {
	ID                string
	StartingTimestamp int64
	EndingTimestamp   int64
	Ended             bool
	Failed            bool
	StageCode         int64
	Txid              string
	UnsignedTx        string
	ConnectorAddress  string
	DustAmount        int64
	Version           int64
	Swept             bool
}

type RoundPaymentVw struct {
	ID      sql.NullString
	RoundID sql.NullString
}

type RoundTxVw struct {
	ID         sql.NullInt64
	Tx         sql.NullString
	RoundID    sql.NullString
	Type       sql.NullString
	Position   sql.NullInt64
	Txid       sql.NullString
	TreeLevel  sql.NullInt64
	ParentTxid sql.NullString
	IsLeaf     sql.NullBool
}

type Tx struct {
	ID         int64
	Tx         string
	RoundID    string
	Type       string
	Position   int64
	Txid       sql.NullString
	TreeLevel  sql.NullInt64
	ParentTxid sql.NullString
	IsLeaf     sql.NullBool
}

type Vtxo struct {
	Txid      string
	Vout      int64
	Pubkey    string
	Amount    int64
	PoolTx    string
	SpentBy   string
	Spent     bool
	Redeemed  bool
	Swept     bool
	ExpireAt  int64
	CreatedAt int64
	PaymentID sql.NullString
	RedeemTx  sql.NullString
}
