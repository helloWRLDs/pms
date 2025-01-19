package interfaces

type FieldExtractor interface {
	Columns(excluding ...string) []string
	Values(excluding ...string) []interface{}
}
