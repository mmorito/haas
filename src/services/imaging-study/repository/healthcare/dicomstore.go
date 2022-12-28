package healthcare

import (
	"bytes"
	"fmt"
	"io"

	"go.uber.org/zap"
	healthcareapi "google.golang.org/api/healthcare/v1beta1"

	"github.com/mnes/haas/src/services/imaging-study/config"
	"github.com/mnes/haas/src/services/imaging-study/driver/healthcare"
)

type DicomStoreRepo interface {
	CreateDICOMStore(datasetID, dicomStoreID string) error
	StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error
}

type dicomStoreRepo struct {
	healthcare *healthcare.Healthcare
	projectID  string
	location   string
}

func NewDicomStoreRepo(h *healthcare.Healthcare) DicomStoreRepo {
	env := config.Env()
	return &dicomStoreRepo{
		healthcare: h,
		projectID:  env.ProjectID,
		location:   env.Location,
	}
}

func (r *dicomStoreRepo) CreateDICOMStore(datasetID, dicomStoreID string) error {

	parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", r.projectID, r.location, datasetID)

	resp, err := r.healthcare.DicomStoresService().Create(parent, &healthcareapi.DicomStore{}).DicomStoreId(dicomStoreID).Do()
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}

	zap.L().Info("Created DICOM store", zap.String("dicomStoreID", resp.Name))
	return nil
}

// StoreInstance stores DICOM inctance.
func (r *dicomStoreRepo) StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error {

	parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/dicomStores/%s", r.projectID, r.location, datasetID, dicomStoreID)

	call := r.healthcare.DicomStoresService().StoreInstances(parent, dicomWebPath, bytes.NewReader(dicomData))
	call.Header().Set("Content-Type", "application/dicom")
	resp, err := call.Do()
	if err != nil {
		return fmt.Errorf("StoreInstances: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("StoreInstances: status %d %s: %s", resp.StatusCode, resp.Status, respBytes)
	}
	fmt.Printf("%s", respBytes)
	return nil
}
