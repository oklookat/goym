package goym

// Получить информацию об аккаунте.
func (c *Client) GetAccountStatus() (*TypicalResponse[Status], error) {
	var endpoint = genApiPath([]string{"account", "status"})

	var data = &TypicalResponse[Status]{}
	resp, err := c.self.R().SetResult(data).SetError(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}
