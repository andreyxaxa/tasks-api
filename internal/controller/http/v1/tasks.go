package v1

import (
	"encoding/json"
	"net/http"

	"github.com/andreyxaxa/tasks-api/internal/entity"
	"github.com/gorilla/mux"
)

func (h *V1) create(w http.ResponseWriter, r *http.Request) {
	var t entity.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task := h.t.CreateTask(t)
	w.WriteHeader(http.StatusCreated)

	resp := map[string]interface{}{
		"id":       task.ID,
		"name":     task.Name,
		"status":   task.Status,
		"created":  task.CreatedAt,
		"duration": task.Duration,
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *V1) get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id != "" {
		t, err := h.t.GetTask(id)
		if err != nil {
			if err.Error() == "task not found" {
				http.Error(w, "task not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		resp := map[string]interface{}{
			"status":     t.Status,
			"created_at": t.CreatedAt,
			"in_work":    t.InWork,
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func (h *V1) delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.t.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
