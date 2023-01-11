package goym

import (
	"context"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *AccountTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s AccountTestSuite) TestGetAccountStatus() {
	stat, err := s.cl.GetAccountStatus(context.Background())
	s.require.Nil(err)
	s.require.NotNil(stat.Account)
	s.require.Positive(stat.Account.UID)
	s.require.NotEmpty(stat.Account.Login)
}
