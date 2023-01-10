package goym

import (
	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SearchTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *SearchTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s SearchTestSuite) TestSearch() {
	// 🤘🤘🤘
	data, err := s.cl.Search("король и шут бедняжка", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)

	var ideed = data.Tracks.Results[0].Title
	s.require.Equal("Бедняжка", ideed)
}

func (s SearchTestSuite) TestSearchSuggest() {
	data, err := s.cl.SearchSuggest("emine")
	s.require.Nil(err)

	var suggestion = data.Suggestions[0]
	s.require.Equal("eminem", suggestion)
}
