package database

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = viper.New()
var db *mongo.Database

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func InitDatabase() *mongo.Database {
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	uri := config.GetString("MONGODB_URI")
	dbname := config.GetString("MONGODB_DBNAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	db = client.Database(dbname)

	return db
}

func CloseDatabase() error {
	return db.Client().Disconnect(context.Background())
}
