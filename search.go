package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Поиск.
//
// text - текст запроса.
//
// page - страница (первая начинается с нуля).
//
// what - тип поиска.
//
// Если тип поиска не SearchTypeAll - поле Bests будет nil.
//
// Например: если тип поиска будет "artist", то
// поля best, playlists, и подобные, будут nil (кроме поля Artists).
//
// noCorrect - не исправлять опечатки?
func (c Client) Search(ctx context.Context, text string, page int, what schema.SearchType, noCorrect bool) (schema.Response[*schema.Search], error) {
	// GET /search
	data := &schema.Response[*schema.Search]{}

	query := schema.SearchQueryParams{
		Text:      text,
		Page:      page,
		Type:      what,
		NoCorrect: noCorrect,
	}

	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("search")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Подсказать что-нибудь по поисковому запросу.
//
// например: SearchSuggest("emine")
func (c Client) SearchSuggest(ctx context.Context, part string) (schema.Response[*schema.Suggestions], error) {
	// GET /search/suggest
	query := schema.SearchSuggestQueryParams{
		Part: part,
	}
	data := &schema.Response[*schema.Suggestions]{}

	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("search", "suggest")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}
