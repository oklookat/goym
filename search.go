package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Поиск.
//
// text - текст запроса.
//
// page - страница (первая страница начинается с нуля).
//
// what - тип поиска.
//
// Если тип поиска не SearchTypeAll - поле Bests будет nil.
//
// Например: если тип поиска будет "artist", то
// поля best, playlists, и подобные, будут nil (кроме поля Artists).
//
// noCorrect - исправить опечатки?
func (c Client) Search(ctx context.Context, text string, page uint16, what schema.SearchType, noCorrect bool) (*schema.Search, error) {
	// GET /search
	query := schema.SearchQueryParams{
		Text:      text,
		Page:      page,
		Type:      what,
		NoCorrect: noCorrect,
	}
	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("search")
	data := &schema.Response[*schema.Search]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Подсказать что-нибудь по поисковому запросу.
//
// например: SearchSuggest("emine")
func (c Client) SearchSuggest(ctx context.Context, part string) (*schema.Suggestions[any], error) {
	// GET /search/suggest
	query := schema.SearchSuggestQueryParams{
		Part: part,
	}
	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("search", "suggest")
	data := &schema.Response[*schema.Suggestions[any]]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
