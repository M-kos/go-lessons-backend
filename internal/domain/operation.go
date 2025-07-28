package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Operation struct {
	ID             uuid.UUID
	UserID         int64
	SequenceNumber int64
	Type           OperationType
	Amount         decimal.Decimal
	CreateTime     time.Time
}

type OperationType string

const (
	OperationTypeDebit  OperationType = "debit"
	OperationTypeCredit OperationType = "credit"
)
