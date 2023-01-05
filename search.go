package goym

import (
	"strconv"

	"github.com/oklookat/goym/holly"
)

// Найти.
//
// query - текст запроса
//
// page - страница. Первая страница начинается с нуля.
//
// eType - тип. Используйте константу SearchType_.
//
// Если eType будет не SearchType_All, то в результатах поиска будет отсутствовать поле Best.
//
// exact - не исправлять запрос, искать ровно то, что написано в query.
func (c *Client) Search(query string, page int, eType string, exact bool) (data *TypicalResponse[*Search], err error) {
	data = &TypicalResponse[*Search]{}
	var endpoint = genApiPath([]string{"search"})

	var noCorrect = strconv.FormatBool(exact)

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetQueryParams(map[string]string{
		"text":      query,
		"page":      i2s(page),
		"type":      eType,
		"nocorrect": noCorrect,
	}).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Подсказать что-нибудь по поисковому запросу.
func (c *Client) SearchSuggest(query string) (data *TypicalResponse[*Suggestions], err error) {
	data = &TypicalResponse[*Suggestions]{}
	var endpoint = genApiPath([]string{"search", "suggest"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetQueryParams(map[string]string{
		"part": query,
	}).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}
