package schema

// Лейбл звукозаписи.
type Label struct {
	// ID лейбла.
	ID int64 `json:"id"`

	// Имя лейбла.
	Name string `json:"name"`
}
