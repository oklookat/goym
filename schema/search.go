package schema

// Результаты поиска.
type Search struct {
	Type SearchType `json:"type"`

	// Текущая страница. Доступно при использовании параметра type.
	Page *uint16 `json:"page"`

	// Результатов на странице. Доступно при использовании параметра type.
	PerPage *uint16 `json:"perPage"`

	// Поисковой запрос (оригинальный или исправленный).
	Text string `json:"text"`

	SearchRequestID string `json:"searchRequestId"`

	// Был ли исправлен запрос. Доступен при Type "all".
	MisspellCorrected *bool `json:"misspellCorrected"`

	// Исправленный поисковой запрос. Не nil, если запрос был исправлен.
	MisspellResult *string `json:"misspellResult"`

	// Оригинальный поисковой запрос. Не nil, если запрос был исправлен.
	MisspellOriginal *string `json:"misspellOriginal"`

	// ID запроса.
	SearchResultID *string `json:"searchResultId"`

	// Лучший результат.
	//
	// Если Type не "all" - будет nil.
	//
	// Например: если тип поиска будет "artist", то
	// поля best, playlists, и подобные, будут nil (кроме поля Artists).
	Best *Best[any] `json:"best"`

	// Найденные треки.
	Tracks SearchResult[*Track] `json:"tracks"`

	// Найденные альбомы.
	Albums SearchResult[*Album] `json:"albums"`

	// Найденные эписозды подкастов.
	PodcastEpisodes SearchResult[any] `json:"podcast_episodes"`

	// Найденные артисты.
	Artists SearchResult[*Artist] `json:"artists"`

	// Найденные плейлисты.
	Playlists SearchResult[*Playlist] `json:"playlists"`

	// Найденные видео.
	Videos SearchResult[*Video] `json:"videos"`

	// Найденные подкасты.
	Podcasts SearchResult[any] `json:"podcasts"`
}

type SearchResult[T any] struct {
	// Количество результатов
	Total uint32 `json:"total"`

	// Максимальное количество результатов на странице.
	PerPage uint16 `json:"perPage"`

	// Позиция блока
	Order uint16 `json:"order"`

	Results []T `json:"results"`
}

// Лучший результат поиска
type Best[T Track | Artist | Album | Playlist | Video | any] struct {
	// Тип лучшего результата
	//
	// track | artist | album | playlist | video
	Type string `json:"type"`

	Text string `json:"text"`

	Result T `json:"result"`
}

// Поисковая подсказка.
type Suggestions[T Track | Artist | Album | Playlist | Video | any] struct {
	// Лучший результат.
	Best Best[T] `json:"best"`

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
type SearchQueryParams struct {
	// Текст запроса.
	Text string `url:"text"`

	// Номер страницы.
	Page uint16 `url:"page"`

	// Тип поиска (default = all).
	Type SearchType `url:"type"`

	// Исправлять опечатки?
	NoCorrect bool `url:"nocorrect"`
}

// GET /search/suggest
type SearchSuggestQueryParams struct {
	// Часть поискового запроса.
	Part string `url:"part"`
}
