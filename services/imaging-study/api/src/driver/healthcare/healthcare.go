package healthcare

import (
	"context"

	healthcare "google.golang.org/api/healthcare/v1beta1"
)

type Healthcare struct {
	HealthcareService *healthcare.Service
}

func NewHealthcare() (*Healthcare, error) {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Healthcare{
		HealthcareService: healthcareService,
	}, nil
}

func (h *Healthcare) DicomStoresService() *healthcare.ProjectsLocationsDatasetsDicomStoresService {
	return h.HealthcareService.Projects.Locations.Datasets.DicomStores
}

func (h *Healthcare) StudiesService() *healthcare.ProjectsLocationsDatasetsDicomStoresStudiesService {
	return h.DicomStoresService().Studies
}

func (h *Healthcare) SeriesService() *healthcare.ProjectsLocationsDatasetsDicomStoresStudiesSeriesService {
	return h.StudiesService().Series
}

func (h *Healthcare) OperationService() *healthcare.ProjectsLocationsDatasetsOperationsService {
	return h.HealthcareService.Projects.Locations.Datasets.Operations
}
