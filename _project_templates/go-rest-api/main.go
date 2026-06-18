package main

import (
"encoding/json"
"log"
"net/http"

"go-rest-api/handlers"
)

func main() {
mux := http.NewServeMux()
mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}
w.Header().Set("Content-Type", "application/json")
_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})

userHandler := handlers.NewUserHandler()
mux.HandleFunc("/api/v1/users", userHandler.Users)
mux.HandleFunc("/api/v1/users/", userHandler.UserByID)

server := &http.Server{
Addr:    ":8080",
Handler: mux,
}

log.Println("go-rest-api listening on :8080")
if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
log.Fatalf("server error: %v", err)
}
}
