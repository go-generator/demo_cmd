package controller

import (
	. "github.com/common-go/echo"
	"project_name/handler"
	"project_name/model"
	"project_name/search-model"
	"project_name/service"
	"reflect"
)

type BookCategoryController struct {
	*GenericController
	*SearchController
}

func NewBookCategoryController(BookCategoryService service.BookCategoryService, logService ActivityLogService) *BookCategoryController {
	var BookCategoryModel model.BookCategory
	modelType := reflect.TypeOf(BookCategoryModel)
	searchModelType := reflect.TypeOf(search_model.BookCategorySM{})
	idNames := BookCategoryService.GetIdNames()
	controlModelHandler := handler.NewControlModelHandler(idNames)
	genericController, searchController := NewGenericSearchController(BookCategoryService, modelType, controlModelHandler, BookCategoryService, searchModelType, nil, logService, true, "")
	return &BookCategoryController{GenericController: genericController, SearchController: searchController}
}
