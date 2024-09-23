package basic

type MetaResult[T any] struct {
	Result       T
	hasError     bool
	ErrorMessage string
	ResponseCode int
}
