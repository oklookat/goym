package goym

import "github.com/oklookat/goym/schema"

// Получить информацию об аккаунте.
func (c Client) GetAccountStatus() (*schema.Status, error) {
	var endpoint = genApiPath([]string{"account", "status"})

	var data = &schema.TypicalResponse[*schema.Status]{}
	resp, err := c.self.R().SetResult(data).SetError(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data.Result, err
}
