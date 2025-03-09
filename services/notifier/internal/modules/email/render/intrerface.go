package render

import "encoding/json"

type Renderable interface {
	Template() string
	Subject() string
}

func Parse(data []byte, dst interface{}) error {
	if err := json.Unmarshal(data, &dst); err != nil {
		return err
	}
	return nil
}
