package vantuz

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestAll(t *testing.T) {
	suite.Run(t, &VantuzTestSuite{})
}

type VantuzTestSuite struct {
	suite.Suite
	require *require.Assertions
}

func (s *VantuzTestSuite) SetupSuite() {
	s.require = s.Require()
}

func (s VantuzTestSuite) TestRequestRateLimit() {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Write([]byte("Hello."))
	}))
	defer server.Close()

	const expected = 5
	req := C().R()
	for i := 0; i < expected; i++ {
		_, err := req.Get(context.Background(), server.URL)
		s.require.Nil(err)
	}

	s.require.EqualValues(expected, requestCount)
}

func (s VantuzTestSuite) TestOneRequestManySend() {
	const key = "grant_type"
	const val = "device_code"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		gotVal := r.Form.Get(key)
		s.require.Equal(val, gotVal)
	}))
	defer server.Close()

	form := map[string]string{
		key: val,
	}
	req := C().R().
		SetFormUrlMap(form)
	for i := 0; i < 10; i++ {
		fmt.Printf("Request: %d\n", i)
		_, err := req.Post(context.Background(), server.URL)
		s.require.Nil(err)
	}
}
