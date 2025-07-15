package account

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB(dataBaseURL string) {
	var err error
	DB, err = pgxpool.New(context.Background(), dataBaseURL)
	if err != nil {
		log.Fatal("Error creating connection pool:", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Database connection established successfully!")
}
