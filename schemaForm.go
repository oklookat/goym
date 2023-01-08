package goym

// {"album-ids": "id,id,id,id,id"}.
func formAlbumIds(albumIds []int64) map[string]string {
	return map[string]string{
		"album-ids": i64Join(albumIds),
	}
}

// {"album-id": "id"}.
func formAlbumId(albumId int64) map[string]string {
	return map[string]string{
		"album-id": i2s(albumId),
	}
}

// {"artist-id": "id"}.
func formArtistId(artistId int64) map[string]string {
	return map[string]string{
		"artist-id": i2s(artistId),
	}
}

// Параметры для поиска.
func searchQuery(text string, page uint, _type string, nocorrect bool) map[string]string {
	return map[string]string{
		"text":      text,
		"page":      i2s(page),
		"type":      _type,
		"nocorrect": b2s(nocorrect),
	}
}

// Параметры для поисковых подсказок.
func searchSuggestQuery(part string) map[string]string {
	return map[string]string{
		"part": part,
	}
}

// {"track-ids": "id,id,id,id,id"}.
func formTrackIds(trackIds []int64) map[string]string {
	return map[string]string{
		"track-ids": i64Join(trackIds),
	}
}

// {"track-id": "id"}.
func formTrackId(trackId int64) map[string]string {
	return map[string]string{
		"track-id": i2s(trackId),
	}
}

// {"title": "hello", "visibility": "public"}
func formTitleVisibility(title string, public bool) map[string]string {
	return map[string]string{
		"title":      title,
		"visibility": visibilityToString(public),
	}
}

// {"value": "hello"}
func formValue(val string) map[string]string {
	return map[string]string{
		"value": val,
	}
}
