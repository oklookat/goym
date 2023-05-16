package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить лайкнутые треки.
func (c Client) LikedTracks(ctx context.Context) (schema.Response[*schema.TracksLibrary], error) {
	return likesDislikes[*schema.TracksLibrary](ctx, &c, true, "tracks")
}

// Получить дизлайкнутые треки.
func (c Client) DislikedTracks(ctx context.Context) (schema.Response[*schema.TracksLibrary], error) {
	return likesDislikes[*schema.TracksLibrary](ctx, &c, false, "tracks")
}

// Лайкнуть трек.
func (c Client) LikeTrack(ctx context.Context, id schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/likes/tracks/add
	body := schema.TrackIdRequestBody{
		TrackId: id,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemove(ctx, &c, vals, true, "tracks")
}

// Лайкнуть треки.
//
// Используйте LikeTrack() для лайка одного трека.
func (c Client) LikeTracks(ctx context.Context, ids []schema.ID) (schema.Response[string], error) {
	return c.likeUnlikeTracks(ctx, ids, true)
}

// Снять лайки с треков.
func (c Client) UnlikeTracks(ctx context.Context, ids []schema.ID) (schema.Response[string], error) {
	return c.likeUnlikeTracks(ctx, ids, false)
}

func (c Client) likeUnlikeTracks(ctx context.Context, ids []schema.ID, like bool) (schema.Response[string], error) {
	// POST /users/{userId}/likes/tracks/add-multiple
	// ||
	// POST /users/{userId}/likes/tracks/remove
	body := schema.TrackIdsRequestBody{}
	body.SetIds(ids...)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemoveMultiple(ctx, &c, vals, like, "tracks")
}

// Получить трек по id.
func (c Client) Track(ctx context.Context, trackId schema.ID) (schema.Response[[]schema.Track], error) {
	// GET /tracks/{trackId}
	endpoint := genApiPath("tracks", trackId.String())
	data := &schema.Response[[]schema.Track]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить треки по id.
func (c Client) Tracks(ctx context.Context, trackIds []schema.ID) (schema.Response[[]schema.Track], error) {
	// POST /tracks
	data := &schema.Response[[]schema.Track]{}

	if trackIds == nil {
		return *data, nil
	}
	body := schema.GetTracksByIdsRequestBody{
		TrackIds: trackIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("tracks")
	resp, err := c.Http.R().SetError(data).SetFormUrlValues(vals).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить информацию о загрузке трека.
func (c Client) TrackDownloadInfo(ctx context.Context, id schema.ID) (schema.Response[[]schema.TrackDownloadInfo], error) {
	// GET /tracks/{trackId}/download-info
	endpoint := genApiPath("tracks", id.String(), "download-info")
	data := &schema.Response[[]schema.TrackDownloadInfo]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить дополнительную информацию о треке (текст песни, видео, etc).
func (c Client) TrackSupplement(ctx context.Context, id schema.ID) (schema.Response[*schema.Supplement], error) {
	// GET /tracks/{trackId}/supplement
	endpoint := genApiPath("tracks", id.String(), "supplement")
	data := &schema.Response[*schema.Supplement]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить похожие треки.
func (c Client) SimilarTracks(ctx context.Context, id schema.ID) (schema.Response[*schema.SimilarTracks], error) {
	// GET /tracks/{trackId}/similar
	endpoint := genApiPath("tracks", id.String(), "similar")
	data := &schema.Response[*schema.SimilarTracks]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}
