package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/aaron-g-sanchez/PROTOTYPE/PROJECT-ATHENA-PROTO/frontend/text-editor/templates"
)

func main() {
	homePage := templates.Home()

	http.Handle("/", templ.Handler(homePage))

	fmt.Println("Listening on port :3000")

	http.ListenAndServe(":3000", nil)
}
