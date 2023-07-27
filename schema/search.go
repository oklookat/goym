package schema

import "encoding/json"

// Тип поиска.
type SearchType string

const (
	// Поиск артистов.
	SearchTypeArtist SearchType = "artist"

	// Поиск альбомов.
	SearchTypeAlbum SearchType = "album"

	// Поиск треков.
	SearchTypeTrack SearchType = "track"

	// Поиск подкастов.
	SearchTypePodcast SearchType = "podcast"

	// Поиск плейлистов.
	SearchTypePlaylist SearchType = "playlist"

	// Поиск видео.
	SearchTypeVideo SearchType = "video"

	// Поиск всего.
	SearchTypeAll SearchType = "all"
)

type (
	// Результаты поиска.
	Search struct {
		// По какому типу был выполнен поиск.
		//
		// Например: если тип поиска будет "artist", то
		// поля best, playlists, и подобные, будут пусты, кроме поля Artists.
		Type SearchType `json:"type"`

		// Текущая страница. Доступно при использовании параметра type.
		Page int `json:"page"`

		// Результатов на странице. Доступно при использовании параметра type.
		PerPage int `json:"perPage"`

		// Поисковой запрос (оригинальный или исправленный).
		Text string `json:"text"`

		// ID запроса.
		SearchRequestID string `json:"searchRequestId"`

		// Был ли исправлен запрос. Доступен при Type "all".
		MisspellCorrected bool `json:"misspellCorrected"`

		// Исправленный поисковой запрос. Не пуст, если запрос был исправлен.
		MisspellResult string `json:"misspellResult"`

		// Оригинальный поисковой запрос. Не пуст, если запрос был исправлен.
		MisspellOriginal string `json:"misspellOriginal"`

		// ID запроса.
		SearchResultID string `json:"searchResultId"`

		// Лучший результат.
		//
		// Не nil если Type == all.
		Best *Best `json:"best"`

		// Найденные треки.
		Tracks SearchResult[Track] `json:"tracks"`

		// Найденные альбомы.
		Albums SearchResult[Album] `json:"albums"`

		// Найденные артисты.
		Artists SearchResult[Artist] `json:"artists"`

		// Найденные плейлисты.
		Playlists SearchResult[Playlist] `json:"playlists"`
	}

	SearchResult[T any] struct {
		// Количество результатов
		Total int `json:"total"`

		// Максимальное количество результатов на странице.
		PerPage int `json:"perPage"`

		// Позиция блока
		Order int `json:"order"`

		Results []T `json:"results"`
	}

	// Поисковая подсказка.
	Suggestions struct {
		// Лучший результат.
		//
		// Альбом, артист, плейлист, видео, и так далее.
		Best Best `json:"best"`

		// Предложения на основе запроса.
		//
		// Например, запрос: "emine"
		//
		// Suggestions будут примерно такие:
		//
		// ["eminem", "mount eminest", "eminen", "eminem - encore"], и так далее.
		Suggestions []string `json:"suggestions"`
	}

	// GET /search
	SearchQueryParams struct {
		// Текст запроса.
		Text string `url:"text"`

		// Номер страницы.
		Page int `url:"page"`

		// Тип поиска (default = all).
		Type SearchType `url:"type"`

		// Исправлять опечатки?
		NoCorrect bool `url:"nocorrect"`
	}

	// GET /search/suggest
	SearchSuggestQueryParams struct {
		// Часть поискового запроса.
		Part string `url:"part"`
	}
)

// Лучший результат поиска
type Best struct {
	// Тип лучшего результата
	Type SearchType `json:"type"`

	Text string `json:"text"`

	// Может быть nil.
	//
	// Для удобства используйте поля Track, Artist, и так далее.
	// Это тот же Result.
	Result any `json:"result"`

	// Поля ниже не входят в ответ API. Сделаны для удобства.
	// Не nil может быть только одно из полей.

	// Лучший трек.
	Track *Track `json:"-"`

	// Лучший артист.
	Artist *Artist `json:"-"`

	// Лучший альбом.
	Album *Album `json:"-"`

	// Лучший плейлист.
	Playlist *Playlist `json:"-"`
}

func (b *Best) UnmarshalJSON(data []byte) error {
	type fake Best
	var theBest fake
	if err := json.Unmarshal(data, &theBest); err != nil {
		return err
	}
	*b = Best(theBest)
	if len(b.Type) == 0 || b.Result == nil {
		return nil
	}

	// Не проверяю ошибки, потому что это опциональный демаршал
	// и от его результата ничего не изменится.
	switch b.Type {
	case SearchTypeAlbum:
		var album Album
		b.resultUnmarshal(&album)
		b.Album = &album
	case SearchTypeArtist:
		var artist Artist
		b.resultUnmarshal(&artist)
		b.Artist = &artist
	case SearchTypeTrack:
		var track Track
		b.resultUnmarshal(&track)
		b.Track = &track
	case SearchTypePlaylist:
		var playlist Playlist
		b.resultUnmarshal(&playlist)
		b.Playlist = &playlist
	}

	return nil
}

func (b Best) resultUnmarshal(where any) error {
	bytes, err := json.Marshal(b.Result)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, where)
}
