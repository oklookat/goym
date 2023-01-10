package schema

// Артист.
type Artist struct {
	ID int64 `json:"id"`

	// Имя.
	Name     string `json:"name"`
	Various  bool   `json:"various"`
	Composer bool   `json:"composer"`

	// (?) Аватар.
	Cover *Cover `json:"cover"`

	// Жанры.
	Genres []string `json:"genres"`
}

// GET /artists/{artistId}/tracks
type GetArtistTracksQueryParams struct {
	// Страница.
	Page int `url:"page"`

	// Кол-во результатов на странице (20, например).
	PageSize int `url:"page-size"`
}

// GET /artists/{artistId}/direct-albums
type GetArtistAlbumsQueryParams struct {
	// Страница.
	Page int `url:"page"`

	// Кол-во результатов на странице (20, например).
	PageSize int `url:"page-size"`

	// year | rating
	SortBy string `url:"sort-by"`
}

// POST /users/{userId}/likes/artists/add
type LikeArtistRequestBody struct {
	ArtistId int64 `url:"artist-id"`
}
