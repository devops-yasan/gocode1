package handlers

import (
"encoding/json"
"net/http"
"strings"

"go-rest-api/models"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
return &UserHandler{}
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
switch r.Method {
case http.MethodGet:
writeJSON(w, http.StatusOK, []models.User{{ID: "1", Name: "Jane Doe", Email: "jane@example.com"}})
case http.MethodPost:
writeJSON(w, http.StatusCreated, map[string]string{"message": "create user stub"})
default:
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
}

func (h *UserHandler) UserByID(w http.ResponseWriter, r *http.Request) {
id := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")
if id == "" {
http.NotFound(w, r)
return
}

switch r.Method {
case http.MethodGet:
writeJSON(w, http.StatusOK, models.User{ID: id, Name: "Sample User", Email: "sample@example.com"})
case http.MethodPut:
writeJSON(w, http.StatusOK, map[string]string{"message": "update user stub", "id": id})
case http.MethodDelete:
writeJSON(w, http.StatusOK, map[string]string{"message": "delete user stub", "id": id})
default:
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(status)
_ = json.NewEncoder(w).Encode(data)
}
