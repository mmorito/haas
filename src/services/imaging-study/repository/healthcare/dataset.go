package healthcare

import (
	"context"
	"fmt"

	healthcareapi "google.golang.org/api/healthcare/v1beta1"

	"github.com/mnes/haas/src/services/imaging-study/config"
	"github.com/mnes/haas/src/services/imaging-study/driver/healthcare"
)

type DatasetRepo interface {
	CreateDataset(ctx context.Context, datasetID string) (string, error)
}

type datasetRepo struct {
	healthcare *healthcare.Healthcare
	projectID  string
	location   string
}

func NewDatasetRepo(h *healthcare.Healthcare) DatasetRepo {
	env := config.Env()
	return &datasetRepo{
		healthcare: h,
		projectID:  env.ProjectID,
		location:   env.Location,
	}
}

// CreateDataset create dataset
// The dataset is not always ready to use immediately, instead a long-running operation is returned.
func (r *datasetRepo) CreateDataset(ctx context.Context, datasetID string) (string, error) {

	parent := fmt.Sprintf("projects/%s/locations/%s", r.projectID, r.location)

	resp, err := r.healthcare.DatasetsService().Create(parent, &healthcareapi.Dataset{}).DatasetId(datasetID).Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("Create: %v", err)
	}
	return resp.Name, nil
}
