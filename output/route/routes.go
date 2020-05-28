package route

import (
	"github.com/common-go/mongo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"project_name/config"
)

type EvaRoutes struct {
	Router *echo.Echo
}

func NewProjectNameRoutes(e *echo.Echo, mongoConfig mongo.MongoConfig) (*EvaRoutes, error) {
	applicationContext, err := config.NewApplicationContext(mongoConfig)
	if err != nil {
		return nil, err
	}

	//middle for all routes
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bookController := applicationContext.BookController
	bookPath := "/project_name/book"
	e.GET(bookPath, bookController.GetAll())
	e.POST(bookPath, bookController.Insert())
	e.GET(bookPath+"/:id", bookController.GetById())
	e.POST(bookPath+"/search", bookController.Search())
	e.PUT(bookPath+"/:id", bookController.Update())

	bookCategoryController := applicationContext.BookCategoryController
	bookCategoryPath := "/project_name/bookCategory"
	e.GET(bookCategoryPath, bookCategoryController.GetAll())
	e.POST(bookCategoryPath, bookCategoryController.Insert())
	e.GET(bookCategoryPath+"/:id", bookCategoryController.GetById())
	e.POST(bookCategoryPath+"/search", bookCategoryController.Search())
	e.PUT(bookCategoryPath+"/:id", bookCategoryController.Update())

	return &EvaRoutes{e}, nil
}
