package goym

import (
	"context"
	"testing"

	"github.com/oklookat/goym/schema"
)

func TestLikedArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.LikedArtists(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtistLikeUnlike(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeArtist(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeArtist(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeUnlikeArtists(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeArtists(ctx, _artistIds[:4])
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeArtists(ctx, _artistIds[:4])
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtistTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.ArtistTracks(ctx, _artistIds[0], 0, 20)
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtistAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.ArtistAlbums(ctx, _artistIds[0], 0, 20, schema.SortByYear, schema.SortOrderDesc)
	if err != nil {
		t.Fatal(err)
	}
}

func TestArtistTopTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.ArtistTopTracks(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetArtistInfo(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.ArtistInfo(ctx, _artistIds[0])
	if err != nil {
		t.Fatal(err)
	}
}
