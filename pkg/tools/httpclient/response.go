package httpclient

import "encoding/json"

type Response struct {
	Data   []byte
	Status int
}

func (r *Response) ScanJSON(dst any) error {
	return json.Unmarshal(r.Data, dst)
}
