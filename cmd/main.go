package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sharofiddin701/medium-clone-simple/config"
	"github.com/Sharofiddin701/medium-clone-simple/storage"
	"github.com/Sharofiddin701/medium-clone-simple/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	log.Println("Postgres Connection successfully done!")
	fmt.Println("Postgres Connection successfully done!")

	strg := storage.NewStorage(psqlConn)

	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	user, err := strg.User().Create(context.TODO(), &repo.User{
		ID:        id.String(),
		FirstName: "Sharofiddin",
		LastName:  "Bobomurodov",
		Email:     "someone@gmail.com",
		Password:  "password123",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
