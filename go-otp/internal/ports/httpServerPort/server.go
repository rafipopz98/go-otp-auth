package httpserverport

type HttpServerInterface interface {
	Start(func())
}