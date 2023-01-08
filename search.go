package goym

// Найти.
//
// text - текст запроса
//
// page - страница. Первая страница начинается с нуля.
//
// _type - тип. Используйте константу SearchType_.
//
// Если eType будет не SearchType_All, то в результатах поиска будет отсутствовать поле Best.
//
// exact - не исправлять запрос, искать ровно то, что написано в query.
func (c *Client) Search(text string, page uint, _type string, nocorrect bool) (*TypicalResponse[Search], error) {
	var endpoint = genApiPath([]string{"search"})
	var query = searchQuery(text, page, _type, nocorrect)

	var data = &TypicalResponse[Search]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetQueryParams(query).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Подсказать что-нибудь по поисковому запросу.
func (c *Client) SearchSuggest(query string) (*TypicalResponse[Suggestions], error) {
	var endpoint = genApiPath([]string{"search", "suggest"})

	var data = &TypicalResponse[Suggestions]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetQueryParams(searchSuggestQuery(query)).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}
