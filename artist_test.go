package goym

import (
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

func (s *ArtistTestSuite) TestArtistLikeUnlike() {
	// search & get artist id
	data, err := s.cl.Search("монеточка", 0, string(SearchTypeArtist), false)
	s.require.Nil(err)

	var best = data.Result.Artists.Results[0].Name
	s.require.Equal("Монеточка", best)

	// like
	var artistID = data.Result.Artists.Results[0].ID
	err = s.cl.LikeArtist(artistID)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeArtist(artistID)
	s.require.Nil(err)
}
