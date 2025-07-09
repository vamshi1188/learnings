package api

import (
	"encoding/json"
	"femproject/internal/stores"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
	WorkoutStore stores.WorkoutStore
}

func NewWorkoutHandler(WorkoutStore stores.WorkoutStore) *WorkoutHandler {
	return &WorkoutHandler{
		WorkoutStore: WorkoutStore,
	}

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
	var workouts stores.Workout

	err := json.NewDecoder(r.Body).Decode(&workouts)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create the workout", http.StatusInternalServerError)
		return
	}

	createdWorkout, err := wh.WorkoutStore.CreateWorkout(&workouts)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create the workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)

}
