package service

import (
	"github.com/common-go/search"
	"github.com/common-go/service"
)

type BookService interface {
	service.GenericService
	search.SearchService
}
