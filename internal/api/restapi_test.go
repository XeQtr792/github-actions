package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_GetStuff(t *testing.T) {
	serverCalled := 0
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		serverCalled++
		writer.WriteHeader(200)
	}))
	defer testServer.Close()
	apii := new(Api)
	apii.Client = http.DefaultClient
	assert.NotPanicsf(t, func() {
		apii.GetStuff(testServer.URL)
	},"")
	assert.Equalf(t, serverCalled, 1, "")
}
