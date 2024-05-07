package main

import (
	httpserver "goOtp/internal/adapters/left/HttpServer"
	MessageQueue "goOtp/internal/adapters/left/messageQueue"
	smsprovider "goOtp/internal/adapters/right/smsProvider"
	"goOtp/internal/application"

	"github.com/gorilla/mux"
)

func main() {

	smsProvider := smsprovider.NewSmsProvider()

	app := application.NewApp(smsProvider)

	messageQ := MessageQueue.NewAdapter("amqps://puxwmcvn:CewRm5Q_7uq_mIGzFnrX6yWpZMEVd1xn@puffin.rmq2.cloudamqp.com/puxwmcvn", app)
	messageQ.MakeConnection()

	httpServerAdpt := httpserver.NewAdapter(mux.NewRouter(), "8001", app)

	httpServerAdpt.Start(messageQ.CloseConnection)
}
