package server

import (
	"log"

	healthcare "github.com/mnes/haas/src/services/imaging-study/driver/healthcare"
	"github.com/mnes/haas/src/services/imaging-study/openapi"
	healthcareRepo "github.com/mnes/haas/src/services/imaging-study/repository/healthcare"
	"github.com/mnes/haas/src/services/imaging-study/usecase/dicomweb"
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
