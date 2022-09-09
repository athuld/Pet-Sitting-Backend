package datasource

import (
	"context"
	"os"
	"pet-sitting-backend/utils/logger"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var (
	Client *pgx.Conn
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
        logger.Error.Println(envErr)
	}

	var err error
	Client, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	logger.Info.Println("Database Connected Successfully")
}
