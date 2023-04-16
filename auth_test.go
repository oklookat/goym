package goym

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/oklookat/goym/auth"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
	require *require.Assertions
}

func (s *AuthTestSuite) SetupSuite() {
	s.require = s.Require()
}

func (s *AuthTestSuite) TestAuth() {
	err := godotenv.Load()
	s.require.Nil(err)

	login := os.Getenv("LOGIN")

	acc, err := auth.New(context.Background(), login, func(url, code string) {
		println("go to " + url + " with code " + code)
	}, nil)
	s.require.Nil(err)

	_, err = New(acc)
	s.require.Nil(err)
}
