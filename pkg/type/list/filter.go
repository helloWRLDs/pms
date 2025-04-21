package list

import (
	"bytes"
	"fmt"
)

type Filters struct {
	Pagination
	Date
	Order
	Fields   map[string]string
	InFields map[string][]string
}

func (f Filters) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "page=%d per_page=%d ", f.Page, f.PerPage)
	fmt.Fprintf(&buf, "created_from=%s created_to=%s ", f.Date.From, f.Date.To)
	fmt.Fprintf(&buf, "order_by=%s ", f.Order.String())
	for k, v := range f.Fields {
		fmt.Fprintf(&buf, "%s=%s ", k, v)
	}

	return buf.String()
}

type Date struct {
	From string `json:"created_from"`
	To   string `json:"created_to"`
}

type Order struct {
	By       string
	Ascended bool
}

func (o Order) String() string {
	if o.Ascended {
		return o.By + " ASC"
	}
	return o.By + " DESC"
}
