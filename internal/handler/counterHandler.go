package handler

import (
	"fmt"
	"net/http"

	"github.com/abhilash26/gohyper/internal/model"
	"github.com/abhilash26/gohyper/internal/storage"
)

func CounterAdd(w http.ResponseWriter, r *http.Request) {
	counter, _ := model.GetCounter(storage.DB)
	counter++
	model.UpdateCounterValue(storage.DB, counter)
	fmt.Fprintf(w, "%d", counter)
}

func CounterSubtract(w http.ResponseWriter, r *http.Request) {
	counter, _ := model.GetCounter(storage.DB)
	counter--
	model.UpdateCounterValue(storage.DB, counter)
	fmt.Fprintf(w, "%d", counter)
}
