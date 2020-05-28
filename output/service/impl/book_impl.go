package impl

import (
	. "github.com/common-go/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	. "project_name/model"
	"reflect"
)

type BookServiceImpl struct {
	database   *mongo.Database
	collection *mongo.Collection
	*DefaultGenericService
	*DefaultSearchService
}

func NewBookServiceImpl(db *mongo.Database, searchResultBuilder SearchResultBuilder) *BookServiceImpl {
	var model Book
	modelType := reflect.TypeOf(model)
	collection := "book"
	mongoService, searchService := NewMongoGenericSearchService(db, modelType, collection, searchResultBuilder, false, "")
	return &BookServiceImpl{db, db.Collection(collection), mongoService, searchService}
}
