package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (*server) StoreInstances(ctx echo.Context, organizationID string) error {
	return ctx.JSON(http.StatusOK, organizationID)
}

func (*server) RetrieveStudy(ctx echo.Context, organizationID string, studyInstanceUID string) error {
	return ctx.JSON(http.StatusOK, studyInstanceUID)
}
