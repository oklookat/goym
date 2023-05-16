package schema

// Артист.
//
// Много полей могут быть nil. Например, когда Artist находится в составе Track.
type Artist struct {
	ID ID `json:"id"`

	// Имя.
	Name string `json:"name"`

	// Исполнитель относится к категории сборник.
	Various bool `json:"various"`

	// Исполнитель является композитором.
	Composer bool `json:"composer"`

	// Фото.
	Cover *Cover `json:"cover"`

	// По сути дублирует Cover.URI.
	OgImage *string `json:"ogImage"`

	// Жанры исполнителя.
	Genres []string `json:"genres"`

	// Количество разных вещей.
	Counts *struct {
		// Общее количество треков исполнителя, доступных в каталоге ЯМ.
		Tracks uint32 `json:"tracks"`

		// Количество собственных альбомов.
		DirectAlbums uint32 `json:"directAlbums"`

		// Количество альбомов, где представлен исполнитель.
		AlsoAlbums uint32 `json:"alsoAlbums"`

		// Количество треков, где представлен исполнитель.
		AlsoTracks uint32 `json:"alsoTracks"`
	} `json:"counts"`

	// Треки исполнителя доступны?
	Available bool `json:"available"`

	// Рейтинги исполнителя.
	Ratings *struct {
		// За месяц.
		Month uint32 `json:"month"`

		// За неделю.
		Week *uint32 `json:"week"`

		// За день.
		Day *uint32 `json:"day"`
	} `json:"ratings"`

	// Список ссылок на сайты исполнителя.
	Links []struct {
		// Заголовок ссылки.
		Title string `json:"title"`

		// Сама ссылка. YouTube, Twitter, персональный сайт, etc.
		Href string `json:"href"`

		// social | official
		Type string `json:"type"`

		// twitter | youtube | vk | telegram. Может быть nil, когда Type == official.
		SocialNetwork *string `json:"socialNetwork"`
	} `json:"links"`

	// Доступны билеты на концерт?
	TicketsAvailable *bool `json:"ticketsAvailable"`
}

type (
	ArtistBriefInfo struct {
		Artist *struct {
			Artist

			// Количество слушателей, оценивших исполнителя.
			LikesCount uint32 `json:"likesCount"`

			// Описание.
			Description struct {
				// Об артисте.
				Text string `json:"text"`

				// Ссылка на источник. Например, на Википедию.
				URI string `json:"uri"`
			} `json:"description"`

			// Откуда артист? Пример: ["Франция"].
			Countries []string `json:"countries"`

			// Год начала карьеры.
			InitDate string `json:"initDate"`

			// Год конца карьеры.
			EndDate string `json:"endDate"`

			// Ссылка на страницу артиста в английской Википедии.
			EnWikipediaLink string `json:"enWikipediaLink"`

			// Список вариантов ввода имени исполнителя в поисковой строке
			//
			// (для облегчения поиска музыки на смартфоне в режиме офлайн).
			//
			// Пример: ["Daft punk", "Дафт Панк", "duft pank", "ダフトパンク"]
			DbAliases []string `json:"dbAliases"`
		} `json:"artist"`

		// Собственные альбомы исполнителя (где он указан исполнителем), в базовой информации.
		Albums []Album `json:"albums"`

		// Альбомы, где представлен исполнитель (где он указан исполнителем), в базовой информации.
		AlsoAlbums []Album `json:"alsoAlbums"`

		// Популярные треки, в базовой информации.
		PopularTracks []Track `json:"popularTracks"`

		// Похожие (по стилю) исполнители, в базовой информации.
		SimilarArtists []Artist `json:"similarArtists"`

		// Все изображения исполнителя.
		AllCovers []Cover `json:"allCovers"`

		// Список концертов исполнителя.
		Concerts      []any   `json:"concerts"`
		Videos        []Video `json:"videos"`
		Clips         []any   `json:"clips"`
		Vinyls        []any   `json:"vinyls"`
		HasPromotions bool    `json:"hasPromotions"`
		LastReleases  []any   `json:"lastReleases"`
		Stats         struct {
			LastMonthListeners uint32 `json:"lastMonthListeners"`
		} `json:"stats"`
		CustomWave struct {
			Title        string `json:"title"`
			AnimationURL string `json:"animationUrl"`
		} `json:"customWave"`
		PlaylistIds []struct {
			UID  ID `json:"uid"`
			Kind ID `json:"kind"`
		} `json:"playlistIds"`
		Playlists []Playlist `json:"playlists"`
	}

	// GET /artists/{artistId}/tracks
	GetArtistTracksQueryParams struct {
		// Страница.
		Page uint16 `url:"page"`

		// Кол-во результатов на странице (20, например).
		PageSize uint16 `url:"page-size"`
	}

	// GET /artists/{artistId}/direct-albums
	GetArtistAlbumsQueryParams struct {
		Page      uint16    `url:"page"`
		PageSize  uint16    `url:"page-size"`
		SortBy    SortBy    `url:"sort-by"`
		SortOrder SortOrder `url:"sort-order"`
	}

	// POST /users/{userId}/likes/artists/add
	ArtistIdRequestBody struct {
		ArtistId ID `url:"artist-id"`
	}

	ArtistTracksPaged struct {
		Pager  Pager   `json:"pager"`
		Tracks []Track `json:"tracks"`
	}

	ArtistAlbumsPaged struct {
		Pager  Pager   `json:"pager"`
		Albums []Album `json:"albums"`
	}
)

type ArtistTopTracks struct {
	Artist *Artist `json:"artist"`
	Tracks []ID    `json:"tracks"`
}

// POST /users/{userId}/likes/artists/add-multiple
//
// POST /users/{userId}/likes/artists/remove
type ArtistIdsRequestBody struct {
	// ID альбомов.
	ArtistIds []ID `url:",artistIds"`
}

func (l *ArtistIdsRequestBody) SetIds(ids ...ID) {
	l.ArtistIds = []ID{}
	l.ArtistIds = append(l.ArtistIds, ids...)
}
