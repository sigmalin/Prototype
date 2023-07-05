package response

type Response interface {
	Add(string, interface{})
	Error(int, string) error
	Message() error
}
