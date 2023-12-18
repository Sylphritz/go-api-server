package util

import (
	"path"
)

func GetApiPath(paths ...string) string {
	apiPath := append([]string{"/api/v1"}, paths...)
	return path.Join(apiPath...)
}
