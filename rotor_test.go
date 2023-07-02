package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RotorTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *RotorTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *RotorTestSuite) getStation() *schema.RotorStation {
	dash, err := s.cl.GetRotorDashboard(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(dash.Result.Stations)
	s.require.NotEmpty(dash.Result.Stations[0].Station.ID.Tag)
	return dash.Result.Stations[0].Station
}

func (s *RotorTestSuite) getTracks(st *schema.RotorStation) *schema.RotorStationTracks {
	res, err := s.cl.GetRotorStationTracks(context.Background(), st, nil)
	s.require.Nil(err)
	s.require.NotEmpty(res.Result.Sequence)
	s.require.NotNil(res.Result.Sequence[0].Track)
	return res.Result
}

func (s *RotorTestSuite) TestGetRotorDashboard() {
	s.getStation()
}

func (s *RotorTestSuite) TestGetRotorStationTracks() {
	s.getTracks(s.getStation())
}

func (s *RotorTestSuite) TestGetRotorStationInfo() {
	st := s.getStation()
	res, err := s.cl.GetRotorStationInfo(context.Background(), st)
	s.require.Nil(err)
	s.require.NotEmpty(res.Result)
	s.require.NotEmpty(res.Result[0].Station.ID.Tag)
}

func (s *RotorTestSuite) TestGetRotorAccountStatus() {
	res, err := s.cl.GetRotorAccountStatus(context.Background())
	s.require.Nil(err)
	s.require.Equal(res.Result.Account.UID, s.cl.UserId)
}

func (s *RotorTestSuite) TestGetRotorStationsList() {
	getWithLang := func(lang *string) {
		res, err := s.cl.GetRotorStationsList(context.Background(), lang)
		s.require.Nil(err)
		s.require.NotEmpty(res)
		s.require.NotEmpty(res.Result[0].Station.ID.Tag)
	}
	lang := "en"
	getWithLang(&lang)
	lang = "ru"
	getWithLang(&lang)
	getWithLang(nil)
}

func (s *RotorTestSuite) TestGetRotorStationsFeedback() {
	ctx := context.Background()
	stat := s.getStation()
	tracks := s.getTracks(stat)
	track := tracks.Sequence[0].Track

	res, err := s.cl.RotorStationFeedback(ctx, stat, schema.RotorStationFeedbackTypeTrackStarted, tracks, track, 12.5)
	s.require.Nil(err)
	s.require.Equal(res, "ok")
}
