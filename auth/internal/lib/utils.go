package lib

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gookit/slog"
)

func InterfaceToStringMap(input map[string]interface{}) (map[string]string, error) {
	result := make(map[string]string)

	for key, value := range input {
		switch v := value.(type) {
		case string:
			result[key] = v
		case int, int32, int64, float32, float64, bool:
			result[key] = fmt.Sprintf("%v", v) // Convert numbers and booleans to string
		default:
			return nil, errors.New(fmt.Sprintf("unsupported type for key %s: %T", key, value))
		}
	}

	return result, nil
}

func intFromEnv(key string) (int, error) {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		slog.Errorf("error parsing %s: %v", key, err)
		return 60, err
	}
	return val, err
}
