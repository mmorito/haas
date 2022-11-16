package healthcare

import (
	"bytes"
	"fmt"
	"io/ioutil"

	healthcare "google.golang.org/api/healthcare/v1beta1"

	"github.com/mnes/haas/services/imaging-study/src/config"
	h "github.com/mnes/haas/services/imaging-study/src/driver/healthcare"
)

type DicomStoreRepo interface {
	StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error
}

type dicomStoreRepo struct {
	dicomStore *healthcare.ProjectsLocationsDatasetsDicomStoresService
}

func NewDicomStoreRepo(hd *h.Healthcare) DicomStoreRepo {
	return &dicomStoreRepo{
		dicomStore: hd.DicomStoresService(),
	}
}

// StoreInstance stores DICOM inctance.
func (r *dicomStoreRepo) StoreInstance(dicomData []byte, datasetID, dicomStoreID, dicomWebPath string) error {
	env := config.Env()

	parent := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/dicomStores/%s", env.ProjectID, env.Location, datasetID, dicomStoreID)

	call := r.dicomStore.StoreInstances(parent, dicomWebPath, bytes.NewReader(dicomData))
	call.Header().Set("Content-Type", "application/dicom")
	resp, err := call.Do()
	if err != nil {
		return fmt.Errorf("StoreInstances: %v", err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response: %v", err)
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("StoreInstances: status %d %s: %s", resp.StatusCode, resp.Status, respBytes)
	}
	fmt.Printf("%s", respBytes)
	return nil
}
