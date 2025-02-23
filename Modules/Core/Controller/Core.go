package coreController

import (
	core "dates/Core"
	coreEntity "dates/Modules/Core/Entity"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	resp := coreEntity.Result{
		Message: "Response OK",
		Status:  http.StatusOK,
	}
	_, err := json.Marshal(resp)
	if err != nil {
		core.ReturnData(w, "Erreur de formattage JSON")
		return
	}
	core.ReturnData(w, resp.Message, resp.Status)
}

func StringExamples(w http.ResponseWriter, r *http.Request) {
	stringValues := [3]string{"John", "Jack", "Bryan"}
	randomIndex := rand.Intn(len(stringValues))
	runes := []rune(stringValues[randomIndex])
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	response := coreEntity.Result{
		Message: core.UcFirst(string(runes)),
		Status:  http.StatusOK,
	}
	_, err := json.Marshal(response)
	if err != nil {
		core.ReturnData(w, "Erreur de formattage JSON")
		return
	}
	core.ReturnData(w, response.Message, response.Status)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		core.ReturnData(w, "This method is not authorized", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		core.ReturnData(w, err.Error())
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		core.ReturnData(w, err.Error())
		return
	}
	defer file.Close()

	if !core.IsImage(handler.Header.Get("Content-type")) {
		core.ReturnData(w, "The image is not valid")
		return
	}

	uploadsDir := "./uploads/"
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {
		os.Mkdir(uploadsDir, os.ModePerm)
	}
	fmt.Println(filepath.Join(uploadsDir, handler.Filename))

	f, err := os.Create(filepath.Join(uploadsDir, "newfile.webp"))
	if err != nil {
		core.ReturnData(w, err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		core.ReturnData(w, err.Error())
		return
	}
	core.ReturnData(w, "File is uploaded", http.StatusOK)
}
