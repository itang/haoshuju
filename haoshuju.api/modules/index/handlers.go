package index

import (
	"net/http"
)

func IndexHandler(resp http.ResponseWriter, req *http.Request) {
	http.Redirect(resp, req, "/public/info.html", http.StatusFound)
}
