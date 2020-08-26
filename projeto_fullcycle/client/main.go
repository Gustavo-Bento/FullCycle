package main

import (
	"context"
	oidc "github.com/coreos/go-oidc"
)

var(
	clientId="app"
	clientSecret="d6789a59-c2d4-42a3-83ca-5872f3641048"
)

func main(){
	ctx : = context.Background()	

	provider, err:=oidc.newProvider(ctx, issuer:"http://localhost:8080/auth/realms/demo")
	if err != nil{
		log.Fatalf(err)
	}

	config :=ouath2.Config{
		ClientId:			clientID,
		ClientSecret:	clientSecret,
		Endpoint:			provider.Endpoint(),
		ReadirectURL: "http://localhost:8081/auth/callback",
		Scopes: 			[]strings{oidc, ScopesOpenID, "profile", "email","roles"},
	}

	state := "magica"

	http.HandleFunc(pattern: "/", func(write http.ResponseWriter, resquest *http.Request){
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(addr:":8081",handler: nil))
}
