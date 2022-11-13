package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mopeneko/h2o-dstat/backend/model"
	"github.com/mopeneko/h2o-dstat/backend/pkg/echoutils"
	"golang.org/x/xerrors"
)

func GetStatus(c echo.Context) error {
	statusResp, err := http.Get("http://localhost/server-status/json")
	if err != nil {
		err = xerrors.Errorf("failed to get: %w", err)
		log.Printf("%+v", err)
		return echoutils.ReturnInternalServerError(c)
	}

	defer statusResp.Body.Close()

	data, err := io.ReadAll(statusResp.Body)
	if err != nil {
		err = xerrors.Errorf("failed to read response body: %w", err)
		log.Printf("%+v", err)
		return echoutils.ReturnInternalServerError(c)
	}

	status := new(model.ServerStatus)
	err = json.Unmarshal(data, status)
	if err != nil {
		err = xerrors.Errorf("failed to unmarshal json: %w", err)
		log.Printf("%+v", err)
		return echoutils.ReturnInternalServerError(c)
	}

	res := &model.GetStatusResponse{
		Connections: status.Connections,
	}
	return c.JSON(http.StatusOK, res)
}
