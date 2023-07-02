package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type PlaylistTestSuite struct {
	suite.Suite
	cl      *Client
	require *require.Assertions
}

func (s *PlaylistTestSuite) SetupSuite() {
	s.cl = getClient(s.T())
	s.require = s.Require()
}

func (s *PlaylistTestSuite) TestLikesPlaylist() {
	// Search.
	ctx := context.Background()
	found, err := s.cl.Search(ctx, "музыка в машину", 0, schema.SearchTypePlaylist, false)
	s.require.Nil(err)
	s.require.NotEmpty(found.Result.Playlists.Results)
	pl := found.Result.Playlists.Results[0]

	// Like.
	_, err = s.cl.LikePlaylist(ctx, pl.Kind, pl.UID)
	s.require.Nil(err)

	// Get liked.
	respPlaylists, err := s.cl.LikedPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(respPlaylists.Result)
	s.require.NotEmpty(respPlaylists.Result[0].Playlist.Title)

	// Unlike.
	_, err = s.cl.UnlikePlaylist(ctx, pl.Kind, pl.UID)
	s.require.Nil(err)

	// Like multiple.
	toLike := map[schema.ID]schema.ID{}
	for i, pl := range found.Result.Playlists.Results {
		toLike[pl.Kind] = pl.UID
		if i > 9 {
			break
		}
	}
	_, err = s.cl.LikePlaylists(ctx, toLike)
	s.require.Nil(err)

	// Unlike multiple.
	_, err = s.cl.UnlikePlaylists(ctx, toLike)
	s.require.Nil(err)
}

func (s *PlaylistTestSuite) TestPlaylistsByKindUid() {
	// Search.
	ctx := context.Background()
	found, err := s.cl.Search(ctx, "phonk", 0, schema.SearchTypePlaylist, false)
	s.require.Nil(err)

	playlists := found.Result.Playlists.Results
	kindUid := map[schema.ID]schema.ID{}
	for i, p := range playlists {
		kindUid[p.Kind] = p.UID
		if i >= 6 {
			break
		}
	}

	// Get.
	foundByKind, err := s.cl.PlaylistsByKindUid(ctx, kindUid)
	s.require.Nil(err)
	s.require.NotEmpty(foundByKind.Result)
	if len(foundByKind.Result) <= 5 {
		s.require.Fail("too few UidKind playlists")
	}
}

// CreatePlaylist()
// GetUserPlaylistById()
// RenamePlaylist()
// DeletePlaylist()
// SetPlaylistVisibility()
// SetPlaylistDescription()
// AddTracksToPlaylist()
// DeleteTrackFromPlaylist()
// GetPlaylistRecommendations()
func (s *PlaylistTestSuite) TestPlaylistCRUD() {
	ctx := context.Background()

	// CreatePlaylist
	pl, err := s.cl.CreatePlaylist(ctx, "goym", "test1", schema.VisibilityPublic)
	s.require.Nil(err)
	s.require.Equal(pl.Result.Title, "goym")
	s.require.Equal(pl.Result.Description, "test1")

	// MyPlaylist
	pl, err = s.cl.MyPlaylist(ctx, pl.Result.Kind)
	s.require.Nil(err)

	// MyPlaylists
	pls, err := s.cl.MyPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(pls)

	// RenamePlaylist
	pl, err = s.cl.RenamePlaylist(ctx, pl.Result.Kind, "goym (renamed)")
	s.require.Nil(err)

	// SetPlaylistVisibility
	pl, err = s.cl.SetPlaylistVisibility(ctx, pl.Result.Kind, schema.VisibilityPrivate)
	s.require.Nil(err)

	// SetPlaylistDescription
	pl, err = s.cl.SetPlaylistDescription(ctx, pl.Result.Kind, "123")
	s.require.Nil(err)
	s.require.Equal(pl.Result.Description, "123")

	// AddToPlaylist
	tracksResp, err := s.cl.Search(ctx, "dubstep", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)

	tracksIds := map[schema.ID]bool{}
	var tracksToAdd []schema.Track
	var tracksToDelete []schema.ID
	for i, tr := range tracksResp.Result.Tracks.Results {
		tracksToAdd = append(tracksToAdd, tr)
		tracksToDelete = append(tracksToDelete, tr.ID)
		tracksIds[tr.ID] = true
		if i >= 9 {
			break
		}
	}
	pl, err = s.cl.AddToPlaylist(ctx, *pl.Result, tracksToAdd)
	s.require.Nil(err)

	// Get with tracks.
	pl, err = s.cl.MyPlaylist(ctx, pl.Result.Kind)
	s.require.Nil(err)
	s.require.NotEmpty(pl.Result.Tracks)

	// PlaylistRecommendations
	recs, err := s.cl.PlaylistRecommendations(ctx, pl.Result.Kind)
	s.require.Nil(err)
	s.require.NotEmpty(recs.Result.Tracks)

	// DeleteFromPlaylist
	pl, err = s.cl.DeleteTracksFromPlaylist(ctx, *pl.Result, tracksToDelete)
	s.require.Nil(err)
	// is track actually removed?
	for _, ti := range pl.Result.Tracks {
		_, ok := tracksIds[ti.ID]
		if ok {
			s.Failf("fail", "track id %s not removed", ti.ID)
		}
	}

	// DeletePlaylist
	_, err = s.cl.DeletePlaylist(ctx, pl.Result.Kind)
	s.require.Nil(err)
}
