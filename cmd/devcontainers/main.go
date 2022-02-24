package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/italolelis/devcontainers/pkg/app/storage/postgres"
	"github.com/italolelis/devcontainers/pkg/app/user"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "postgres://admin:qwerty123@localhost:5432/admin-db?sslmode=disable")
	if err != nil {
		log.Fatalf("could not connect to the database: %s", err)
	}
	defer db.Close()

	r := postgres.NewUserRepository(db)

	http.HandleFunc("/", handleUsers(ctx, r))

	http.ListenAndServe(":8080", nil)
}

func handleUsers(ctx context.Context, r user.Repository) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		users, err := r.GetAll(ctx)
		if err != nil {
			http.Error(w, "could not fetch users", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")

		e := json.NewEncoder(w)
		if err := e.Encode(users); err != nil {
			http.Error(w, "could not marshal users", http.StatusInternalServerError)
			return
		}
	}
}
