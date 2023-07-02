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

func (s *SearchTestSuite) TestSearch() {
	// Track.
	found, err := s.cl.Search(context.Background(), "трек", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	s.require.NotNil(found.Result)
	s.require.NotEmpty(found.Result.Tracks)

	// Any.
	found, err = s.cl.Search(context.Background(), "что-то", 0, schema.SearchTypeAll, false)
	s.require.Nil(err)
	s.require.NotNil(found.Result)
}

func (s *SearchTestSuite) TestSearchSuggest() {
	sugg, err := s.cl.SearchSuggest(context.Background(), "emine")
	s.require.Nil(err)
	s.require.NotNil(sugg.Result)
	s.require.NotEmpty(sugg.Result.Suggestions)

	suggestion := sugg.Result.Suggestions[0]
	s.require.Equal("eminem", suggestion)
}
