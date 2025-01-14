package transaction

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"pms.pkg/errs"
	"pms.pkg/logger"
)

func TestMain(m *testing.M) {
	setupLogger()
	code := m.Run()
	os.Exit(code)
}

func setupLogger() {
	conf := logger.Config{
		Dev:  true,
		Path: "",
	}
	logger.Init(conf)
}

func TestStart_ReuseTransaction(t *testing.T) {
	ctx := context.Background()

	// Mock DB and expectations
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Add a mock transaction to context
	mockTx := &sqlx.Tx{}
	ctx = context.WithValue(ctx, TX_KEY, mockTx)

	// Start should reuse the transaction
	newCtx, err := StartCTX(ctx, sqlxDB)
	assert.NoError(t, err)
	assert.Equal(t, ctx, newCtx)
}

func TestStart_NewTransaction(t *testing.T) {
	ctx := context.Background()

	// Mock DB and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Expect a new transaction to be created
	mock.ExpectBegin()

	// Start should create a new transaction
	newCtx, err := StartCTX(ctx, sqlxDB)
	assert.NoError(t, err)

	// Verify transaction in context
	retrievedTx, err := RetrieveCTX(newCtx)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedTx)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEnd_CommitTransaction(t *testing.T) {
	ctx := context.Background()

	// Mock DB and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Expect a new transaction and commit
	mock.ExpectBegin()
	mock.ExpectCommit()

	// Start a transaction
	ctx, err = StartCTX(ctx, sqlxDB)
	assert.NoError(t, err)

	// End the transaction successfully
	EndCTX(ctx, nil)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEnd_RollbackTransaction(t *testing.T) {
	ctx := context.Background()

	// Mock DB and expectations
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Expect a new transaction and rollback
	mock.ExpectBegin()
	mock.ExpectRollback()

	// Start a transaction
	ctx, err = StartCTX(ctx, sqlxDB)
	assert.NoError(t, err)

	// End the transaction with an error
	EndCTX(ctx, errs.ErrInternal{Reason: "some error"})

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEnd_NoTransaction(t *testing.T) {
	ctx := context.Background()

	// Test: End should handle missing transaction gracefully
	EndCTX(ctx, nil) // No panic or crash
}
