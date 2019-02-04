package siprocket

import (
	"fmt"
	"strings"
)

type Auth struct {
	Username string
	Realm    string
	Nonce    string
	Opaque   string
	URI      string
	Response string
	Src      string
}

func AuthParser(authLine string, auth *Auth) {
	values := strings.Split(authLine, ",")
	values = strings.Split(strings.Join(values, ""), "=")
	values = strings.Split(strings.Join(values, ""), "\"")
	auth.Src = authLine
	for index := range values {
		line := strings.TrimSpace(values[index])
		switch line {
		case "Digest username":
			auth.Username = values[index+1]
		case "realm":
			auth.Realm = values[index+1]
		case "nonce":
			auth.Nonce = values[index+1]
		case "opaque":
			auth.Opaque = values[index+1]
		case "uri":
			auth.URI = values[index+1]
		case "response":
			auth.Response = values[index+1]
		default:
			continue
		}
	}
}
