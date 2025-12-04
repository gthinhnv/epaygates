package commonhandler

import (
	"apigateway/pkg/utils/modelutil"
	"net/http"
	"shared/models/commonmodel"
	"shared/pkg/utils/apiutil"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	statusesOnce sync.Once
	statusesList []*commonmodel.StatusItem
)

type CommonHandler struct{}

// NewCommonHandler creates handler but avoids rebuilding statuses
func NewCommonHandler() *CommonHandler {
	statusesOnce.Do(func() {
		statusesList = modelutil.BuildStatuses()
	})
	return &CommonHandler{}
}

func (h *CommonHandler) ListStatuses(c *gin.Context) {
	c.JSON(http.StatusOK, apiutil.Response{
		Code:    apiutil.CODE_SUCCESS,
		Message: "success",
		Data:    statusesList, // zero copy
	})
}
