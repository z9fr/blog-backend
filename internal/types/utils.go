package types

type MemUsage struct {
	Alloc      uint64 `json:"alloc"`
	TotalAlloc uint64 `json:"total_alloc"`
	Sys        uint64 `json:"sys"`
	// NumGC is the number of completed GC cycles.
	NumGC uint32 `json:"numGC"`
}
