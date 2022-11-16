package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) StoreInstances(ctx echo.Context, organizationID string) error {
	fmt.Printf("start StoreInstances")
	f, _ := io.ReadAll(ctx.Request().Body)
	// fmt.Printf("%s", f)

	// return ctx.JSON(http.StatusOK, organizationID)

	return s.dicomwebUsecase.StoreInstance(f, "mmorito", "mmorito", "stidies")
}

func (s *server) RetrieveStudy(ctx echo.Context, organizationID string, studyInstanceUID string) error {
	return ctx.JSON(http.StatusOK, studyInstanceUID)
}
