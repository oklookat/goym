package schema

type (
	// Результаты поиска.
	Search struct {
		// По какому типу был выполнен поиск.
		Type SearchType `json:"type"`

		// Текущая страница. Доступно при использовании параметра type.
		Page *uint16 `json:"page"`

		// Результатов на странице. Доступно при использовании параметра type.
		PerPage *uint16 `json:"perPage"`

		// Поисковой запрос (оригинальный или исправленный).
		Text string `json:"text"`

		// ID запроса.
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
		Best *Best `json:"best"`

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

	SearchResult[T any] struct {
		// Количество результатов
		Total uint32 `json:"total"`

		// Максимальное количество результатов на странице.
		PerPage uint16 `json:"perPage"`

		// Позиция блока
		Order uint16 `json:"order"`

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

	// Лучший результат поиска
	Best struct {
		// Тип лучшего результата
		//
		// track | artist | album | playlist | video
		Type string `json:"type"`

		Text string `json:"text"`

		// Может быть nil.
		Result any `json:"result"`
	}

	// GET /search
	SearchQueryParams struct {
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
	SearchSuggestQueryParams struct {
		// Часть поискового запроса.
		Part string `url:"part"`
	}
)
