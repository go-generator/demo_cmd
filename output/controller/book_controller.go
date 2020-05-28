package controller

import (
	. "github.com/common-go/echo"
	"project_name/handler"
	"project_name/model"
	"project_name/search-model"
	"project_name/service"
	"reflect"
)

type BookController struct {
	*GenericController
	*SearchController
}

func NewBookController(BookService service.BookService, logService ActivityLogService) *BookController {
	var BookModel model.Book
	modelType := reflect.TypeOf(BookModel)
	searchModelType := reflect.TypeOf(search_model.BookSM{})
	idNames := BookService.GetIdNames()
	controlModelHandler := handler.NewControlModelHandler(idNames)
	genericController, searchController := NewGenericSearchController(BookService, modelType, controlModelHandler, BookService, searchModelType, nil, logService, true, "")
	return &BookController{GenericController: genericController, SearchController: searchController}
}
