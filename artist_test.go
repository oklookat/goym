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

func (s ArtistTestSuite) TestGetLikedArtists() {
	artists, err := s.cl.GetLikedArtists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(artists)
	s.require.Positive(artists[0].ID)
}

func (s ArtistTestSuite) TestArtistLikeUnlike() {
	var ctx = context.Background()

	// search & get artist id
	found, err := s.cl.Search(ctx, "монеточка", 0, schema.SearchTypeArtist, false)
	s.require.Nil(err)
	s.require.NotNil(found.Artists)
	s.require.NotEmpty(found.Artists.Results)
	var ar = found.Artists.Results[0]
	s.require.Positive(ar.ID)

	// like
	err = s.cl.LikeArtist(ctx, ar)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeArtist(ctx, ar)
	s.require.Nil(err)
}
