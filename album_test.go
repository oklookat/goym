package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
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

func (s AlbumTestSuite) getAlbumId() schema.ID {
	res, err := s.cl.Search(context.Background(), "album", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(res)
	s.require.NotEmpty(res.Albums.Results)
	id := res.Albums.Results[0].ID
	s.require.Positive(id)
	return id
}

func (s AlbumTestSuite) getAlbumIds() []schema.ID {
	res, err := s.cl.Search(context.Background(), "album", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(res)
	s.require.NotEmpty(res.Albums.Results)
	ids := []schema.ID{}
	for i, al := range res.Albums.Results {
		if i == 8 {
			break
		}
		ids = append(ids, al.ID)
	}
	return ids
}

func (s AlbumTestSuite) TestAlbum() {
	ctx := context.Background()

	// without tracks
	id := s.getAlbumId()
	data, err := s.cl.Album(ctx, id, false)
	s.require.Nil(err)
	s.require.Positive(data.ID)

	// with tracks
	data, err = s.cl.Album(ctx, 231541, true)
	s.require.Nil(err)
	s.require.Positive(data.ID)
	s.require.NotEmpty(data.Volumes)
}

func (s AlbumTestSuite) TestAlbums() {
	ids := s.getAlbumIds()
	albums, err := s.cl.Albums(context.Background(), ids)
	s.require.Nil(err)
	s.require.NotEmpty(albums)
	s.require.Positive(albums[0].ID)
}

func (s AlbumTestSuite) TestLikeUnlikeAlbum() {
	ctx := context.Background()

	found, err := s.cl.Search(ctx, "mujuice downshifting", 0, schema.SearchTypeAlbum, false)
	s.require.Nil(err)
	s.require.NotNil(found)
	s.require.NotEmpty(found.Albums.Results)
	al := found.Albums.Results[0]
	s.require.Positive(al.ID)

	// like
	err = s.cl.LikeAlbum(ctx, al.ID)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeAlbum(ctx, al.ID)
	s.require.Nil(err)
}

func (s AlbumTestSuite) TestLikeUnlikeAlbums() {
	ctx := context.Background()
	ids := s.getAlbumIds()

	// like
	err := s.cl.LikeAlbums(ctx, ids)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeAlbums(ctx, ids)
	s.require.Nil(err)
}

func (s AlbumTestSuite) TestLikedAlbums() {
	albums, err := s.cl.LikedAlbums(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(albums)
	s.require.Positive(albums[0].ID)
}
