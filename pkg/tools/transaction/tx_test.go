package transaction

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestStart_NewTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()

	ctx := context.Background()
	newCtx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	tx := Retrieve(newCtx)
	assert.NotNil(t, tx)
}

func TestStart_ExistingTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()

	ctx := context.Background()

	ctx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	newCtx, err := Start(ctx, sqlxDB)
	assert.NoError(t, err)

	assert.Equal(t, ctx, newCtx)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStart_FailToBeginTransaction(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin().WillReturnError(assert.AnError)

	ctx := context.Background()
	newCtx, err := Start(ctx, sqlxDB)

	assert.Error(t, err)
	assert.Equal(t, ctx, newCtx)
}

func TestRetrieve_NotFound(t *testing.T) {
	ctx := context.Background()
	tx := Retrieve(ctx)
	assert.Nil(t, tx)
}

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

func TestEnd_CommitSuccess(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectCommit()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	End(ctx, nil)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEnd_RollbackOnError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	mock.ExpectBegin()
	mock.ExpectRollback()

	ctx := context.Background()
	ctx, _ = Start(ctx, sqlxDB)

	End(ctx, assert.AnError)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEnd_NoTransaction(t *testing.T) {
	ctx := context.Background()
	End(ctx, nil)
}

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

func TestCommit_NoTransaction(t *testing.T) {
	ctx := context.Background()
	err := Commit(ctx)
	assert.NoError(t, err)
}

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

func TestRollback_NoTransaction(t *testing.T) {
	ctx := context.Background()
	err := Rollback(ctx)
	assert.NoError(t, err)
}
