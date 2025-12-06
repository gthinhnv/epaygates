package requtil

import (
	"errors"
	"shared/models/staticpagemodel"
	"shared/pkg/utils/stringutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUint64List(c *gin.Context, key string) ([]uint64, error) {
	if arr := c.QueryArray(key); len(arr) > 0 {
		return parseUint64List(arr)
	}

	raw := c.Query(key)
	if raw == "" {
		return nil, nil
	}
	return parseUint64List(strings.Split(raw, ","))
}

func GetInt32List(c *gin.Context, key string) ([]int32, error) {
	if arr := c.QueryArray(key); len(arr) > 0 {
		return parseInt32List(arr)
	}

	raw := c.Query(key)
	if raw == "" {
		return nil, nil
	}
	return parseInt32List(strings.Split(raw, ","))
}

func GetStringList(c *gin.Context, key string) ([]string, error) {
	if arr := c.QueryArray(key); len(arr) > 0 {
		return arr, nil
	}

	raw := c.Query(key)
	if raw == "" {
		return nil, nil
	}
	return strings.Split(raw, ","), nil
}

func GetBool(c *gin.Context, key string) (bool, error) {
	return strconv.ParseBool(c.Query(key))
}

func GetUInt32(c *gin.Context, key string) (uint32, error) {
	parsedValue, err := strconv.ParseUint(c.Query(key), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(parsedValue), nil
}

func GetOrderBy(c *gin.Context) (string, error) {
	raw := c.Query("sort")
	if raw == "" {
		return "", nil // no ordering
	}

	parts := strings.Split(raw, ".")
	if len(parts) != 2 {
		return "", errors.New("invalid orderBy format, expected: field.asc or field.desc")
	}

	field := stringutil.ToSnake(parts[0])
	direction := strings.ToUpper(parts[1])

	// validate direction
	if direction != "ASC" && direction != "DESC" {
		return "", errors.New("invalid order direction, must be asc or desc")
	}

	// validate field to avoid SQL injection
	if !staticpagemodel.AllowedOrderFields[field] {
		return "", errors.New("invalid order field")
	}

	// safe because field + direction are whitelisted
	return field + " " + direction, nil
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

func parseInt32List(values []string) ([]int32, error) {
	out := make([]int32, 0, len(values))
	for _, v := range values {
		if v == "" {
			continue
		}
		parsedValue, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, err
		}
		out = append(out, int32(parsedValue))
	}
	return out, nil
}
