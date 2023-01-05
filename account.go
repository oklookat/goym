package goym

import "github.com/oklookat/goym/holly"

// Получить информацию об аккаунте.
func (c *Client) GetAccountStatus() (data *GetResponse[*Status], err error) {
	data = &GetResponse[*Status]{}

	var endpoint = genApiPath([]string{"account", "status"})
	var resp *holly.Response
	resp, err = c.self.R().SetResult(data).SetError(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}
