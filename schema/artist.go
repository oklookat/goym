package schema

import (
	"encoding/json"
)

// Артист.
type Artist struct {
	ID UniqueID `json:"-"`

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
		Tracks       uint32 `json:"tracks"`
		DirectAlbums uint32 `json:"directAlbums"`
		AlsoAlbums   uint32 `json:"alsoAlbums"`
		AlsoTracks   uint32 `json:"alsoTracks"`
	} `json:"counts"`

	Available bool `json:"available"`

	// Это не количество прослушиваний.
	//
	// Похоже на какую-то позицию в топе артистов.
	//
	// Если артист непопулярен, будет доступно только поле Month.
	Ratings *struct {
		// В месяц.
		Month uint32 `json:"month"`
		// В неделю.
		Week *uint32 `json:"week"`
		// В день.
		Day *uint32 `json:"day"`
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

	// ---- Поля ниже доступны, при получении Brief Info. ---- //

	// Сколько людей лайкнули артиста?
	LikesCount *uint32 `json:"likesCount"`

	// Описание.
	Description *struct {
		// Об артисте.
		Text string `json:"text"`

		// Ссылка на источник. Например, на Википедию.
		URI string `json:"uri"`
	} `json:"description"`

	// Откуда артист? Пример: ["Франция"].
	Countries []string `json:"countries"`

	// Год начала карьеры.
	InitDate *string `json:"initDate"`

	// Год конца карьеры.
	EndDate *string `json:"endDate"`

	// Ссылка на страницу артиста в английской Википедии.
	EnWikipediaLink string `json:"enWikipediaLink"`

	// Стили написания имени артиста.
	//
	// Пример: ["Daft punk", "Дафт Панк", "duft pank", "ダフトパンク"]
	DbAliases []string `json:"dbAliases"`

	// ---- Поля выше доступны, при получении Brief Info. ---- //
}

func (a *Artist) UnmarshalJSON(data []byte) error {
	var dem = func(id UniqueID, data []byte) error {
		type Fake Artist
		var faked Fake
		if err := json.Unmarshal(data, &faked); err != nil {
			return err
		}
		*a = Artist(faked)
		a.ID = id
		return nil
	}
	return unmarshalID(dem, data)
}

type ArtistBriefInfo struct {
	Artist         Artist    `json:"artist"`
	Albums         []*Album  `json:"albums"`
	AlsoAlbums     []*Album  `json:"alsoAlbums"`
	PopularTracks  []*Track  `json:"popularTracks"`
	SimilarArtists []*Artist `json:"similarArtists"`
	AllCovers      []*Cover  `json:"allCovers"`
	Concerts       []any     `json:"concerts"`
	Videos         []*Video  `json:"videos"`
	Clips          []any     `json:"clips"`
	Vinyls         []any     `json:"vinyls"`
	HasPromotions  bool      `json:"hasPromotions"`
	LastReleases   []any     `json:"lastReleases"`
	Stats          struct {
		LastMonthListeners uint32 `json:"lastMonthListeners"`
	} `json:"stats"`
	CustomWave struct {
		Title        string `json:"title"`
		AnimationURL string `json:"animationUrl"`
	} `json:"customWave"`
	PlaylistIds []struct {
		UID  UniqueID `json:"uid"`
		Kind KindID   `json:"kind"`
	} `json:"playlistIds"`
	Playlists []*Playlist `json:"playlists"`
}

// GET /artists/{artistId}/tracks
type GetArtistTracksQueryParams struct {
	// Страница.
	Page uint16 `url:"page"`

	// Кол-во результатов на странице (20, например).
	PageSize uint16 `url:"page-size"`
}

// GET /artists/{artistId}/direct-albums
type GetArtistAlbumsQueryParams struct {
	// Страница.
	Page uint16 `url:"page"`

	// Кол-во результатов на странице (20, например).
	PageSize uint16 `url:"page-size"`

	SortBy SortBy `url:"sort-by"`
}

// POST /users/{userId}/likes/artists/add
type LikeArtistRequestBody struct {
	ArtistId UniqueID `url:"artist-id"`
}

type ArtistTracksPaged struct {
	Pager  Pager    `json:"pager"`
	Tracks []*Track `json:"tracks"`
}

type ArtistAlbumsPaged struct {
	Pager  Pager    `json:"pager"`
	Albums []*Album `json:"albums"`
}

type ArtistTopTracks struct {
	Artist *Artist    `json:"artist"`
	Tracks []UniqueID `json:"tracks"`
}

// Разбираемся с ID.
func (a *ArtistTopTracks) UnmarshalJSON(data []byte) error {
	// TODO: сделать демаршал []string в []int64
	type real struct {
		Artist *Artist  `json:"artist"`
		Tracks []string `json:"tracks"`
	}
	var realVal = &real{}
	if err := json.Unmarshal(data, realVal); err != nil {
		return err
	}
	if len(realVal.Tracks) == 0 {
		return nil
	}
	a.Tracks = make([]UniqueID, 0)
	for _, id := range realVal.Tracks {
		var uid UniqueID = 0
		err := uid.FromString(id)
		if err != nil {
			return err
		}
		a.Tracks = append(a.Tracks, uid)
	}

	return nil
}
