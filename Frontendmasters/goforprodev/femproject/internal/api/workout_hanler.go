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

	workout, err := wh.WorkoutStore.GetWorkoutByID(workoutId)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to fetch the workout", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workout)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdWorkout)

}

func (wh *WorkoutHandler) HandleUpdateWorkoutById(w http.ResponseWriter, r *http.Request) {

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

	existingWorkout, err := wh.WorkoutStore.GetWorkoutByID(workoutId)
	if err != nil {
		http.Error(w, "failed to fetch workout", http.StatusInternalServerError)
		return
	}

	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	var updateWorkoutRequest struct {
		Title           *string               `json:"title"`
		Description     *string               `json:"description"`
		DurationMinutes *int                  `json:"duration_minutes"`
		CaloriesBurned  *int                  `json:"calories_burned"`
		Entries         []stores.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}

	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}

	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}

	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}

	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	err = wh.WorkoutStore.UpdateWorkout(existingWorkout)
	if err != nil {
		fmt.Println("update workout error", err)
		http.Error(w, "failed to update the workout", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingWorkout)

}
