package schema

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// Преобразовать struct (НЕ указатель на struct) в url.Values.
//
// Доступно для структур, название которых заканчивается на "Params" и "Body".
//
// Но не всегда. В некоторых структурах есть дополнительные методы. Читайте доки (c).
//
// После получения Values можно сделать Encode(), и отправить GET или POST (request body).
func ParamsToValues(s any) (url.Values, error) {
	return query.Values(s)
}
