package files

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardoraider/go-box/internal/bucket"
	"github.com/eduardoraider/go-box/internal/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type TransactionSuite struct {
	suite.Suite
	conn    *sql.DB
	mock    sqlmock.Sqlmock
	handler handler
	entity  *File
}

func (ts *TransactionSuite) SetupTest() {
	var err error

	ts.conn, ts.mock, err = sqlmock.New()
	assert.NoError(ts.T(), err)

	b, err := bucket.New(bucket.MockProvider, nil)
	assert.NoError(ts.T(), err)

	q, err := queue.New(queue.Mock, nil)
	assert.NoError(ts.T(), err)

	ts.handler = handler{ts.conn, b, q}

	ts.entity = &File{
		OwnerId: 1,
		Name:    "test_gopher.png",
		Type:    "application/octet-stream",
		Path:    "/test_gopher.png",
	}
}

func (ts *TransactionSuite) TearDownTest(_, _ string) {
	assert.NoError(ts.T(), ts.mock.ExpectationsWereMet())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
