package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aaron-g-sanchez/PROTOTYPE/PROJECT-ATHENA-PROTO/frontend/auth/templates"
	"github.com/joho/godotenv"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

var (
	domain      string
	key         string
	clientId    string
	redirectUri string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env", err)
	}

	domain = os.Getenv("DOMAIN")
	key = os.Getenv("KEY")
	clientId = os.Getenv("CLIENT_ID")
	redirectUri = os.Getenv("REDIRECT_URI")
}

func main() {

	ctx := context.Background()

	authN, err := authentication.New(ctx, zitadel.New(domain), key, openid.DefaultAuthentication(clientId, redirectUri, key))
	if err != nil {
		log.Fatal("zitadel sdk could not be initialized", err)
	}

	mw := authentication.Middleware(authN)

	router := http.NewServeMux()

	router.Handle("/auth/", authN)

	router.Handle("/profile", mw.RequireAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authCtx := mw.Context(req.Context())
		data, err := json.MarshalIndent(authCtx.UserInfo, "", "")
		if err != nil {
			log.Fatal("error marshalling profile response", err)
		}
		component := templates.Profile(string(data))
		component.Render(req.Context(), w)
	})))

	router.Handle("/", mw.CheckAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if authentication.IsAuthenticated(req.Context()) {
			http.Redirect(w, req, "/profile", http.StatusFound)
		}
		homeComponent := templates.Home()
		homeComponent.Render(req.Context(), w)
	})))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", router)
}
