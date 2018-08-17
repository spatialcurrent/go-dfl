package dfl

type Reader interface {
	ReadAll() ([]byte, error)
	ReadRange(start int, end int) ([]byte, error)
}
