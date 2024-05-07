package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
)

type VerifyOTPstruct struct {
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}

func (adpt *Adapter) HandleVerifyOTP(w http.ResponseWriter, r *http.Request) error {
	var body VerifyOTPstruct
	log.Println("handling otp verification")
	if err := json.NewDecoder((r.Body)).Decode(&body); err != nil {
		return err
	}

	result := adpt.app.VerifyOTP(body.Otp, body.PhoneNumber)
	adpt.WriteJSONResponse(w, http.StatusAccepted, map[string]interface{}{"phoneNumber": body.PhoneNumber, "verified": result}, nil)
	return nil
}
