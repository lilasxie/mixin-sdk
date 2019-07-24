package main

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"log"

	"github.com/fox-one/mixin-sdk/mixin"
	jsoniter "github.com/json-iterator/go"
)

func printJSON(prefix string, item interface{}) {
	msg, err := jsoniter.MarshalToString(item)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(prefix, msg)
}

func main() {
	user := &mixin.User{
		UserID:    ClientID,
		SessionID: SessionID,
		PINToken:  PINToken,
	}

	block, _ := pem.Decode([]byte(SessionKey))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Panicln(err)
	}
	user.SetPrivateKey(privateKey)

	ctx := context.Background()

	p := "123456"
	assetID := "965e5c6e-434c-3fa9-b780-c50f43cd955c"

	publicKey := doAsset(ctx, user, assetID)
	u := doCreateUser(ctx, user, p)

	{
		pkce := "uPSuDZM_9QL_M5fbdSSo01LJjwWpPBbwIDg10x02RyU"
		code := "1357994f-619a-4a3a-bcc2-00eb4dd87d13"
		a, err := u.GetAuthorization(ctx, code)
		log.Println(a, err)
		a1, err := u.Authorize(ctx, a.AuthorizationID, a.Scopes)
		log.Println(a1, err)
		user.GetOAuthToken(ctx, a1.AuthorizationCode, a1.App.AppID, "", pkce)
	}

	doAssetFee(ctx, u)

	publicKey1 := doAsset(ctx, u, assetID)

	doAssets(ctx, u)

	doTransfer(ctx, user, assetID, u.UserID, "0.1", "ping", PIN)
	snap := doTransfer(ctx, u, assetID, user.UserID, "0.1", "pong", p)

	doWithdraw(ctx, user, assetID, publicKey1, "0.1", "ping", PIN)
	doWithdraw(ctx, u, assetID, publicKey, "0.1", "pong", p)

	doReadNetwork(ctx, u)

	doReadSnapshot(ctx, u, snap.SnapshotID)

	doReadTransfer(ctx, u, snap.TraceID)

	doReadExternal(ctx, u)
}
