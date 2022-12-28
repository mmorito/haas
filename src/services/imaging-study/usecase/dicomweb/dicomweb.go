package dicomweb

import "github.com/mnes/haas/src/services/imaging-study/repository/healthcare"

type Usecase interface {
	StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error
}

type usecase struct {
	dicomStoreRepo healthcare.DicomStoreRepo
}

func NewUsecase(dicomStoreRepo healthcare.DicomStoreRepo) Usecase {
	return &usecase{
		dicomStoreRepo: dicomStoreRepo,
	}
}

func (u *usecase) StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error {
	return u.dicomStoreRepo.StoreInstance(dicomData, datasetID, dicomStoreID, dicomWebPath)
}
