package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var jsonContentType = []string{"application/json;charset=utf-8"}

const (
	MESSAGE_INTERNAL_SERVER_ERROR = "Internal Server Error"
	MESSAGE_FILE_EXCEED_LIMIT     = "File Exceeded Limit"
)

type ResponseBody struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, code int, obj interface{}) error {
	w.WriteHeader(code)
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(jsonBytes)
	return nil
}
func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
func BindJSON(r *http.Request, obj interface{}) error {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return err

	}
	return json.Unmarshal(b, obj)
}
