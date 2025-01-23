package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type UploadHandler struct {
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (u *UploadHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/upload/{id}", u.uploadFile).Methods("POST")
}

func (u *UploadHandler) uploadFile(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//@todo create directory if doesn't exist
	filePath := fmt.Sprintf("./resumes/%s", userID)
	dst, err := os.Create(filepath.Join(filePath, header.Filename))
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
		return
	}

}
