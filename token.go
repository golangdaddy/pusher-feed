package pusherfeed

import (
	"fmt"
	"time"
	//
	"github.com/dgrijalva/jwt-go"
)

/*
"app": instance_id,
"iss": "api_keys/<key_id>",
"iat": issued_at,
"exp": expiry_time,
"sub": user_id,
"feeds": {
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
	Sub string `json:"sub"`
	Feeds *Feeds `json:"feeds"`
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

	sigKey := []byte(client.keySecret)

	// Create the Claims
	claims := Claims{
		client.instance,
		"api_keys/"+client.keyId,
		fmt.Sprintf("%v", time.Now().UTC().Unix()),
		fmt.Sprintf("%v", time.Now().UTC().Add(time.Hour).Unix()),
		"user_id",
		&Feeds{
			&Permission{
				"action",
				"path",
			},
		},
		jwt.StandardClaims{
	        ExpiresAt: 15000,
	        Issuer:    "test",
	    },
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(sigKey)
	if err != nil {
		panic(err)
	}

	return ss
}
