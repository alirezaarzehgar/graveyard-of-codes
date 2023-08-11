package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHi(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(http.MethodGet, "/hi", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Hi(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"OK!\"\n", rec.Body.String())
	}
}
