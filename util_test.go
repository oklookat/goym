package goym

import "testing"

func TestGenApiPath(t *testing.T) {
	const expected = "https://api.music.yandex.net/users/1234/playlists/list"
	result := genApiPath("users", "1234", "playlists", "list")
	if expected != result {
		t.Fatalf("expected: %s, got: %s", expected, result)
	}
}
