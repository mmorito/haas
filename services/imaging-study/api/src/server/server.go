package server

import (
	"log"

	"github.com/mnes/haas/services/imaging-study/openapi"
	healthcare "github.com/mnes/haas/services/imaging-study/src/driver/healthcare"
	healthcareRepo "github.com/mnes/haas/services/imaging-study/src/repository/healthcare"
	"github.com/mnes/haas/services/imaging-study/src/usecase/dicomweb"
)

type server struct {
	dicomwebUsecase dicomweb.Usecase
}

func NewServer() openapi.ServerInterface {
	h, err := healthcare.NewHealthcare()
	if err != nil {
		log.Fatal(err)
	}

	// repository
	dicomStoreRepo := healthcareRepo.NewDicomStoreRepo(h)

	// usecase
	dicomwebUsecase := dicomweb.NewUsecase(dicomStoreRepo)

	return &server{
		dicomwebUsecase: dicomwebUsecase,
	}
}
