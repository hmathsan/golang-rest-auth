package catchAllHandler

import "net/http"

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
