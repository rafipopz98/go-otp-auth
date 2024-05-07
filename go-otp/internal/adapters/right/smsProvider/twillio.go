package smsprovider

import (
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type smsproviderStruct struct {
	accountSid       string
	authToken        string
	verifyServiceSid string
	client           *twilio.RestClient
}

func NewSmsProvider() *smsproviderStruct {
	sms := &smsproviderStruct{}
	sms.InitializeClient()
	return sms
}

func (S *smsproviderStruct) InitializeClient() {
	S.accountSid = "AC98646446de23bf93a7123c23dad0306a"
	S.authToken = "7adc7df80153f9f2ecf4c26130e49b9d"
	S.verifyServiceSid = "VA144555c36dba186c7ebd33a793e1a63b"

	S.client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: S.accountSid,
		Password: S.authToken,
	})
}

func (S *smsproviderStruct) SendOtp(to string) error {
	params := &openapi.CreateVerificationParams{}
	log.Println("generating OTPPPPPPPPPPPPPPPPPP")
	to = "+91 " + to
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := S.client.VerifyV2.CreateVerification(S.verifyServiceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
		return nil
	}

}

func (S *smsproviderStruct) CheckOtpIsCorrect(to, code string) bool {
	params := &openapi.CreateVerificationCheckParams{}
	log.Println("to:", to, " code:", code)
	to = "+91 " + to
	params.SetTo(to)
	params.SetCode(code)
	resp, err := S.client.VerifyV2.CreateVerificationCheck(S.verifyServiceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Println("Incorrect!")
		return false
	}
}
