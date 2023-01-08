package goym

import (
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

func (s *TrackTestSuite) TestGetLikedDislikedTracks() {
	// liked
	data, err := s.cl.GetLikedDislikedTracks(true)
	s.require.Nil(err)
	s.require.NotEmpty(data.Result.Library.Uid)

	// disliked
	data, err = s.cl.GetLikedDislikedTracks(false)
	s.require.Nil(err)
	s.require.Positive(data.Result.Library.Uid)
}

func (s *TrackTestSuite) TestLikeUnlikeTrack() {
	// search & get track id
	data, err := s.cl.Search("llll falling alone", 0, SearchTypeTrack, false)
	s.require.Nil(err)

	var tracks = data.Result.Tracks
	s.require.NotNil(tracks)

	// :(
	s.require.NotEmpty(tracks.Results)

	// :)
	var trackID = tracks.Results[0].ID

	// like
	err = s.cl.LikeUnlikeTracks([]int64{trackID}, true)
	s.require.Nil(err)

	// unlike
	err = s.cl.LikeUnlikeTracks([]int64{trackID}, false)
	s.require.Nil(err)
}

func (s *TrackTestSuite) TestLikeUnlikeTracks() {
	// search & get tracks id
	data, err := s.cl.Search("mick gordon", 0, SearchTypeTrack, false)
	s.require.Nil(err)

	var tracks = data.Result.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)

	var ids = []int64{}
	for _, t := range tracks.Results {
		ids = append(ids, t.ID)
	}

	// like
	err = s.cl.LikeUnlikeTracks(ids, true)
	s.require.Nil(err)

	// unlike
	err = s.cl.LikeUnlikeTracks(ids, false)
	s.require.Nil(err)
}

func (s *TrackTestSuite) TestGetTrackById() {
	s.getTrackId()
}

// get random track id.
func (s *TrackTestSuite) getTrackId() int64 {
	// search & get track id
	data, err := s.cl.Search("привет с большого бодуна", 0, SearchTypeTrack, false)
	s.require.Nil(err)

	var tracks = data.Result.Tracks
	s.require.NotNil(tracks)
	s.require.NotEmpty(tracks.Results)

	return tracks.Results[0].ID
}

func (s *TrackTestSuite) TestGetTracksById() {
	liked, err := s.cl.GetLikedDislikedTracks(true)
	s.require.Nil(err)

	var trackIds = []int64{}
	for _, ts := range liked.Result.Library.Tracks {
		trackId, err := s2i64(ts.Id)
		s.require.Nil(err)
		trackIds = append(trackIds, trackId)
	}

	_, err = s.cl.GetTracksById(trackIds)
	s.require.Nil(err)
}

func (s *TrackTestSuite) TestGetTrackDownloadInfo() {
	var id = s.getTrackId()

	// get info
	respInfo, err := s.cl.GetTrackDownloadInfo(id)
	s.require.Nil(err)
	s.require.NotEmpty(respInfo.Result)
	s.require.Positive(respInfo.Result[0].BitrateInKbps)
}

func (s *TrackTestSuite) TestGetTrackSupplement() {
	var id = s.getTrackId()

	// get info
	resp, err := s.cl.GetTrackSupplement(id)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result.Id)
}

func (s *TrackTestSuite) TestGetSimilarTracks() {
	var id = s.getTrackId()

	// get info
	resp, err := s.cl.GetSimilarTracks(id)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result.SimilarTracks)
}
