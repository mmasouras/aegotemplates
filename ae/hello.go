package hello

import (
	"aegotemplates/homepage"
	"net/http"
)

func init() {
	http.HandleFunc("/", homepage.Homepage)
}
