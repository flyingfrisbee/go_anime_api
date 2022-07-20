package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Conn        *pgx.Conn
	MongoClient *mongo.Client
)

func StartConnectionToDB() {
	StartConnectionToPostgre()
	StartConnectionToMongoDB()
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	return ctx, cancel
}
