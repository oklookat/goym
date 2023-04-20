package goym

import (
	"context"
	"math"

	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DebugTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *DebugTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s DebugTestSuite) TestLikedPlaylistsPager() {
	// Search.
	ctx := context.Background()

	var page uint16
	toLike := map[schema.ID]schema.ID{}

	var totalPages float64
	for len(toLike) >= 100 || totalPages >= float64(page) {
		found, err := s.cl.Search(ctx, "музыка в машину", page, schema.SearchTypePlaylist, false)
		s.require.Nil(err)
		s.require.NotNil(found.Playlists)
		s.require.NotEmpty(found.Playlists.Results)
		totalPages = math.Ceil(float64(*found.PerPage) / float64(*found.Page))

		for i, pl := range found.Playlists.Results {
			toLike[pl.Kind] = pl.UID
			if i >= 100 {
				break
			}
		}

		page++
	}

	err := s.cl.LikePlaylists(ctx, toLike)
	s.require.Nil(err)

	// Get liked.
	playlists, err := s.cl.LikedPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(playlists)
	s.require.NotEmpty(playlists[0].Playlist.Title)

	// Unlike multiple.
	err = s.cl.UnlikePlaylists(ctx, toLike)
	s.require.Nil(err)

	// при ~47 лайкнутых плейлистах в Pager ничего не менялось.
	// может пагинация плейлистов вообще не работает?
}
