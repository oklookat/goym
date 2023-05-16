package goym

import (
	"context"

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

func (s AlbumTestSuite) TestAlbum() {
	ctx := context.Background()

	// without tracks
	data, err := s.cl.Album(ctx, albumIds[0], false)
	s.require.Nil(err)

	// with tracks
	data, err = s.cl.Album(ctx, albumIds[0], true)
	s.require.Nil(err)
	s.require.NotEmpty(data.Result.Volumes)
}

func (s AlbumTestSuite) TestAlbums() {
	resp, err := s.cl.Albums(context.Background(), albumIds[:])
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result)
}

func (s AlbumTestSuite) TestLikeUnlikeAlbum() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeAlbum(ctx, albumIds[0])
	s.require.Nil(err)

	// unlike
	_, err = s.cl.UnlikeAlbum(ctx, albumIds[0])
	s.require.Nil(err)
}

func (s AlbumTestSuite) TestLikeUnlikeLikedAlbums() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeAlbums(ctx, albumIds[:])
	s.require.Nil(err)

	liked, err := s.cl.LikedAlbums(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(liked.Result)

	// unlike
	_, err = s.cl.UnlikeAlbums(ctx, albumIds[:])
	s.require.Nil(err)
}
