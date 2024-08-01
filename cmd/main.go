package main

import (
	"fmt"
	"log"

	"github.com/Sharofiddin701/medium-clone-simple/config"
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

	log.Println("Postgres Connetion succesfully done !")
	fmt.Println("Postgres Connetion succesfully done !")

	_ = psqlConn

}