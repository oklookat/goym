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
	tracks := found.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)
	track := tracks.Results[0]
	s.require.Positive(track.ID)
	return tracks.Results[0]
}

// get random tracks.
func (s TrackTestSuite) getTracks() []*schema.Track {
	found, err := s.cl.Search(context.Background(), "mick gordon", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	tracks := found.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)

	tracksData := []*schema.Track{}
	for i, t := range tracks.Results {
		tracksData = append(tracksData, t)
		if i == 5 {
			break
		}
	}
	return tracksData
}

func (s *TrackTestSuite) TestLikedTracks() {
	tracks, err := s.cl.LikedTracks(context.Background())
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
	s.require.NotEmpty(tracks.Library.Tracks)
}

func (s TrackTestSuite) TestDislikedTracks() {
	tracks, err := s.cl.DislikedTracks(context.Background())
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
}

func (s TrackTestSuite) TestLikeUnlikeTrack() {
	ctx := context.Background()
	track := s.getTrack()

	// like
	err := s.cl.LikeTrack(ctx, track.ID)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks(ctx, []schema.ID{track.ID})
	s.require.Nil(err)
}

func (s TrackTestSuite) TestLikeUnlikeTracks() {
	ctx := context.Background()
	tracks := s.getTracks()

	// like
	var ids []schema.ID
	for i := range tracks {
		ids = append(ids, tracks[i].ID)
	}
	err := s.cl.LikeTracks(ctx, ids)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks(ctx, ids)
	s.require.Nil(err)
}

func (s TrackTestSuite) TestGetTrackById() {
	track := s.getTrack()
	tracks, err := s.cl.Track(context.Background(), track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
	s.require.Equal(tracks[0].ID, track.ID)
}

func (s TrackTestSuite) TestTracksById() {
	ids := []schema.ID{}
	searched := s.getTracks()
	for _, t := range searched {
		ids = append(ids, t.ID)
	}

	tracks, err := s.cl.Tracks(context.Background(), ids)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
}

func (s TrackTestSuite) TestTrackDownloadInfo() {
	track := s.getTrack()

	// get info
	respInfo, err := s.cl.TrackDownloadInfo(context.Background(), track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(respInfo)
	s.require.Positive(respInfo[0].BitrateInKbps)
}

func (s TrackTestSuite) TestTrackSupplement() {
	track := s.getTrack()

	// get info
	resp, err := s.cl.TrackSupplement(context.Background(), track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(resp.ID)
}

func (s TrackTestSuite) TestSimilarTracks() {
	track := s.getTrack()
	resp, err := s.cl.SimilarTracks(context.Background(), track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(resp.SimilarTracks)
}
