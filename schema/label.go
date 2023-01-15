package schema

// Лейбл звукозаписи.
type Label struct {
	// ID лейбла.
	ID UniqueID `json:"id"`

	// Имя лейбла.
	Name string `json:"name"`
}
