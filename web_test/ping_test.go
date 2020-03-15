package web_test

import (
	"gin-study-case/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T)  {
	router := routes.Route()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET" ,"/ping",nil)
	router.ServeHTTP(w , req)
	assert.Equal(t , http.StatusOK , w.Code)
	assert.Equal(t,"pong" , w.Body.String())
}
