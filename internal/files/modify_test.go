package files

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
)

func (ts *TransactionSuite) TestModify() {
	f := File{
		ID:   1,
		Name: "Gopher.png",
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&f)
	assert.NoError(ts.T(), err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/{id}", &b)

	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	setMockGet(ts.mock)
	setMockUpdate(ts.mock)

	ts.handler.Modify(rr, req)
	assert.Equal(ts.T(), http.StatusOK, rr.Code)

}

func (ts *TransactionSuite) TestUpdate() {
	setMockUpdate(ts.mock)

	err := Update(ts.conn, 1, &File{Name: "Gopher.png"})
	assert.NoError(ts.T(), err)
}

func setMockUpdate(mock sqlmock.Sqlmock) {
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE files SET name=$1, modified_at=$2, deleted=$3 WHERE id=$4`)).
		WithArgs("Gopher.png", AnyTime{}, false, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
}
