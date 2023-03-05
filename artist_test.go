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

func (s ArtistTestSuite) getArtist() *schema.Artist {
	// search & get artist id
	found, err := s.cl.Search(context.Background(), "daft punk", 0, schema.SearchTypeArtist, false)
	s.require.Nil(err)
	s.require.NotNil(found.Artists)
	s.require.NotEmpty(found.Artists.Results)
	ar := found.Artists.Results[0]
	s.require.Positive(ar.ID)
	return ar
}

func (s ArtistTestSuite) getNonameArtist() *schema.Artist {
	// search & get artist id
	found, err := s.cl.Search(context.Background(), "LLLL", 0, schema.SearchTypeArtist, false)
	s.require.Nil(err)
	s.require.NotNil(found.Artists)
	s.require.NotEmpty(found.Artists.Results)
	ar := found.Artists.Results[0]
	s.require.Positive(ar.ID)
	return ar
}

func (s ArtistTestSuite) TestLikedArtists() {
	artists, err := s.cl.LikedArtists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(artists)
	s.require.Positive(artists[0].ID)
}

func (s ArtistTestSuite) TestArtistLikeUnlike() {
	ctx := context.Background()
	ar := s.getArtist()

	// like
	err := s.cl.LikeArtist(ctx, ar.ID)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeArtist(ctx, ar.ID)
	s.require.Nil(err)
}

func (s ArtistTestSuite) TestArtistTracks() {
	ctx := context.Background()
	ar := s.getArtist()
	resp, err := s.cl.ArtistTracks(ctx, ar.ID, 0, 20)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Tracks)
	s.require.Positive(resp.Tracks[0].ID)
}

func (s ArtistTestSuite) TestArtistAlbums() {
	ctx := context.Background()
	ar := s.getArtist()
	resp, err := s.cl.ArtistAlbums(ctx, ar.ID, 0, 20, schema.SortByYear)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Albums)
	s.require.Positive(resp.Albums[0].ID)
}

func (s ArtistTestSuite) TestArtistTopTracks() {
	ctx := context.Background()
	ar := s.getArtist()
	resp, err := s.cl.ArtistTopTracks(ctx, ar.ID)
	s.require.Nil(err)
	s.require.NotEmpty(resp.Tracks)
	s.require.Positive(resp.Tracks[0])
}

func (s ArtistTestSuite) TestGetArtistInfo() {
	verify := func(ar *schema.Artist, br *schema.ArtistBriefInfo) {
		s.require.Equal(ar.ID, br.Artist.ID)
		s.require.NotEmpty(br.Albums)
		s.require.NotEmpty(br.AllCovers)
		s.require.NotEmpty(br.PopularTracks)
		s.require.NotEmpty(br.SimilarArtists)
	}
	ctx := context.Background()
	ar := s.getArtist()
	br, err := s.cl.ArtistInfo(ctx, ar.ID)
	s.require.Nil(err)
	verify(ar, br)

	ar = s.getNonameArtist()
	br, err = s.cl.ArtistInfo(ctx, ar.ID)
	s.require.Nil(err)
	verify(ar, br)
}
