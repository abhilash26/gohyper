package handler

import (
	"log"
	"net/http"

	"github.com/abhilash26/gohyper/internal/model"
	"github.com/abhilash26/gohyper/internal/storage"
	"github.com/abhilash26/gohyper/internal/view"
)

type PageData struct {
	PageName string
	Counter  int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	model.InitCounterTable(storage.DB)
	countervalue, err := model.GetCounter(storage.DB)
	if err != nil {
		log.Fatal("The counter could not be retrieved: ", err)
	}
	data := PageData{
		Counter: countervalue,
	}
	templates := []string{
		"partial/header",
		"partial/footer",
		"component/counter",
		"page/index",
	}
	view.RenderTemplate(w, templates, "index", data)
}
