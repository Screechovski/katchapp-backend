package helper

import (
	"log"
	"net/http"
	"os"
)

func IsDev() bool {
	mode, modeExists := os.LookupEnv("MODE")

	return (modeExists && mode == "dev")
}

func HandleError(w http.ResponseWriter, err error, statusCode int, userMessage string) {
	if IsDev() {
		log.Printf("Error [%d]: %v", statusCode, err)
		http.Error(w, err.Error(), statusCode)
	} else {
		log.Printf("Error [%d]: %v", statusCode, err)
		http.Error(w, userMessage, statusCode)
	}
}
