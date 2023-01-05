package main

import (
	"errors"

	"github.com/oklookat/goym/auth"
	"github.com/oklookat/goym/holly"
)

type Client struct {
	UserId string
	self   holly.Client
}

func New(tokens *auth.Tokens) (*Client, error) {
	if tokens == nil {
		return nil, errors.New("nil tokens")
	}

	var cl = &Client{
		self: *holly.AC(tokens.AccessToken),
	}

	// get uid
	resp, err := cl.GetAccountStatus()
	if err != nil {
		return nil, err
	}

	cl.UserId = i2s(resp.Result.Account.UID)
	return cl, err
}
