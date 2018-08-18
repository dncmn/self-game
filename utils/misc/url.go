package misc

import (
	"net/http"
	"net/url"
)

func GetUrlValues(w http.ResponseWriter, r *http.Request) (err error, values url.Values) {
	err = r.ParseForm()
	if err != nil {
		return
	}
	values = r.Form
	return
}
