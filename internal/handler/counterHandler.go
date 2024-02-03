package handler

import (
	"net/http"

	"github.com/abhilash26/gohyper/internal/model"
	"github.com/abhilash26/gohyper/internal/storage"
	"github.com/abhilash26/gohyper/internal/view"
)

func CounterAdd(w http.ResponseWriter, r *http.Request) {
  counter, _ := model.GetCounter(storage.DB)
  counter++;
	model.UpdateCounterValue(storage.DB, counter)
	templates := []string{
		"component/counter",
	}
	view.RenderTemplate(w, templates, "counter", counter)
}

func CounterSubtract(w http.ResponseWriter, r *http.Request) {
  counter, _ := model.GetCounter(storage.DB)
  counter--;
	model.UpdateCounterValue(storage.DB, counter)
	templates := []string{
		"component/counter",
	}
	view.RenderTemplate(w, templates, "counter", counter)
}
