package belajar

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}

func GetOneDoc[T any](db *mongo.Database, collection string, filter bson.M) (doc T, err error) {
	err = db.Collection(collection).FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		return
	}
	return
}



