package ctx

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestStartTx(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	c := New(context.Background())

	// Expect Begin for the first transaction creation
	mock.ExpectBegin()
	tx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	// Subsequent StartTx calls should not trigger another Begin
	sameTx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.Equal(t, tx, sameTx)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEndTx_Commit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	c := New(context.Background())

	mock.ExpectBegin()
	tx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	mock.ExpectCommit()

	c.EndTx(nil) // No error, should commit
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEndTx_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	c := New(context.Background())

	mock.ExpectBegin()
	tx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	mock.ExpectRollback()

	c.EndTx(errors.New("rollback error")) // Error passed, should rollback
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCommit(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	c := New(context.Background())

	mock.ExpectBegin()
	tx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	mock.ExpectCommit()

	c.Commit()
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	c := New(context.Background())

	mock.ExpectBegin()
	tx, err := c.StartTx(sqlxDB)
	assert.NoError(t, err)
	assert.NotNil(t, tx)

	mock.ExpectRollback()

	c.Rollback()
	assert.NoError(t, mock.ExpectationsWereMet())
}
