package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
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

func (s AccountTestSuite) TestGetChangeAccountSettings() {
	var getSettings = func() *schema.AccountSettings {
		sett, err := s.cl.GetAccountSettings(context.Background())
		s.require.Nil(err)
		s.require.NotNil(sett)
		s.require.Positive(sett.UID)
		return sett
	}
	var newSet = schema.AccountSettings{}
	newSet.VolumePercents = 10
	_, err := s.cl.ChangeAccountSettings(context.Background(), newSet)
	s.require.Nil(err)

	sett := getSettings()
	s.require.Equal(newSet.VolumePercents, sett.VolumePercents)

	newSet.VolumePercents = 33
	_, err = s.cl.ChangeAccountSettings(context.Background(), newSet)
	s.require.Nil(err)

	sett = getSettings()
	s.require.Equal(newSet.VolumePercents, sett.VolumePercents)
}
