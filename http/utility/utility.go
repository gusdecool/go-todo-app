package utility

import "net/http"

func HandleErrorResponse(err error, response http.ResponseWriter) {
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(err.Error()))
}
