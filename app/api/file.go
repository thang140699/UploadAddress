package api

import (
	ultilitiesFS "WeddingUtilities/utilities/firebase"
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(30 * 1024 * 1024)
	files := r.MultipartForm.File["file"]
	var arrayFileName = []string{}

	for _, file := range files {

		fmt.Println("file Name:", file.Filename)
		fmt.Println("File Size :", file.Size)
		fmt.Println("File Type:", file.Header.Get("Content-Type"))
		fmt.Println("-------------------")
		f, _ := file.Open()
		arrayFileName = append(arrayFileName, file.Filename)
		reader := bufio.NewReader(f)
		buf := &bytes.Buffer{}
		buf.ReadFrom(reader)
		data := buf.Bytes()
		err := ultilitiesFS.Upload(data, file.Filename, file.Header.Get("Content-Type"))
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ResponseBody{
				Message: "Cannot upload file" + err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
	}
	WriteJSON(w, http.StatusOK, ResponseBody{
		Message: "upload successfully",
		Code:    200,
		Data: map[string][]string{
			"data": arrayFileName,
		},
	})
}
