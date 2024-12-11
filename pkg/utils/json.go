package utils

import (
	"encoding/json"
)

func JSON(t any) string {
	j, _ := json.MarshalIndent(t, "", "    ")
	return string(j)
}
