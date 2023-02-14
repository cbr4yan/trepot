package company

import (
	"github.com/cbr4yan/trepot/core"
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group, service core.CompanyService) {
	handler := &Handler{companyService: service}

	g.POST("/companies", handler.Create)
	g.DELETE("/companies/:id", handler.Delete)
}

type Handler struct {
	companyService core.CompanyService
}

func (h *Handler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	in := &core.CompanyRequest{}

	if err := c.Bind(in); err != nil {
		return err
	}

	out, err := h.companyService.Create(ctx, in)
	if err != nil {
		return err
	}

	return c.JSON(200, out)
}

func (h *Handler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	if err := h.companyService.Delete(ctx, id); err != nil {
		return err
	}

	return c.NoContent(200)
}
