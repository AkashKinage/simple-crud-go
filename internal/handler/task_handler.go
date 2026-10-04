package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/AkashKinage/simple-crud-go/internal/model"
	"github.com/AkashKinage/simple-crud-go/internal/repository"
	"github.com/AkashKinage/simple-crud-go/internal/service"
)

type TaskHandler struct {
	service service.TaskService
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
    Priority    string `json:"priority"`
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return	
	}

	task := &model.Task{
		Title: req.Title,
		Description: req.Description,
		Status: req.Status,
		Priority: req.Priority,
	}

	created, err := h.service.Create(r.Context(), task)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidTitle), errors.Is(err, service.ErrInvalidStatus), errors.Is(err, service.ErrInvalidPriority):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "something went wrong")
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *TaskHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		writeError(w, http.StatusUnprocessableEntity, "id is not present")
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, "something went wrong while parsing id")
		return
	}

	task, err := h.service.GetByID(r.Context(), intId)

	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, "task not found")
		} else {
			writeError(w, http.StatusInternalServerError, "something went wrong")
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	queryStatus := query.Get("status")
	queryPriority := query.Get("priority")
	queryPage := query.Get("page")
	queryPerPage := query.Get("per_page")

	var page, perPage int

	if queryPage != "" {
		p, err := strconv.Atoi(queryPage)
		if err != nil {
			writeError(w, http.StatusBadRequest, "error parsing page param")
			return
		}

		page = p
	}

	if queryPerPage != "" {
		p, err := strconv.Atoi(queryPerPage)
		if err != nil {
			writeError(w, http.StatusBadRequest, "error parsing per_page param")
			return
		}

		perPage = p
	}
	

	filter := repository.TaskFilter{
		Status: queryStatus,
		Priority: queryPriority,
		Page: page,
		PerPage: perPage,
	}

	tasks, err := h.service.List(r.Context(), filter)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidStatus), errors.Is(err, service.ErrInvalidPriority), errors.Is(err, service.ErrInvalidPerPage):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "something went wrong")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}