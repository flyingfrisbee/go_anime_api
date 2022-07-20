package db

import (
	"GithubRepository/go_anime_api/model"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StartConnectionToMongoDB() {
	ctx, cancel := CreateContext()
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MongoDBURI")))
	if err != nil {
		log.Println(err)
		return
	}

	MongoClient = mongoClient
}

func GetEpisodes(internalID string) []model.Episode {
	coll := MongoClient.Database("episodes").Collection(internalID)

	ctx, cancel := CreateContext()
	defer cancel()

	filterCursor, _ := coll.Find(ctx, bson.M{})
	defer filterCursor.Close(ctx)

	var episodes []model.Episode
	filterCursor.All(ctx, &episodes)

	return episodes
}

func InsertBulkEpisodes(internalID string, episodes []model.Episode) {
	coll := MongoClient.Database("episodes").Collection(internalID)

	convertedEpisodes := convertToInterfaces(episodes)

	ctx, cancel := CreateContext()
	defer cancel()

	_, err := coll.InsertMany(ctx, convertedEpisodes)
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(len(result.InsertedIDs))
}

// returns -1 when no data detected, -2 if document not exist
func CheckIfContainNewData(internalID string, inputLength int) int {
	coll := MongoClient.Database("episodes").Collection(internalID)
	ctx, cancel := CreateContext()
	defer cancel()

	count, _ := coll.CountDocuments(ctx, bson.M{})
	if count == 0 {
		return -2
	}

	if count < int64(inputLength) {
		return inputLength - int(count)
	}

	return -1
}

func convertToInterfaces(episodes []model.Episode) []interface{} {
	length := len(episodes)
	interfaces := make([]interface{}, length)

	for i, j := 0, (length - 1); i <= j; i, j = i+1, j-1 {
		interfaces[i] = episodes[i]
		interfaces[j] = episodes[j]
	}

	return interfaces
}
