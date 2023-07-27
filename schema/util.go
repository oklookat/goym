package schema

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func ParamsToValues(s any) (url.Values, error) {
	return query.Values(s)
}
