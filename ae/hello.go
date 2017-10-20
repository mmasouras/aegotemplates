package hello

import (
    "net/http"
    "aegotemplates/homepage"
)

func init() {
    http.HandleFunc("/", homepage.Homepage)
}
