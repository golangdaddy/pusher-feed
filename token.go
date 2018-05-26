package pusherfeed

import (
	"fmt"
	"time"
	"encoding/base64"
	"encoding/json"
	//
	"github.com/dgrijalva/jwt-go"
)

/*
"app": instance_id,
"iss": "api_keys/<key_id>",
"iat": issued_at,
"exp": expiry_time,
"sub": user_id,
"clients": {
  "permission": {
	"action": action,
	"path": path
  }
}
*/


type Claims struct {
	App string `json:"app"`
	Iss string `json:"iss"`
	Iat string `json:"iat"`
	Exp string `json:"exp"`
	Feeds *Feeds `json:"clients"`
	jwt.StandardClaims
}

type Feeds struct {
	Permission *Permission `json:"permission"`
}

type Permission struct {
	Action string `json:"action"`
	Path string `json:"path"`
}

func (client *Client) NewToken() string {

	_, err := base64.StdEncoding.DecodeString(client.secretKey)
	if err != nil {
		panic(err)
	}

	claims := jwt.MapClaims{
		"app": client.instanceLocator,
		"iss": "api_keys/"+client.keyId,
		"iat": time.Now().UTC().Unix(),
		"exp": time.Now().UTC().Add(24 * time.Hour).Unix(),
		"feeds": &Feeds{
			&Permission{
				Action: "*",
				Path: "*",
			},
		},
	}

	b, _ := json.Marshal(claims)
	fmt.Println(string(b))

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	ss, err := token.SignedString([]byte(client.secretKey))
	if err != nil {
		panic(err)
	}

	fmt.Println(ss)

	return ss
}
