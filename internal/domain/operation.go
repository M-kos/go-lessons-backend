package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Operation struct {
	ID             uuid.UUID       `db:"id"`
	UserID         int64           `db:"user_id"`
	SequenceNumber int64           `db:"sequence_number"`
	Type           OperationType   `db:"operation_type"`
	Amount         decimal.Decimal `db:"amount"`
	CreateTime     time.Time       `db:"create_time"`
}

type OperationType string

const (
	OperationTypeDebit  OperationType = "debit"
	OperationTypeCredit OperationType = "credit"
)

type BalanceSnapshot struct {
	ID             int64           `db:"id"`
	UserID         int64           `db:"user_id"`
	SequenceNumber int64           `db:"sequence_number"`
	Balance        decimal.Decimal `db:"balance"`
	UpdateTime     time.Time       `db:"update_time"`
}
