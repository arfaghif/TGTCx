package handlers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arfaghif/TGTCx/backend/dictionary"
	"github.com/arfaghif/TGTCx/backend/domain/product"
	time_helper "github.com/arfaghif/TGTCx/backend/helpers"
	"github.com/arfaghif/TGTCx/backend/service"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.DeleteProduct(context.Background(), p.ID)
	fmt.Fprintf(w, "success")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var p dictionary.Product

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	product.UpdateProduct(context.Background(), p)

	fmt.Fprintf(w, "success")
}

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	var banner dictionary.Banner

	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	banner.Name = r.FormValue("name")
	banner.Description = r.FormValue("description")
	banner.Tags = strings.Split(r.FormValue("tags"), ",")
	var err error
	banner.StartDate, err = time.Parse(time.RFC3339, r.FormValue("start_date"))
	if err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Bad Request",
			Error: err.Error(),
		}), http.StatusBadRequest)
		return
	}
	banner.EndDate, err = time.Parse(time.RFC3339, r.FormValue("end_date"))
	if err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Bad Request",
			Error: err.Error(),
		}), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Internal Server error",
			Error: err.Error(),
		}), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dir, err := os.Getwd()

	if err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Internal Server error",
			Error: err.Error(),
		}), http.StatusInternalServerError)
		return
	}

	filename := handler.Filename
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	filename = hex.EncodeToString(randBytes) + "_" + filename

	fileLocation := filepath.Join(dir, "files", "banners", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Internal Server error",
			Error: err.Error(),
		}), http.StatusInternalServerError)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Internal Server error",
			Error: err.Error(),
		}), http.StatusInternalServerError)
		return
	}
	banner.ImgPath = fileLocation

	if err := service.UploadBanner(banner); err != nil {
		http.Error(w, time_helper.BuildResponse(dictionary.APIResponse{
			Data:  "Unprocessable Entity",
			Error: err.Error(),
		}), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, _ := json.Marshal(
		dictionary.APIResponse{
			Data: "Success",
		},
	)
	w.Write(resp)
}
