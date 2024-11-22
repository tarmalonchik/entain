package response

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

type SuccessfulResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
}

func JSON200(w http.ResponseWriter) {
	data := "OK"
	RenderJSON(w, http.StatusOK, data)
}

func JSON500(w http.ResponseWriter) {
	JSON500txt(w, "internal server error")
}

func JSON500txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusInternalServerError,
	}
	RenderJSON(w, http.StatusInternalServerError, data)
}

func JSON409txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusConflict,
	}
	RenderJSON(w, http.StatusConflict, data)
}

func JSON400(w http.ResponseWriter) {
	JSON400txt(w, "bad request")
}

func JSON400txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusBadRequest,
	}
	RenderJSON(w, http.StatusBadRequest, data)
}

func JSON404txt(w http.ResponseWriter, errText string) {
	data := ErrorResponse{
		Error:      errText,
		StatusCode: http.StatusNotFound,
	}
	RenderJSON(w, http.StatusNotFound, data)
}

func RenderJSON(w http.ResponseWriter, status int, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.WithField("data", data).Error("Could not encode response")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		log.WithError(err).Error("RenderJSON w.Write error")
	}
}
