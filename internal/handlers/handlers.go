package handlers

import (
	"net/http"

	db "github.com/backstagefood/backstagefood/internal/repositories"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echoEngine *echo.Echo
	database   *db.ApplicationDatabase
}

func New(echoEngine *echo.Echo, databaseConnection *db.ApplicationDatabase) *Handler {
	return &Handler{
		echoEngine: echoEngine,
		database:   databaseConnection,
	}
}

// Health godoc
// @Summary Health check
// @Description Check if the server and the database are up and running.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) Health(c echo.Context) error {
	databaseStatus := "UP"
	if err := h.database.DataBaseHeatlh(); err != nil {
		databaseStatus = "DOWN"
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "UP", "database": databaseStatus})

}
