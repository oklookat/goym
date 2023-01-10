package goym

import (
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

func (s ArtistTestSuite) TestArtistLikeUnlike() {
	// search & get artist id
	data, err := s.cl.Search("монеточка", 0, schema.SearchTypeArtist, false)
	s.require.Nil(err)
	s.require.NotNil(data.Artists)
	s.require.NotEmpty(data.Artists.Results)
	var ar = data.Artists.Results[0]

	// like
	err = s.cl.LikeArtist(ar)
	s.require.Nil(err)

	// unlike
	err = s.cl.UnlikeArtist(ar)
	s.require.Nil(err)
}
