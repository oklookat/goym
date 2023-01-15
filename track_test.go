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

// get random track.
func (s TrackTestSuite) getTrack() *schema.Track {
	found, err := s.cl.Search(context.Background(), "привет с большого бодуна", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	var tracks = found.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)
	var track = tracks.Results[0]
	s.require.Positive(track.ID)
	return tracks.Results[0]
}

// get random tracks.
func (s TrackTestSuite) getTracks() []*schema.Track {
	found, err := s.cl.Search(context.Background(), "mick gordon", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	var tracks = found.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)

	var tracksData = []*schema.Track{}
	for i, t := range tracks.Results {
		tracksData = append(tracksData, t)
		if i == 5 {
			break
		}
	}
	return tracksData
}

func (s *TrackTestSuite) TestGetLikedTracks() {
	tracks, err := s.cl.GetLikedTracks(context.Background())
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
	s.require.NotEmpty(tracks.Library.Tracks)
}

func (s TrackTestSuite) TestGetDislikedTracks() {
	tracks, err := s.cl.GetDislikedTracks(context.Background())
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
}

func (s TrackTestSuite) TestLikeUnlikeTrack() {
	var ctx = context.Background()
	var track = s.getTrack()

	// like
	err := s.cl.LikeTrack(ctx, track)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks(ctx, []*schema.Track{track})
	s.require.Nil(err)
}

func (s TrackTestSuite) TestLikeUnlikeTracks() {
	var ctx = context.Background()
	var tracks = s.getTracks()

	// like
	err := s.cl.LikeTracks(ctx, tracks)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks(ctx, tracks)
	s.require.Nil(err)
}

func (s TrackTestSuite) TestGetTrackById() {
	var track = s.getTrack()
	tracks, err := s.cl.GetTrackById(context.Background(), track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
	s.require.Equal(tracks[0].ID, track.ID)
}

func (s TrackTestSuite) TestGetTracksById() {
	var ids = []schema.UniqueID{}
	var searched = s.getTracks()
	for _, t := range searched {
		ids = append(ids, t.ID)
	}

	tracks, err := s.cl.GetTracksByIds(context.Background(), ids)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
}

func (s TrackTestSuite) TestGetTrackDownloadInfo() {
	var track = s.getTrack()

	// get info
	respInfo, err := s.cl.GetTrackDownloadInfo(context.Background(), track)
	s.require.Nil(err)
	s.require.NotEmpty(respInfo)
	s.require.Positive(respInfo[0].BitrateInKbps)
}

func (s TrackTestSuite) TestGetTrackSupplement() {
	var track = s.getTrack()

	// get info
	resp, err := s.cl.GetTrackSupplement(context.Background(), track)
	s.require.Nil(err)
	s.require.NotEmpty(resp.ID)
}

func (s TrackTestSuite) TestGetSimilarTracks() {
	var track = s.getTrack()
	resp, err := s.cl.GetSimilarTracks(context.Background(), track)
	s.require.Nil(err)
	s.require.NotEmpty(resp.SimilarTracks)
}
