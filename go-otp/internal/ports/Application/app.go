package application

type ApplicationInterface interface {
	GenerateOTPForPhoneNumber(phoneNumber string) error
	VerifyOTP(code, PhoneNumber string) bool
	// sendMessageToMQ(msg []byte)
	ReturnMessageQueueChan() *chan []byte
}
