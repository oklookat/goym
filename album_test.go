package goym

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AlbumTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *AlbumTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *AlbumTestSuite) TestGetAlbum() {
	data, err := s.cl.GetAlbum(231541, false)
	s.require.Nil(err)

	var title = data.Result.Title
	s.require.Equal("Club Bizarre", title)
}

func (s *AlbumTestSuite) TestGetAlbumWithTracks() {
	data, err := s.cl.GetAlbum(231541, true)
	s.require.Nil(err)

	var title = data.Result.Title
	s.require.Equal("Club Bizarre", title)

	s.require.NotEmpty(data.Result.Volumes)
}

func (s *AlbumTestSuite) TestGetAlbums() {
	data, err := s.cl.GetAlbums([]int64{1944241})
	s.require.Nil(err)

	var title = data.Result[0].Title
	s.require.Equal("Ocean Death", title)
}

func (s *AlbumTestSuite) TestLikeAlbum() {
	var err = s.cl.LikeAlbum(1944241)
	s.require.Nil(err)
}

func (s *AlbumTestSuite) TestUnlikeAlbum() {
	var err = s.cl.UnlikeAlbum(1944241)
	s.require.Nil(err)
}
