package smsprovider

type SMSProviderInterface interface {
	InitializeClient()
	SendOtp(to string) error
	CheckOtpIsCorrect(to, code string) bool
}
