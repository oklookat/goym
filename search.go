package main

import "github.com/oklookat/goym/holly"

// Найти.
//
// query - текст запроса
//
// page - страница. Первая страница начинается с нуля.
//
// eType - тип. Используйте константу SearchType_.
//
// exact - не исправлять запрос, искать ровно то, что написано в query.
func (c *Client) Search(query string, page int, eType string, exact bool) (data *GetResponse[*Search], err error) {
	data = &GetResponse[*Search]{}
	var endpoint = genApiPath([]string{"search"})

	var noCorrect = "false"
	if exact {
		noCorrect = "true"
	}

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetQueryParams(map[string]string{
		"text":      query,
		"page":      i2s(page),
		"type":      eType,
		"nocorrect": noCorrect,
	}).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Подсказать что-нибудь по поисковому запросу.
func (c *Client) SearchSuggest(query string) (data *GetResponse[*Suggestions], err error) {
	data = &GetResponse[*Suggestions]{}
	var endpoint = genApiPath([]string{"search", "suggest"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetQueryParams(map[string]string{
		"part": query,
	}).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}
