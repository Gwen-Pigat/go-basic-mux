package core

import (
	coreEntity "dates/Modules/Core/Entity"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func UcFirst(data string) string {
	title := cases.Title(language.French)
	shuffledString := title.String(strings.ToLower(string(data)))
	return shuffledString
}

func ReturnData(w http.ResponseWriter, message string, status ...int) {
	w.Header().Set("Content-type", "application/json")
	if len(status) == 0 {
		status = append(status, http.StatusBadRequest)
	}
	w.WriteHeader(status[0])
	jsonData, _ := json.Marshal(coreEntity.Result{
		Message: message,
		Status:  status[0],
	})
	w.Write(jsonData)
}

func IsImage(contentType string) bool {
	validTypes := []string{"image/jpeg", "image.png", "image.gif"}
	for _, t := range validTypes {
		if strings.HasPrefix(contentType, t) {
			return true
		}
	}
	return false
}
