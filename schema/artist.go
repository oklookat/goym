package schema

import "encoding/json"

// Артист.
type Artist struct {
	ID int64 `json:"-"`

	// Имя.
	Name     string `json:"name"`
	Various  bool   `json:"various"`
	Composer bool   `json:"composer"`

	// Фото.
	Cover *Cover `json:"cover"`

	// По сути дублирует Cover.URI.
	OgImage string `json:"ogImage"`

	// Жанры.
	Genres []string `json:"genres"`

	// Количество разных вещей.
	Counts struct {
		Tracks       int `json:"tracks"`
		DirectAlbums int `json:"directAlbums"`
		AlsoAlbums   int `json:"alsoAlbums"`
		AlsoTracks   int `json:"alsoTracks"`
	} `json:"counts"`

	Available bool `json:"available"`

	// Это не количество прослушиваний.
	//
	// Похоже на какую-то позицию в топе артистов.
	//
	// Если артист непопулярен, будет доступно только поле Month.
	Ratings *struct {
		// В месяц.
		Month int `json:"month"`
		// В неделю.
		Week *int `json:"week"`
		// В день.
		Day *int `json:"day"`
	} `json:"ratings"`

	// Ссылки на ресурсы артиста (сайты, соц.сети).
	Links []struct {
		// Заголовок ссылки.
		Title string `json:"title"`
		// Сама ссылка. YouTube, Twitter, персональный сайт, etc.
		Href string `json:"href"`
		// social | official
		Type string `json:"type"`
		// twitter | youtube | vk | telegram. Nil, вероятно когда Type == official.
		SocialNetwork *string `json:"socialNetwork"`
	} `json:"links"`

	// Доступны билеты на концерт?
	TicketsAvailable bool `json:"ticketsAvailable"`
}

// Разбираемся с ID.
func (a *Artist) UnmarshalJSON(data []byte) error {
	var dem = func(id int64, data []byte) error {
		type ArtistFake Artist
		var faked ArtistFake
		if err := json.Unmarshal(data, &faked); err != nil {
			return err
		}
		*a = Artist(faked)
		a.ID = id
		return nil
	}
	return unmarshalID(dem, data)
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
