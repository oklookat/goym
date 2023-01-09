package goym

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PlaylistTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *PlaylistTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *PlaylistTestSuite) TestGetUserPlaylists() {
	_, err := s.cl.GetUserPlaylists(s.cl.UserId)
	s.require.Nil(err)
}

// GetUserPlaylist()
// CreatePlaylist()
// RenamePlaylist()
// DeletePlaylist()
// ChangePlaylistVisibility()
// AddTracksToPlaylist()
// GetPlaylistRecommendations()
func (s *PlaylistTestSuite) TestPlaylistCRUD() {
	// CreatePlaylist
	resp, err := s.cl.CreatePlaylist("goymtesting", false)
	s.require.Nil(err)
	var pl = resp.Result

	// GetUserPlaylist
	same, err := s.cl.GetUserPlaylist(s.cl.UserId, pl.Kind)
	s.require.Nil(err)
	s.require.Equal(pl.Kind, same.Result.Kind)

	// RenamePlaylist
	resp, err = s.cl.RenamePlaylist(pl.Kind, "goymtesting (renamed)")
	var renamed = resp.Result
	s.require.Nil(err)
	s.require.Equal(pl.Kind, renamed.Kind)

	// ChangePlaylistVisibility
	resp, err = s.cl.ChangePlaylistVisibility(renamed.Kind, true)
	s.require.Nil(err)
	var changed = resp.Result
	s.require.Equal(pl.Kind, changed.Kind)

	// AddPlaylistTracks (add)
	// tracksResp, err := s.cl.Search("dubstep", 0, SearchTypeTrack, false)
	// s.require.Nil(err)
	// s.require.NotNil(tracksResp.Result.Tracks)
	// s.require.NotEmpty(tracksResp.Result.Tracks.Results)
	// var tracks = tracksResp.Result.Tracks.Results
	// var tracksLittle = []*Track{}
	// for i := range tracks {
	// 	tracksLittle = append(tracksLittle, tracks[i])
	// 	if len(tracksLittle) == 10 {
	// 		break
	// 	}
	// }
	// resp, err = s.cl.AddPlaylistTracks(pl, tracksLittle)
	// s.require.Nil(err)
	// var changed2 = resp.Result
	// s.require.Equal(changed.Kind, changed2.Kind)

	// GetPlaylistRecommendations
	// recsResp, err := s.cl.GetPlaylistRecommendations(changed2.Kind)
	// s.require.Nil(err)
	// s.require.NotEmpty(recsResp.Result.Tracks)
	// tracks = recsResp.Result.Tracks
	// s.require.NotEmpty(tracks[0].Title)

	// // RemovePlaylistTracks (remove)
	// resp, err = s.cl.RemovePlaylistTracks(changed2, 1, 8)
	// s.require.Nil(err)
	// var changed3 = resp.Result
	// s.require.Equal(changed2.Kind, changed3.Kind)

	// DeletePlaylist
	err = s.cl.DeletePlaylist(changed.Kind)
	s.require.Nil(err)
}
