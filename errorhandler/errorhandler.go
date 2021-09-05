package errorhandler

import (
	"fmt"
	"net/http"
	"github.com/marvel/model"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		var response model.HTTPError404
		response.Code = status
		response.Message = "Not Found"
		fmt.Fprint(w, response)
	}
}
