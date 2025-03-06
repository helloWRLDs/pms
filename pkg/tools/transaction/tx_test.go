package transaction

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// Test Start: New Transaction
func TestStart_NewTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin() // Expect transaction start

	ctx := context.Background()
	newCtx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	tx := Retrieve(newCtx)
	assert.NotNil(t, tx) // Ensure transaction is set in context
}

// Test Start: Transaction Already Exists
func TestStart_ExistingTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// ✅ Expect only ONE transaction start
	mock.ExpectBegin()

	ctx := context.Background()

	// First call: starts a new transaction
	ctx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	// ✅ Second call: should NOT start a new transaction
	newCtx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	// ✅ Ensure the returned context is the same
	assert.Equal(t, ctx, newCtx)

	// ✅ Ensure no unexpected Beginx() calls
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Test Start: Failure to Begin Transaction
func TestStart_FailToBeginTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin().WillReturnError(assert.AnError) // Simulate DB failure

	ctx := context.Background()
	newCtx, err := Start(ctx, sqlxDB)

	assert.Error(t, err)
	assert.Equal(t, ctx, newCtx) // Ensure ctx remains unchanged
}

// Test Retrieve: No Transaction
func TestRetrieve_NotFound(t *testing.T) {
	ctx := context.Background()
	tx := Retrieve(ctx)
	assert.Nil(t, tx)
}

// Test Retrieve: Transaction Exists
func TestRetrieve_Found(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	tx := Retrieve(ctx)
	assert.NotNil(t, tx)
}

// Test End: Commit Success
func TestEnd_CommitSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectCommit()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	End(ctx, nil) // Commit transaction

	assert.NoError(t, mock.ExpectationsWereMet()) // Ensure commit was called
}

// Test End: Rollback on Error
func TestEnd_RollbackOnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectRollback()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	End(ctx, assert.AnError) // Expect rollback

	assert.NoError(t, mock.ExpectationsWereMet())
}

// Test End: No Transaction Exists
func TestEnd_NoTransaction(t *testing.T) {
	ctx := context.Background()
	End(ctx, nil) // Should not panic
}

// Test Commit: Success
func TestCommit_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectCommit()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	err := Commit(ctx)
	assert.NoError(t, err)
}

// Test Commit: No Transaction
func TestCommit_NoTransaction(t *testing.T) {
	ctx := context.Background()
	err := Commit(ctx)
	assert.NoError(t, err) // Should return nil
}

// Test Rollback: Success
func TestRollback_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectRollback()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	err := Rollback(ctx)
	assert.NoError(t, err)
}

// Test Rollback: No Transaction
func TestRollback_NoTransaction(t *testing.T) {
	ctx := context.Background()
	err := Rollback(ctx)
	assert.NoError(t, err) // Should return nil
}
