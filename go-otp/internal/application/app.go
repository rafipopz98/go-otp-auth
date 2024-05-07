package application

import (
	smsprovider "goOtp/internal/ports/SmsProvider"
)

type AppStruct struct {
	smsProvider      smsprovider.SMSProviderInterface
	MessageQueueChan chan []byte
}

func NewApp(smsProvider smsprovider.SMSProviderInterface) *AppStruct {
	return &AppStruct{
		smsProvider:      smsProvider,
		MessageQueueChan: make(chan []byte),
	}
}

// Only 1 feature: generate OTP for the Phone and send it to the user through Twillio

func (a *AppStruct) GenerateOTPForPhoneNumber(phoneNumber string) error {
	return a.smsProvider.SendOtp(phoneNumber)
}

func (a *AppStruct) VerifyOTP(code, PhoneNumber string) bool {
	result := a.smsProvider.CheckOtpIsCorrect(PhoneNumber, code)

	// msg := make(map[string]interface{})
	// msg["Operation"] = "OTP_VERIFICATION_RESULT"
	// msg["FromService"] = "OTP_SERVICE"
	// msg["Data"] = map[string]interface{}{
	// 	"phoneNumber": PhoneNumber,
	// 	"Verifed":     result,
	// }
	// messageInBytes, err := json.Marshal(msg)
	// if err != nil {
	// 	log.Panic("error while converting to json")
	// }
	// a.sendMessageToMQ(messageInBytes)

	return result
}

func (app *AppStruct) ReturnMessageQueueChan() *chan []byte {
	return &app.MessageQueueChan
}

func (app *AppStruct) sendMessageToMQ(msg []byte) {
	app.MessageQueueChan <- msg

}

// func (A *AppStruct) generateOTP() string {
// 	// Seed the random number generator with the current timestamp
// 	rand.Seed(time.Now().UnixNano())

// 	// Generate a random 6-digit number
// 	otp := rand.Intn(900000) + 100000

// 	// Convert the number to a string and return it
// 	return fmt.Sprintf("%06d", otp)
// }
