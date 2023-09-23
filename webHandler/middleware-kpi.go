package webHandler

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	// map to stock requests (endpoints) & their number of hits
	mapKPIendpoint map[string]int
	//mutex to ensure safe concurrent access to the map : mapEndpoint
	mu sync.Mutex
)

func HitEndPoint(c *gin.Context) {
	// Lock the mutex to ensure safe concurrent access to the map
	mu.Lock()

	// if the first time, we need to initialize the map
	if mapKPIendpoint == nil {
		mapKPIendpoint = make(map[string]int)
	}

	// retrieve path
	path := c.FullPath()

	// increment
	mapKPIendpoint[path]++

	// unlock mutex
	mu.Unlock()

	//continue to handler
	c.Next()
}

func getMostUsedRequest() (string, int) {
	// Lock the mutex to ensure safe concurrent access to the map, and defer the unlock
	mu.Lock()
	defer mu.Unlock()

	mostFrequentRequest := ""
	maxHits := 0

	for path, hits := range mapKPIendpoint {
		if hits > maxHits {
			mostFrequentRequest = path
			maxHits = hits
		}
	}

	return mostFrequentRequest, maxHits
}
