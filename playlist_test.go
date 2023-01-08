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
	// get user playlists
	data, err := s.cl.GetUserPlaylists(s.cl.UserId)
	s.require.Nil(err)
	s.require.NotEmpty(data.Result)

	// возможно, это плейлист "Мне нравится"
	var first = data.Result[0]
	s.require.NotEmpty(first.Kind)
}

// GetUserPlaylist()
// CreatePlaylist()
// RenamePlaylist()
// DeletePlaylist()
// ChangePlaylistVisibility()
func (s *PlaylistTestSuite) TestPlaylistCRUD() {
	// create
	resp, err := s.cl.CreatePlaylist("goymtesting", false)
	s.require.Nil(err)
	var pl = resp.Result

	// read
	same, err := s.cl.GetUserPlaylist(s.cl.UserId, pl.Kind)
	s.require.Nil(err)
	s.require.Equal(pl.Kind, same.Result.Kind)

	// update
	resp, err = s.cl.RenamePlaylist(pl.Kind, "goymtesting (renamed)")
	var renamed = resp.Result
	s.require.Nil(err)
	s.require.Equal(pl.Kind, renamed.Kind)

	// update 2
	resp, err = s.cl.ChangePlaylistVisibility(renamed.Kind, true)
	s.require.Nil(err)
	var changed = resp.Result
	s.require.Equal(pl.Kind, changed.Kind)

	// delete
	err = s.cl.DeletePlaylist(changed.Kind)
	s.require.Nil(err)
}

func (s *PlaylistTestSuite) TestGetPlaylistRecommendations() {
	// get user playlists
	data, err := s.cl.GetUserPlaylists(s.cl.UserId)
	s.require.Nil(err)
	s.require.NotEmpty(data.Result)

	// возможно, это плейлист "Мне нравится"
	var first = data.Result[0]
	s.require.NotEmpty(first.Kind)

	recsResp, err := s.cl.GetPlaylistRecommendations(first.Kind)
	s.require.Nil(err)
	s.require.NotEmpty(recsResp.Result.Tracks)
	var tracks = recsResp.Result.Tracks
	s.require.NotEmpty(tracks[0].Title)
}
