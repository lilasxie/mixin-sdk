package mixin

import (
	"context"
	"encoding/json"

	mixin_sdk "github.com/lilasxie/mixin-sdk"
)

// RequestWithPIN sign pin and request
func (user *User) RequestWithPIN(ctx context.Context, method, uri string, payload map[string]interface{}, pin string) ([]byte, error) {
	if payload == nil {
		payload = map[string]interface{}{}
	}
	pinToken, err := user.signPIN(pin)
	if err != nil {
		return nil, err
	}
	payload["pin"] = pinToken

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return user.Request(ctx, method, uri, data)
}

// Request sign and request
func (user *User) Request(ctx context.Context, method, uri string, payload []byte) ([]byte, error) {
	ctx = mixin_sdk.WithAuth(ctx, user)
	resp, err := mixin_sdk.Request(ctx).SetBody(payload).Execute(method, uri)
	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
