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

func (s PlaylistTestSuite) TestGetMyPlaylists() {
	pls, err := s.cl.GetMyPlaylists(context.Background())
	s.require.Nil(err)
	s.require.NotEmpty(pls)
	s.require.Positive(pls[0].Kind)
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
	var ctx = context.Background()

	pl, err := s.cl.CreatePlaylist(ctx, "goymtesting", schema.VisibilityPublic)
	s.require.Nil(err)
	s.require.NotNil(pl.Kind)

	pl2, err := s.cl.GetMyPlaylistByKind(ctx, pl.Kind)
	s.require.Nil(err)
	s.require.Equal(pl.Kind, pl2.Kind)

	pl3, err := s.cl.RenamePlaylist(ctx, pl2, "goymtesting (renamed)")
	s.require.Nil(err)
	s.require.Equal(pl2.Kind, pl3.Kind)

	pl4, err := s.cl.ChangePlaylistVisibility(ctx, pl3, schema.VisibilityPrivate)
	s.require.Nil(err)
	s.require.Equal(pl3.Kind, pl4.Kind)

	// AddTracksToPlaylist
	tracksResp, err := s.cl.Search(ctx, "dubstep", 0, schema.SearchTypeTrack, false)
	s.require.Nil(err)
	s.require.NotEmpty(tracksResp.Tracks)
	var tracks = tracksResp.Tracks.Results
	// 10 tracks
	var tracksLittle = []*schema.Track{}
	for i := range tracks {
		tracksLittle = append(tracksLittle, tracks[i])
		if len(tracksLittle) == 10 {
			break
		}
	}
	pl5, err := s.cl.AddTracksToPlaylist(ctx, pl, tracksLittle)
	s.require.Nil(err)
	s.require.Equal(pl4.Kind, pl5.Kind)
	s.require.Greater(pl5.Revision, pl4.Revision)
	// get with tracks
	pl5, err = s.cl.GetMyPlaylistByKind(ctx, pl5.Kind)
	s.require.Nil(err)

	// GetPlaylistRecommendations
	recs, err := s.cl.GetPlaylistRecommendations(ctx, pl5)
	s.require.Nil(err)
	s.require.NotEmpty(recs.Tracks)
	s.require.Positive(recs.Tracks[0].ID)

	// DeleteTrackFromPlaylist (remove)
	var trackToDelete = pl5.Tracks[0]
	pl6, err := s.cl.DeleteTrackFromPlaylist(ctx, pl5, trackToDelete)
	s.require.Nil(err)
	s.require.Equal(pl5.Kind, pl6.Kind)
	s.require.Greater(pl6.Revision, pl5.Revision)
	// is track actually removed?
	for _, ti := range pl6.Tracks {
		s.require.NotEqual(ti.Track.ID, trackToDelete.ID)
	}

	// DeletePlaylist
	err = s.cl.DeletePlaylist(ctx, pl6)
	s.require.Nil(err)
}
