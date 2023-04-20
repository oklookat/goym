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
) error {
	// POST /users/{userId}/likes/ENTITIES/add-multiple
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add-multiple"
	if !add {
		endEndPoint = "remove"
	}

	endpoint := genApiPath("users", client.userId, "likes", entityName, endEndPoint)
	data := &schema.Response[any]{}
	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// If add: use "id" request body.
//
// If remove: use "ids" request body.
func addRemove(
	ctx context.Context,
	client *Client,
	vals url.Values,
	add bool,
	entityName string) error {
	// POST /users/{userId}/likes/ENTITIES/add
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	endEndPoint := "add"
	if !add {
		endEndPoint = "remove"
	}

	endpoint := genApiPath("users", client.userId, "likes", entityName, endEndPoint)
	data := &schema.Response[any]{}
	resp, err := client.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

func likesDislikes[T any](ctx context.Context, client *Client, likes bool, entityName string) (T, error) {
	// GET /users/{userId}/likes/ENTITY
	// ||
	// GET /users/{userId}/dislikes/ENTITY
	lord := "likes"
	if !likes {
		lord = "dislikes"
	}

	endpoint := genApiPath("users", client.userId, lord, entityName)
	data := &schema.Response[T]{}
	resp, err := client.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
