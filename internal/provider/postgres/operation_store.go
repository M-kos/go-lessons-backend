package postgres

import (
	"context"
	_ "embed"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

var (
	//go:embed queries/create_operation.sql
	createOperationQuery string
	//go:embed queries/get_balance_by_user_id.sql
	getBalanceByUserIdQuery string
	//go:embed queries/get_operations_by_user_id.sql
	getOperationsByUserIdQuery string
	//go:embed queries/get_balance_snapshot_by_user_id.sql
	getBalanceSnapshotByUserIdQuery string
	//go:embed queries/get_max_sequence_number_by_user_id.sql
	getMaxSequenceNumberByUserId string
)

type OperationStore struct {
	db *pgxpool.Pool
}

func NewOperationStore(db *pgxpool.Pool) *OperationStore {
	return &OperationStore{
		db: db,
	}
}

func (s *OperationStore) ListOperationsByUserID(ctx context.Context, userID int64) ([]domain.Operation, error) {
	rows, err := s.db.Query(ctx, getOperationsByUserIdQuery, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	operations, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Operation])
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func (s *OperationStore) GetUserBalance(ctx context.Context, userID int64) (decimal.Decimal, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx) //TODO Обсудить, как принято обрабатывать в таких местах
			return
		}
		_ = tx.Commit(ctx)
	}()

	row := tx.QueryRow(ctx, getBalanceSnapshotByUserIdQuery, userID)

	var snapshot domain.BalanceSnapshot

	if err = row.Scan(&snapshot.ID, &snapshot.UserID, &snapshot.SequenceNumber, &snapshot.Balance); err != nil {
		return decimal.Zero, err
	}

	row = tx.QueryRow(ctx, getBalanceByUserIdQuery, userID, snapshot.SequenceNumber)

	var balance decimal.Decimal

	if err = row.Scan(&balance); err != nil {
		return decimal.Zero, err
	}

	balanceFromSnapshot := snapshot.Balance
	balance = balance.Add(balanceFromSnapshot)

	return balance, nil
}

func (s *OperationStore) CreateTransfer(
	ctx context.Context,
	sourceUserID int64,
	targetUserID int64,
	amount decimal.Decimal,
) (uuid.UUID, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	row := tx.QueryRow(ctx, getMaxSequenceNumberByUserId, sourceUserID)

	var sourceUserSeqNumber int64

	if err = row.Scan(&sourceUserSeqNumber); err != nil {
		return uuid.Nil, err
	}

	row = tx.QueryRow(ctx, getMaxSequenceNumberByUserId, targetUserID)

	var targetUserSeqNumber int64

	if err = row.Scan(&targetUserSeqNumber); err != nil {
		return uuid.Nil, err
	}

	_, err = tx.Exec(ctx, createOperationQuery, uuid.New(), sourceUserID, sourceUserSeqNumber, domain.OperationTypeCredit, amount.InexactFloat64()) //TODO Как с этим работать
	if err != nil {
		return uuid.Nil, err
	}

	_, err = tx.Exec(ctx, createOperationQuery, uuid.New(), targetUserID, targetUserSeqNumber, domain.OperationTypeDebit, amount.InexactFloat64()) //TODO Как с этим работать
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.Nil, nil // Не понял, чей UUID возвращать
}
