package http

import (
	"net/http"

	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/utils"
)

func (h *Handler) GetApplicationHealth(w http.ResponseWriter, r *http.Request) {
	usage := utils.GetMemUsage()

	h.sendOkResponse(w, struct {
		Status string         `json:"status"`
		Usage  types.MemUsage `json:"usage"`
	}{
		Status: "running",
		Usage:  usage,
	})
}
