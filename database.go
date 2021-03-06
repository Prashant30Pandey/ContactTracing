package databse

import (
	"ContactTracing/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(conf config.MongoConfiguration) *mongo.Database {
	connection := options.Client().ApplyURI(conf.Server)
	client, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		panic(err)
		return client.Database(cont.Database)
	}
}
