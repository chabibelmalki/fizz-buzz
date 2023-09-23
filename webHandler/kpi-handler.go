package webHandler

import (
	"github.com/gin-gonic/gin"
)

type stKPI_MostUsedRequest struct {
	Request string `json:"request"`
	Hits    int    `json:"hit_count"`
}

func Handler_KPI_MostUsedRequest(c *gin.Context) {
	requestPath, nbHits := getMostUsedRequest()

	kpi := stKPI_MostUsedRequest{
		Request: requestPath,
		Hits:    nbHits,
	}
	HTTPDeclare_Success(c, kpi)
}
