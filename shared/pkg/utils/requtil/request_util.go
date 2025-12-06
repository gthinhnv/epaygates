package requtil

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetIDs(c *gin.Context) ([]uint64, error) {
	if arr := c.QueryArray("ids"); len(arr) > 0 {
		return parseUint64List(arr)
	}

	raw := c.Query("ids")
	if raw == "" {
		return nil, nil
	}
	return parseUint64List(strings.Split(raw, ","))
}

func parseUint64List(values []string) ([]uint64, error) {
	out := make([]uint64, 0, len(values))
	for _, v := range values {
		if v == "" {
			continue
		}
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, nil
}
