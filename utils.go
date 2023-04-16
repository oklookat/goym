package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

func likeUnlikeMultiple(ctx context.Context,
	entityName string,
	like bool,
	client *Client,
	body likeUnlikeMultiplyBody,
	ids []schema.ID) error {
	// POST /users/{userId}/likes/ENTITIES/add-multiple
	// ||
	// POST /users/{userId}/likes/ENTITIES/remove
	body.SetIds(ids)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endEndPoint := "add-multiple"
	if !like {
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

type likeUnlikeMultiplyBody interface {
	SetIds(ids []schema.ID)
}
