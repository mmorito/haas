package server

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/mnes/haas/src/services/imaging-study/openapi"
)

func (s *server) CreateDicomStore(ctx echo.Context, organizationID string) error {
	body := new(openapi.CreateDicomStoreJSONRequestBody)
	if err := ctx.Bind(body); err != nil {
		return ctx.JSON(http.StatusBadRequest, "NG")
	}
	return ctx.JSON(http.StatusOK, organizationID)
}

func (s *server) StoreInstances(ctx echo.Context, organizationID string) error {
	zap.L().Info("start StoreInstances", zap.String("organizationID", organizationID))
	f, _ := io.ReadAll(ctx.Request().Body)
	// fmt.Printf("%s", f)

	// return ctx.JSON(http.StatusOK, organizationID)

	return s.dicomwebUsecase.StoreInstance(f, "mmorito", "mmorito", "stidies")
}

func (s *server) RetrieveStudy(ctx echo.Context, organizationID string, studyInstanceUID string) error {
	return ctx.JSON(http.StatusOK, studyInstanceUID)
}
