package goym

import (
	"context"
	"testing"
)

func TestAlbum(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// without tracks
	_, err := cl.Album(ctx, _albumIds[0], false)
	if err != nil {
		t.Fatal(err)
	}

	// with tracks
	_, err = cl.Album(ctx, _albumIds[0], true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.Albums(ctx, _albumIds[:4])
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeUnlikeAlbum(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeAlbum(ctx, _albumIds[0])
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeAlbum(ctx, _albumIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeUnlikeLikedAlbums(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeAlbums(ctx, _albumIds[:4])
	if err != nil {
		t.Fatal(err)
	}

	_, err = cl.LikedAlbums(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeAlbums(ctx, _albumIds[:4])
	if err != nil {
		t.Fatal(err)
	}
}
