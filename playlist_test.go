package goym

import (
	"context"
	"testing"

	"github.com/oklookat/goym/schema"
)

func TestLikesPlaylist(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Search.
	found, err := cl.Search(ctx, "музыка в машину", 0, schema.SearchTypePlaylist, false)
	if err != nil {
		t.Fatal(err)
	}
	pl := found.Result.Playlists.Results[0]

	// Like.
	_, err = cl.LikePlaylist(ctx, pl.Kind, pl.UID)
	if err != nil {
		t.Fatal(err)
	}

	// Get liked.
	_, err = cl.LikedPlaylists(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// Unlike.
	_, err = cl.UnlikePlaylist(ctx, pl.Kind, pl.UID)
	if err != nil {
		t.Fatal(err)
	}

	// Like multiple.
	toLike := map[schema.ID]schema.ID{}
	for i, pl := range found.Result.Playlists.Results {
		toLike[pl.Kind] = pl.UID
		if i > 9 {
			break
		}
	}
	_, err = cl.LikePlaylists(ctx, toLike)
	if err != nil {
		t.Fatal(err)
	}

	// Unlike multiple.
	_, err = cl.UnlikePlaylists(ctx, toLike)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPlaylistsByKindUid(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Search.
	found, err := cl.Search(ctx, "phonk", 0, schema.SearchTypePlaylist, false)
	if err != nil {
		t.Fatal(err)
	}

	playlists := found.Result.Playlists.Results
	kindUid := map[schema.ID]schema.ID{}
	for i, p := range playlists {
		kindUid[p.Kind] = p.UID
		if i >= 6 {
			break
		}
	}

	// Get.
	_, err = cl.PlaylistsByKindUid(ctx, kindUid)
	if err != nil {
		t.Fatal(err)
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
func TestPlaylistCRUD(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// CreatePlaylist
	pl, err := cl.CreatePlaylist(ctx, "goym", "test1", schema.VisibilityPublic)
	if err != nil {
		t.Fatal(err)
	}

	// MyPlaylist
	pl, err = cl.MyPlaylist(ctx, pl.Result.Kind)
	if err != nil {
		t.Fatal(err)
	}

	// MyPlaylists
	_, err = cl.MyPlaylists(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// RenamePlaylist
	pl, err = cl.RenamePlaylist(ctx, pl.Result.Kind, "goym (renamed)")
	if err != nil {
		t.Fatal(err)
	}

	// SetPlaylistVisibility
	pl, err = cl.SetPlaylistVisibility(ctx, pl.Result.Kind, schema.VisibilityPrivate)
	if err != nil {
		t.Fatal(err)
	}

	// SetPlaylistDescription
	pl, err = cl.SetPlaylistDescription(ctx, pl.Result.Kind, "123")
	if err != nil {
		t.Fatal(err)
	}

	// AddToPlaylist
	tracksResp, err := cl.Search(ctx, "dubstep", 0, schema.SearchTypeTrack, false)
	if err != nil {
		t.Fatal(err)
	}

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
	pl, err = cl.AddToPlaylist(ctx, pl.Result, tracksToAdd)
	if err != nil {
		t.Fatal(err)
	}

	// Get with tracks.
	pl, err = cl.MyPlaylist(ctx, pl.Result.Kind)
	if err != nil {
		t.Fatal(err)
	}

	// PlaylistRecommendations
	_, err = cl.PlaylistRecommendations(ctx, pl.Result.Kind)
	if err != nil {
		t.Fatal(err)
	}

	// DeleteFromPlaylist
	pl, err = cl.DeleteTracksFromPlaylist(ctx, pl.Result, tracksToDelete)
	if err != nil {
		t.Fatal(err)
	}
	// is track actually removed?
	for _, ti := range pl.Result.Tracks {
		_, ok := tracksIds[ti.ID]
		if ok {
			t.Fatalf("track id %s not removed", ti.ID)
		}
	}

	// DeletePlaylist
	_, err = cl.DeletePlaylist(ctx, pl.Result.Kind)
	if err != nil {
		t.Fatal(err)
	}
}
