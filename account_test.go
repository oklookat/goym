package goym

import (
	"context"
	"testing"
)

func TestAccountStatus(t *testing.T) {
	ctx := context.Background()
	cl := getClient(t)

	_, err := cl.AccountStatus(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
