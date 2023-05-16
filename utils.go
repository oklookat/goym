package goym

import (
	"context"
	"net/url"

	"github.com/oklookat/goym/schema"
)

func addRemoveMultiple(
	ctx context.Context,
	client *Client,
	vals url.Values,
	add bool,
	entityName string,
) (schema.Response[string], error) {
	// POST /users/{userId}/likes/ENTITIES/add-multiple
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add-multiple"
	if !add {
		endEndPoint = "remove"
	}
	endpoint := genApiPath("users", string(client.UserId), "likes", entityName, endEndPoint)
	data := &schema.Response[string]{}

	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// If add: use "id" request body.
//
// If remove: use "ids" request body.
func addRemove(
	ctx context.Context,
	client *Client,
	vals url.Values,
	add bool,
	entityName string) (schema.Response[string], error) {
	// POST /users/{userId}/likes/ENTITIES/add
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add"
	if !add {
		endEndPoint = "remove"
	}

	endpoint := genApiPath("users", string(client.UserId), "likes", entityName, endEndPoint)
	data := &schema.Response[string]{}
	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

func likesDislikes[T any](ctx context.Context, client *Client, likes bool, entityName string) (schema.Response[T], error) {
	// GET /users/{userId}/likes/ENTITY
	// ||
	// GET /users/{userId}/dislikes/ENTITY
	lord := "likes"
	if !likes {
		lord = "dislikes"
	}
	endpoint := genApiPath("users", string(client.UserId), lord, entityName)
	data := &schema.Response[T]{}

	resp, err := client.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return *data, err
}
