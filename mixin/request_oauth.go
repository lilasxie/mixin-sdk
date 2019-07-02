package mixin

import (
	"context"
	"encoding/json"

	"github.com/fox-one/mixin-sdk/utils"
	jsoniter "github.com/json-iterator/go"
)

// GetAuthorization get authorization by code
func (user User) GetAuthorization(ctx context.Context, code string) (*Authorization, error) {
	data, err := user.Request(ctx, "GET", "/codes/"+code, nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Authorization *Authorization `json:"data,omitempty"`
		Error         *Error         `json:"error,omitempty"`
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, requestError(err)
	} else if resp.Error != nil {
		return nil, resp.Error
	}

	return resp.Authorization, nil
}

// Authorize authorize authorization_id
func (user User) Authorize(ctx context.Context, authorizationID string, scopes []string) (*Authorization, error) {
	params := map[string]interface{}{
		"authorization_id": authorizationID,
		"scopes":           scopes,
	}
	bts, _ := jsoniter.Marshal(params)
	data, err := user.Request(ctx, "POST", "/oauth/authorize", bts)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Authorization *Authorization `json:"data,omitempty"`
		Error         *Error         `json:"error,omitempty"`
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, requestError(err)
	}

	return resp.Authorization, resp.Error
}

// GetOAuthToken get oauth token with code
func GetOAuthToken(ctx context.Context, code, appID, appSecret string, pkce ...string) (string, error) {
	params := map[string]interface{}{
		"client_id": appID,
		"code":      code,
	}
	if len(appSecret) > 0 {
		params["client_secret"] = appSecret
	} else if len(pkce) > 0 && len(pkce[0]) > 0 {
		params["code_verifier"] = pkce[0]
	}
	bts, _ := jsoniter.Marshal(params)
	url := "https://api.mixin.one/oauth/token"
	method := "POST"

	data, err := utils.SendRequest(ctx, url, method, string(bts), "Content-Type", "application/json").Bytes()
	if err != nil {
		return "", err
	}

	var resp struct {
		Data *struct {
			AccessToken string `json:"access_token"`
			Scope       string `json:"scope"`
		} `json:"data,omitempty"`
		Error *Error `json:"error,omitempty"`
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return "", requestError(err)
	} else if resp.Error != nil {
		return "", resp.Error
	}

	return resp.Data.AccessToken, nil
}
