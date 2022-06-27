package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/z9fr/blog-backend/internal/types"
	"github.com/z9fr/blog-backend/internal/utils"
)

func (h *Handler) GetApplicationHealth(w http.ResponseWriter, r *http.Request) {
	usage := utils.GetMemUsage()
	dbstatus := h.PerformanceDatabase.Stats()
	totalPostCount := h.PostService.TotalPostCount()
	publishedPostCount := h.PostService.TotalPublishedPostCount()

	h.sendOkResponse(w, struct {
		Status                   string           `json:"status"`
		Usage                    types.MemUsage   `json:"usage"`
		DBstatus                 [5]time.Duration `json:"dbStatus"`
		Uptime                   string           `json:"uptime"`
		TotalPostsCount          int64            `json:"total_posts"`
		TotalPublishedPostsCount int64            `json:"published_posts"`
	}{
		Status:                   "running",
		Usage:                    usage,
		DBstatus:                 dbstatus,
		Uptime:                   fmt.Sprintf("%s", utils.Uptime(h.StartTime)),
		TotalPostsCount:          totalPostCount,
		TotalPublishedPostsCount: publishedPostCount,
	})
}
