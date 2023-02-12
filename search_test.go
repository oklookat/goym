package goym

import (
	"context"

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
	found, err := s.cl.Search(context.Background(), "король и шут бедняжка", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	s.require.NotEmpty(found.Tracks.Results)
	s.require.Positive(found.Tracks.Results[0].ID)
}

func (s SearchTestSuite) TestSearchSuggest() {
	sugg, err := s.cl.SearchSuggest(context.Background(), "emine")
	s.require.Nil(err)
	s.require.NotEmpty(sugg.Suggestions)
	s.require.NotNil(sugg.Best.Result)

	suggestion := sugg.Suggestions[0]
	s.require.Equal("eminem", suggestion)
}
