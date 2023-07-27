package goym

import (
	"context"
	"testing"

	"github.com/oklookat/goym/schema"
)

func TestLikedTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.LikedTracks(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDislikedTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.DislikedTracks(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeUnlikeTrack(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeTrack(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeTracks(ctx, []schema.ID{_trackIds[0]})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLikeUnlikeTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// like
	_, err := cl.LikeTracks(ctx, _trackIds[:4])
	if err != nil {
		t.Fatal(err)
	}

	// unlike
	_, err = cl.UnlikeTracks(ctx, _trackIds[:4])
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTrackById(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.Track(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestTracksById(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.Tracks(ctx, _trackIds[:4])
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrackDownloadInfo(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// get info
	_, err := cl.TrackDownloadInfo(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestTrackSupplement(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// get info
	_, err := cl.TrackSupplement(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestSimilarTracks(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.SimilarTracks(ctx, _trackIds[0])
	if err != nil {
		t.Fatal(err)
	}
}
