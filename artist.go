package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить лайкнутых артистов.
func (c Client) LikedArtists(ctx context.Context) (schema.Response[[]schema.Artist], error) {
	// GET /users/{userId}/likes/artists
	return likesDislikes[[]schema.Artist](ctx, &c, true, "artists")
}

// Лайкнуть артиста по ID.
func (c Client) LikeArtist(ctx context.Context, id schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/likes/artists/add
	body := schema.ArtistIdRequestBody{
		ArtistId: id,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemove(ctx, &c, vals, true, "artists")
}

// Убрать лайк с артиста по ID.
func (c Client) UnlikeArtist(ctx context.Context, id schema.ID) (schema.Response[string], error) {
	// POST /users/{userId}/likes/artists/{artistId}/remove
	body := schema.ArtistIdsRequestBody{}
	body.SetIds(id)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemove(ctx, &c, vals, false, "artists")
}

// Лайкнуть артистов.
//
// Используйте LikeArtists() для лайка одного альбома.
func (c Client) LikeArtists(ctx context.Context, ids []schema.ID) (schema.Response[string], error) {
	return c.likeUnlikeArtists(ctx, ids, true)
}

// Снять лайки с артистов.
func (c Client) UnlikeArtists(ctx context.Context, ids []schema.ID) (schema.Response[string], error) {
	return c.likeUnlikeArtists(ctx, ids, false)
}

func (c Client) likeUnlikeArtists(ctx context.Context, ids []schema.ID, like bool) (schema.Response[string], error) {
	// POST /users/{userId}/likes/artists/add-multiple
	// ||
	// POST /users/{userId}/likes/artists/remove
	body := schema.ArtistIdsRequestBody{}
	body.SetIds(ids...)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return schema.Response[string]{}, err
	}
	return addRemoveMultiple(ctx, &c, vals, like, "artists")
}

// Получить список треков артиста по его ID.
func (c Client) ArtistTracks(ctx context.Context, id schema.ID, page, pageSize int) (schema.Response[*schema.ArtistTracksPaged], error) {
	// GET /artists/{artistId}/tracks
	body := schema.GetArtistTracksQueryParams{
		Page:     page,
		PageSize: pageSize,
	}
	data := &schema.Response[*schema.ArtistTracksPaged]{}

	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("artists", string(id), "tracks")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить альбомы артиста по его ID.
//
// Приложение под Windows в качестве pageSize обычно использует 50.
func (c Client) ArtistAlbums(ctx context.Context, id schema.ID, page, pageSize int, sortBy schema.SortBy, sortOrder schema.SortOrder) (schema.Response[*schema.ArtistAlbumsPaged], error) {
	// GET /artists/{artistId}/direct-albums
	body := schema.GetArtistAlbumsQueryParams{
		Page:      page,
		PageSize:  pageSize,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
	data := &schema.Response[*schema.ArtistAlbumsPaged]{}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("artists", string(id), "direct-albums")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить лучшие треки артиста по его ID.
func (c Client) ArtistTopTracks(ctx context.Context, id schema.ID) (schema.Response[*schema.ArtistTopTracks], error) {
	// GET /artists/{artistId}/track-ids-by-rating
	endpoint := genApiPath("artists", string(id), "track-ids-by-rating")
	data := &schema.Response[*schema.ArtistTopTracks]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить полную информацию об артисте по его ID.
func (c Client) ArtistInfo(ctx context.Context, id schema.ID) (schema.Response[*schema.ArtistBriefInfo], error) {
	// GET /artists/{artistId}/brief-info
	endpoint := genApiPath("artists", string(id), "brief-info")
	data := &schema.Response[*schema.ArtistBriefInfo]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}
