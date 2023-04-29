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

func (s PlaylistTestSuite) TestLikesPlaylist() {
	// Search.
	ctx := context.Background()
	found, err := s.cl.Search(ctx, "музыка в машину", 0, schema.SearchTypePlaylist, false)
	s.require.Nil(err)
	s.require.NotNil(found.Playlists)
	s.require.NotEmpty(found.Playlists.Results)
	pl := found.Playlists.Results[0]

	// Like.
	err = s.cl.LikePlaylist(ctx, pl.Kind, pl.UID)
	s.require.Nil(err)

	// Get liked.
	playlists, err := s.cl.LikedPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(playlists)
	s.require.NotEmpty(playlists[0].Playlist.Title)

	// Unlike.
	err = s.cl.UnlikePlaylist(ctx, pl.Kind, pl.UID)
	s.require.Nil(err)

	// Like multiple.
	toLike := map[schema.ID]schema.ID{}
	for i, pl := range found.Playlists.Results {
		toLike[pl.Kind] = pl.UID
		if i >= 10 {
			break
		}
	}
	err = s.cl.LikePlaylists(ctx, toLike)
	s.require.Nil(err)

	// Unlike multiple.
	err = s.cl.UnlikePlaylists(ctx, toLike)
	s.require.Nil(err)
}

func (s PlaylistTestSuite) TestPlaylistsByKindUid() {
	// Search.
	ctx := context.Background()
	found, err := s.cl.Search(ctx, "phonk", 0, schema.SearchTypePlaylist, false)
	s.require.Nil(err)
	s.require.NotNil(found.Playlists)
	s.require.NotEmpty(found.Playlists.Results)
	if len(found.Playlists.Results) < 5 {
		s.require.Fail("too few playlists")
	}

	playlists := found.Playlists.Results
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
	s.require.NotEmpty(foundByKind)
	if len(foundByKind) <= 5 {
		s.require.Fail("too few UidKind playlists")
	}
}

// CreatePlaylist()
// GetUserPlaylistById()
// RenamePlaylist()
// DeletePlaylist()
// ChangePlaylistVisibility()
// AddTracksToPlaylist()
// DeleteTrackFromPlaylist()
// GetPlaylistRecommendations()
func (s PlaylistTestSuite) TestPlaylistCRUD() {
	ctx := context.Background()

	// Create.
	pl, err := s.cl.CreatePlaylist(ctx, "goymtesting", schema.VisibilityPublic)
	s.require.Nil(err)
	s.require.NotNil(pl.Kind)

	// Get.
	pl2, err := s.cl.MyPlaylist(ctx, pl.Kind)
	s.require.Nil(err)
	s.require.Equal(pl.Kind, pl2.Kind)
	s.require.NotNil(pl2.Tracks)

	// Get many.
	pls, err := s.cl.MyPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(pls)
	s.require.Positive(pls[0].Kind)

	// Rename.
	pl3, err := s.cl.RenamePlaylist(ctx, pl2.Kind, "goymtesting (renamed)")
	s.require.Nil(err)
	s.require.Equal(pl2.Kind, pl3.Kind)

	// Change vis.
	pl4, err := s.cl.SetPlaylistVisibility(ctx, pl3.Kind, schema.VisibilityPrivate)
	s.require.Nil(err)
	s.require.Equal(pl3.Kind, pl4.Kind)

	// Add tracks.
	tracksResp, err := s.cl.Search(ctx, "dubstep", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	s.require.NotEmpty(tracksResp.Tracks)
	tracks := tracksResp.Tracks.Results
	// 10 tracks
	tracksLittle := []*schema.Track{}
	for i := range tracks {
		tracksLittle = append(tracksLittle, tracks[i])
		if len(tracksLittle) >= 10 {
			break
		}
	}
	pl5, err := s.cl.AddToPlaylist(ctx, pl, tracksLittle)
	s.require.Nil(err)
	s.require.Equal(pl4.Kind, pl5.Kind)
	s.require.Greater(*pl5.Revision, *pl4.Revision)

	// Get with tracks.
	pl5, err = s.cl.MyPlaylist(ctx, pl5.Kind)
	s.require.Nil(err)

	// Get recs.
	recs, err := s.cl.PlaylistRecommendations(ctx, pl5.Kind)
	s.require.Nil(err)
	s.require.NotEmpty(recs.Tracks)
	s.require.Positive(recs.Tracks[0].ID)

	// Delete from playlist.
	trackToDelete := pl5.Tracks[0]
	pl6, err := s.cl.DeleteFromPlaylist(ctx, pl5, trackToDelete)
	s.require.Nil(err)
	s.require.Equal(pl5.Kind, pl6.Kind)
	s.require.Greater(*pl6.Revision, *pl5.Revision)
	// is track actually removed?
	for _, ti := range pl6.Tracks {
		s.require.NotEqual(ti.Track.ID, trackToDelete.ID)
	}

	// Delete playlist.
	err = s.cl.DeletePlaylist(ctx, pl6.Kind)
	s.require.Nil(err)
}
