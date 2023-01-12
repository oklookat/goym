package schema

type Cover struct {
	// Если не nil, значит остальные поля структуры будут nil.
	//
	// Example: "cover doesn't exist".
	Error *string `json:"error"`

	Custom *bool `json:"custom"`

	// Существует когда поле type = "pic".
	Dir *string `json:"dir"`

	// pic | mosaic | from-artist-photos | from-album-cover.
	//
	// "from-artist-photos" используется, когда у артиста есть официальное фото в профиле ЯМ.
	//
	// "from-album-cover" используется, когда у артиста нет официального фото в профиле ЯМ,
	// в таком случае в качестве фото артиста используется обложка одного из альбомов.
	Type *string `json:"type"`

	// пример: "7d7e16a0.p.ЕЩЁ_КАКОЙ_ТО_ID/"
	Prefix *string `json:"prefix"`

	// Существует когда поле type = "mosaic".
	ItemsUri []string `json:"itemsUri"`

	// Существует, когда поле Type не mosaic.
	//
	// пример: "avatars.yandex.net/get-music-content/КАКОЙ_ТО_ID/7d7e16a0.p.ЕЩЁ_КАКОЙ_ТО_ID/%%"
	URI *string `json:"uri"`

	Version *string `json:"version"`
}
