package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mopeneko/h2o-dstat/backend/model"
	"golang.org/x/xerrors"
)

var (
	connections = 0
)

func init() {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				getter()
			}
		}
	}()
}

func getter() {
	statusResp, err := http.Get("http://localhost/server-status/json")
	if err != nil {
		err = xerrors.Errorf("failed to get: %w", err)
		log.Printf("%+v", err)
		return
	}

	defer statusResp.Body.Close()

	data, err := io.ReadAll(statusResp.Body)
	if err != nil {
		err = xerrors.Errorf("failed to read response body: %w", err)
		log.Printf("%+v", err)
		return
	}

	status := new(model.ServerStatus)
	err = json.Unmarshal(data, status)
	if err != nil {
		err = xerrors.Errorf("failed to unmarshal json: %w", err)
		log.Printf("%+v", err)
		return
	}

	connections = status.Connections
}

func GetStatus(c echo.Context) error {
	res := &model.GetStatusResponse{
		Connections: connections,
	}
	return c.JSON(http.StatusOK, res)
}
