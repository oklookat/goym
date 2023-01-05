package goym

import "github.com/oklookat/goym/holly"

// Получить информацию об аккаунте.
func (c *Client) GetAccountStatus() (data *TypicalResponse[*Status], err error) {
	data = &TypicalResponse[*Status]{}

	var endpoint = genApiPath([]string{"account", "status"})
	var resp *holly.Response
	resp, err = c.self.R().SetResult(data).SetError(data).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}
