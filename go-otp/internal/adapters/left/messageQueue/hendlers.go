package MessageQueue

import (
	"encoding/json"
	application "goOtp/internal/ports/Application"
	"log"
)

type generateOTPStruct struct {
	PhoneNumber string `json:"phoneNumber"`
}

type VerifyOTPStruct struct {
	PhoneNumber string `json:"phoneNumber"`
	OTP         string `json:"otp"`
}

func trialHandler(body message, app application.ApplicationInterface) bool {
	log.Println(body)
	return true
}

func GenerateOTPHandler(body message, app application.ApplicationInterface) bool {
	var data generateOTPStruct
	err := json.Unmarshal(body.Data, &data)
	if err != nil {
		log.Println("error while unmarshalling in GenerateOTPHandler", err)
		return false
	}
	log.Println("generating otp at mq handler")
	if err := app.GenerateOTPForPhoneNumber(data.PhoneNumber); err != nil {
		return false
	}

	return true
}

func VerifyOTPHandler(body message, app application.ApplicationInterface) bool {
	var data VerifyOTPStruct
	err := json.Unmarshal(body.Data, &data)
	if err != nil {
		log.Println("error while unmarshalling in GenerateOTPHandler", err)
		return false
	}

	return app.VerifyOTP(data.OTP, data.PhoneNumber)
}
