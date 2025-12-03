package commonhandler

import (
	"apigateway/gen/go/commonpb"
	"net/http"
	"shared/pkg/utils/apiutil"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
)

type commonItem struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

var (
	statusesOnce sync.Once
	statusesList []commonItem
)

func buildStatuses() {
	// Allocate exact size
	statusesList = make([]commonItem, 0, len(commonpb.Status_value))

	for name, val := range commonpb.Status_value {
		statusesList = append(statusesList, commonItem{
			Name:  name,
			Value: val,
		})
	}

	// optional but recommended: stable order
	sort.Slice(statusesList, func(i, j int) bool {
		return statusesList[i].Value < statusesList[j].Value
	})
}

type CommonHandler struct{}

// NewCommonHandler creates handler but avoids rebuilding statuses
func NewCommonHandler() *CommonHandler {
	statusesOnce.Do(buildStatuses)
	return &CommonHandler{}
}

func (h *CommonHandler) ListStatuses(c *gin.Context) {
	c.JSON(http.StatusOK, apiutil.Response{
		Code:    apiutil.CODE_SUCCESS,
		Message: "success",
		Data:    statusesList, // zero copy
	})
}
