package schema

type Cover struct {
	// Если не nil, значит остальные поля структуры будут nil.
	//
	// Example: "cover doesn't exist".
	Error *string `json:"error"`

	Custom *bool `json:"custom"`

	// Существует когда поле type = "pic".
	Dir *string `json:"dir"`

	// pic | mosaic.
	Type *string `json:"type"`

	// Существует когда поле type = "mosaic".
	ItemsUri []string `json:"itemsUri"`

	// Существует когда поле type = "pic".
	Uri *string `json:"uri"`

	Version *string `json:"version"`
}
