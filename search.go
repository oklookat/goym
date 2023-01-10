package goym

import "github.com/oklookat/goym/schema"

// Найти.
//
// text - текст запроса
//
// page - страница. Первая страница начинается с нуля.
//
// sType - тип поиска.
//
// Если eType будет не SearchType_All, то в результатах поиска будет отсутствовать поле Best.
//
// noCorrect - исправить опечатки?
//
// GET /search
func (c Client) Search(text string, page uint16, sType schema.SearchType, noCorrect bool) (*schema.Search, error) {
	var query = schema.SearchQueryParams{
		Text:      text,
		Page:      page,
		Type:      sType,
		NoCorrect: noCorrect,
	}
	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"search"})
	var data = &schema.TypicalResponse[*schema.Search]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Подсказать что-нибудь по поисковому запросу.
//
// GET /search/suggest
func (c Client) SearchSuggest(part string) (*schema.Suggestions[any], error) {
	var query = schema.SearchSuggestQueryParams{
		Part: part,
	}
	vals, err := schema.ParamsToValues(query)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"search", "suggest"})
	var data = &schema.TypicalResponse[*schema.Suggestions[any]]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}
