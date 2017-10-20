package homepage

import (
	"html/template"
    "net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("../homepage/homepage.html")).ExecuteTemplate(w, "base", nil)
 }