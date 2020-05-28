package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/common-go/config"
	"github.com/common-go/echo"

	c "project_name/config"
	. "project_name/route"
)

func main() {
	parentPath := "project_name"
	resource := "resource"
	env := os.Getenv("ENV")
	var conf c.Root
	config.LoadConfig(parentPath, resource, env, &conf, "application")
	//resourceMap := server.LoadMap(parentPath, resource, env, "resource_map")
	log.Println(" host ", conf)
	e := echo.New()
	route, err := NewProjectNameRoutes(e, conf)
	if err != nil {
		panic(fmt.Errorf("create route failed"))
	}

	route.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))
	server.StartServer(route.Router, conf.Server)
}
