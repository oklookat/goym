package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ArtistTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *ArtistTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *ArtistTestSuite) TestLikedArtists() {
	_, err := s.cl.LikedArtists(context.Background())
	s.require.Nil(err)
}

func (s *ArtistTestSuite) TestArtistLikeUnlike() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeArtist(ctx, artistIds[0])
	s.require.Nil(err)

	// unlike
	_, err = s.cl.UnlikeArtist(ctx, artistIds[0])
	s.require.Nil(err)
}

func (s *ArtistTestSuite) TestLikeUnlikeArtists() {
	ctx := context.Background()

	// like
	_, err := s.cl.LikeArtists(ctx, artistIds[:])
	s.require.Nil(err)

	// unlike
	_, err = s.cl.UnlikeArtists(ctx, artistIds[:])
	s.require.Nil(err)
}

func (s *ArtistTestSuite) TestArtistTracks() {
	ctx := context.Background()
	resp, err := s.cl.ArtistTracks(ctx, artistIds[0], 0, 20)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result.Tracks)
}

func (s *ArtistTestSuite) TestArtistAlbums() {
	ctx := context.Background()
	resp, err := s.cl.ArtistAlbums(ctx, artistIds[0], 0, 20, schema.SortByYear, schema.SortOrderDesc)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result.Albums)
}

func (s *ArtistTestSuite) TestArtistTopTracks() {
	ctx := context.Background()
	resp, err := s.cl.ArtistTopTracks(ctx, artistIds[0])
	s.require.Nil(err)
	s.require.NotEmpty(resp.Result.Tracks)
}

func (s *ArtistTestSuite) TestGetArtistInfo() {
	ctx := context.Background()
	br, err := s.cl.ArtistInfo(ctx, artistIds[0])
	s.require.Nil(err)
	s.require.NotEmpty(br.Result.Albums)
	s.require.NotEmpty(br.Result.AllCovers)
	s.require.NotEmpty(br.Result.PopularTracks)
	s.require.NotEmpty(br.Result.SimilarArtists)
}
