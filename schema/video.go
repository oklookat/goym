package schema

type (
	// Видеоклипы.
	VideoSupplement struct {
		// URL на обложку видео.
		Cover string `json:"cover"`

		// Сервис поставляющий видео.
		Provider string `json:"provider"`

		// Название видео.
		Title string `json:"title"`

		// Уникальный идентификатор видео на сервисе.
		ProviderVideoId string `json:"providerVideoId"`

		// URL на видео.
		Url string `json:"url"`

		// URL на видео, находящегося на серверах Яндекса.
		EmbedUrl string `json:"embedUrl"`

		// HTML тег для встраивания видео.
		Embed string `json:"embed"`
	}

	// Видео.
	Video struct {
		// Название видео.
		Title string `json:"title"`

		// Ссылка на изображение.
		Cover string `json:"cover"`

		// Ссылка на видео.
		EmbedUrl string `json:"embedUrl"`

		// Сервис поставляющий видео.
		Provider string `json:"provider"`

		// Уникальный идентификатор видео на сервисе.
		ProviderVideoId string `json:"providerVideoId"`

		// Ссылка на видео YouTube.
		YoutubeUrl string `json:"youtubeUrl"`

		// Ссылка на изображение.
		ThumbnailUrl string `json:"thumbnailUrl"`

		// Длительность видео в секундах.
		Duration int `json:"duration"`

		// Текст.
		Text string `json:"text"`

		// HTML тег для встраивания в разметку страницы.
		HtmlAutoPlayVideoPlayer string `json:"htmlAutoPlayVideoPlayer"`

		// example: ["RUSSIA_PREMIUM", "RUSSIA"].
		Regions []string `json:"regions"`
	}
)
