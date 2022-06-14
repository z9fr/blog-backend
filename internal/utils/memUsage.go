package utils

import (
	"runtime"

	"github.com/z9fr/blog-backend/internal/types"
)

func GetMemUsage() types.MemUsage {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats

	var usage types.MemUsage
	usage.Alloc = bToMb(m.Alloc)
	usage.TotalAlloc = bToMb(m.TotalAlloc)
	usage.Sys = bToMb(m.Sys)
	usage.NumGC = m.NumGC

	return usage
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
