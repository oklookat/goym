package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TrackTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *TrackTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *TrackTestSuite) TestLikedTracks() {
	tracks, err := s.cl.LikedTracks(context.Background())
	s.require.Nil(err)
	s.require.NotNil(tracks.Result)
}

func (s *TrackTestSuite) TestDislikedTracks() {
	tracks, err := s.cl.DislikedTracks(context.Background())
	s.require.Nil(err)
	s.require.NotNil(tracks.Result)
}

func (s *TrackTestSuite) TestLikeUnlikeTrack() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeTrack(ctx, trackIds[0])
	s.require.Nil(err)

	// unlike
	_, err = s.cl.UnlikeTracks(ctx, []schema.ID{trackIds[0]})
	s.require.Nil(err)
}

func (s *TrackTestSuite) TestLikeUnlikeTracks() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeTracks(ctx, trackIds[:])
	s.require.Nil(err)

	// unlike
	_, err = s.cl.UnlikeTracks(ctx, trackIds[:])
	s.require.Nil(err)
}

func (s *TrackTestSuite) TestGetTrackById() {
	tracks, err := s.cl.Track(context.Background(), trackIds[0])
	s.require.Nil(err)
	s.require.NotEmpty(tracks.Result)
}

func (s *TrackTestSuite) TestTracksById() {
	tracks, err := s.cl.Tracks(context.Background(), trackIds[:])
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
}

func (s *TrackTestSuite) TestTrackDownloadInfo() {
	// get info
	respInfo, err := s.cl.TrackDownloadInfo(context.Background(), trackIds[0])
	s.require.Nil(err)
	s.require.NotEmpty(respInfo.Result)
}

func (s *TrackTestSuite) TestTrackSupplement() {
	// get info
	resp, err := s.cl.TrackSupplement(context.Background(), trackIds[0])
	s.require.Nil(err)
	s.require.NotNil(resp.Result)
}

func (s *TrackTestSuite) TestSimilarTracks() {
	_, err := s.cl.SimilarTracks(context.Background(), trackIds[0])
	s.require.Nil(err)
}
