package goym

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DebugTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *DebugTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}
