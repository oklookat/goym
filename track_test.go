package goym

import (
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
	// search & get track id
	data, err := s.cl.Search("привет с большого бодуна", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	var tracks = data.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)
	return tracks.Results[0]
}

// get random tracks.
func (s TrackTestSuite) getTracks() []*schema.Track {
	found, err := s.cl.Search("mick gordon", 0, schema.SearchTypeTrack, false)
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
	tracks, err := s.cl.GetLikedTracks()
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
}

func (s TrackTestSuite) TestGetDislikedTracks() {
	tracks, err := s.cl.GetDislikedTracks()
	s.require.Nil(err)
	s.require.Positive(tracks.Library.Uid)
}

func (s TrackTestSuite) TestUnlikeLikeTrack() {
	var track = s.getTrack()

	// like
	err := s.cl.LikeTrack(track)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks([]*schema.Track{track})
	s.require.Nil(err)
}

func (s TrackTestSuite) TestLikeUnlikeTracks() {
	var tracks = s.getTracks()

	// like
	err := s.cl.LikeTracks(tracks)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeTracks(tracks)
	s.require.Nil(err)
}

func (s TrackTestSuite) TestGetTrackById() {
	var track = s.getTrack()
	tracks, err := s.cl.GetTrackById(track.ID)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
	s.require.Equal(tracks[0].ID, track.ID)
}

func (s TrackTestSuite) TestGetTracksById() {
	var ids = []int64{}
	var searched = s.getTracks()
	for _, t := range searched {
		ids = append(ids, t.ID)
	}

	tracks, err := s.cl.GetTracksByIds(ids)
	s.require.Nil(err)
	s.require.NotEmpty(tracks)
}

func (s TrackTestSuite) TestGetTrackDownloadInfo() {
	var track = s.getTrack()

	// get info
	respInfo, err := s.cl.GetTrackDownloadInfo(track)
	s.require.Nil(err)
	s.require.NotEmpty(respInfo)
	s.require.Positive(respInfo[0].BitrateInKbps)
}

func (s TrackTestSuite) TestGetTrackSupplement() {
	var track = s.getTrack()

	// get info
	resp, err := s.cl.GetTrackSupplement(track)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Id)
}

func (s TrackTestSuite) TestGetSimilarTracks() {
	var track = s.getTrack()

	// get info
	resp, err := s.cl.GetSimilarTracks(track)
	s.require.Nil(err)
	s.require.NotEmpty(resp.SimilarTracks)
}
