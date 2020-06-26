package handler

import (
	"github.com/Calenaur/pandemic/handler/response"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (h *Handler) getDiseasesHandler(c echo.Context) error {
	events, err := h.ds.GetDiseases()
	if err != nil {
		return response.MessageHandler(err, "This isn't working", c)
	}

	return c.JSON(http.StatusOK, events)

}

func (h *Handler) getDiseaseByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	event, err := h.ds.GetByID(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) getDiseasesForUserHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	diseases, err := h.ds.GetDiseasesForUser(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, diseases)

}

func (h *Handler) getAvailableDiseasesHandler(c echo.Context) error {
	id, _, _ := getUserFromToken(c)

	diseases, err := h.ds.GetDiseasesList(id)
	if err != nil {
		return response.MessageHandler(err, "", c)
	}

	return c.JSON(http.StatusOK, diseases)

}
