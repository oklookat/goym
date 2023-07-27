package goym

import (
	"context"
	"testing"

	"github.com/oklookat/goym/schema"
)

func TestSearch(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	// Track.
	_, err := cl.Search(ctx, "трек", 0, schema.SearchTypeTrack, false)
	if err != nil {
		t.Fatal(err)
	}

	// Any.
	_, err = cl.Search(ctx, "хороший", 0, schema.SearchTypeAll, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchSuggest(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.SearchSuggest(ctx, "emine")
	if err != nil {
		t.Fatal(err)
	}
}
