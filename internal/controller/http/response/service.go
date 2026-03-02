package response

type ServiceResponse[T any] struct {
	data             T
	errorCode        string
	errorDescription string
	errorURI         string
}

func NewServiceResponse[T any](data T, code, description, uri string) ServiceResponse[T] {
	return ServiceResponse[T]{
		data:             data,
		errorCode:        code,
		errorDescription: description,
		errorURI:         uri,
	}
}

func (r *ServiceResponse[T]) Data() T {
	return r.data
}

func (r *ServiceResponse[T]) ErrorCode() string {
	return r.errorCode
}

func (r *ServiceResponse[T]) ErrorDescription() string {
	return r.errorDescription
}

func (r *ServiceResponse[T]) ErrorURI() string {
	return r.errorURI
}

func (r *ServiceResponse[T]) IsSuccess() bool {
	return r.errorCode == "" && r.errorDescription == "" && r.errorURI == ""
}
