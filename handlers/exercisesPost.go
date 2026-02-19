package handlers

import (
	"fmt"
	"io"
	"katchapp-backend/db"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func convertAndSave(file multipart.File, name, format string) error {
	file.Seek(0, io.SeekStart)
	newName := "images/exercise/" + name + format
	dst, err := os.Create(newName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)

	return err
}

func stringToIntSlice(strSlice []string) ([]int, error) {
	intSlice := make([]int, 0, len(strSlice))

	for _, s := range strSlice {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("invalid integer string: %w", err)
		}
		intSlice = append(intSlice, i)
	}

	return intSlice, nil
}

func ExercisesPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	if handler.Size > 1*1024*1024 { // 1MB in bytes
		http.Error(w, "file must be less than 1MB", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "no name", http.StatusBadRequest)
		return
	} else if len(name) < 4 {
		http.Error(w, "name should > 3", http.StatusBadRequest)
		return
	} else if len(name) > 50 {
		http.Error(w, "name should < 50", http.StatusBadRequest)
		return
	}

	mainGroup := r.FormValue("main-group")
	if mainGroup == "" {
		http.Error(w, "'main-group' is required", http.StatusBadRequest)
		return
	}
	mainGroupInt, err := strconv.Atoi(mainGroup)
	if err != nil {
		http.Error(w, "main-group invalid data", http.StatusBadRequest)
		return
	}

	subgroups := r.MultipartForm.Value["subgroup"]
	subgroupInts, err := stringToIntSlice(subgroups)
	if err != nil {
		http.Error(w, "subgroup invalid data", http.StatusBadRequest)
		return
	}

	subs := []db.SubGroup{}

	for _, sub := range subgroupInts {
		subs = append(subs, db.SubGroup{
			Id:    uint(sub),
			Level: 1,
		})
	}

	id, err := db.SaveExercise(
		name,
		"",
		uint(mainGroupInt),
		subs,
	)

	ext := filepath.Ext(handler.Filename)
	err = convertAndSave(file, fmt.Sprintf("%d", id), ext)
	if err != nil {
		http.Error(w, "error saving image", http.StatusInternalServerError)
		return
	}

	err = db.UpdateImage(id, fmt.Sprintf("%d", id)+ext)
	if err != nil {
		http.Error(w, "error update image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
