package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) GetHandlerbyId(w http.ResponseWriter, r *http.Request) {
	paramsWorkourtId := chi.URLParam(r, "id")

	if paramsWorkourtId == "" {
		http.NotFound(w, r)

		return
	}

	workoutId, err := strconv.ParseInt(paramsWorkourtId, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this is our workout Id %d\n", workoutId)
}

func (wh *WorkoutHandler) Handlercreateworkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "we have cretaed workout \n")
}
