package http

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "PONG"); err != nil {
		msg := fmt.Sprintf("coulfn't ping? %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}
