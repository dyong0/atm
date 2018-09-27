package test

import (
	"bytes"
	"encoding/json"
)

func ParseBody(b *bytes.Buffer) map[string]interface{} {
	var res map[string]interface{}
	json.Unmarshal([]byte(b.String()), &res)

	return res
}
