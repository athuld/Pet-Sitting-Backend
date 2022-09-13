package datasource

import (
	"context"
	"os"
	"pet-sitting-backend/utils/logger"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	Client *pgxpool.Pool
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		logger.Error.Println(envErr)
	}

	var err error
	Client, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	logger.Info.Println("Database Connected Successfully")
}
