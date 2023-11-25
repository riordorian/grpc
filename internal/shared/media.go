package shared

type MediaType int

const (
	IMAGE = iota
	VIDEO
	DOCUMENT
)

type Media struct {
	Name string
	Path string
	Type MediaType
}
