package api

import (
	"fmt"

	twilio "github.com/sfreiberg/gotwilio"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

func deliverSMS(ctx context.Context, body string) error {
	client := twilio.NewTwilioClientCustomHTTP(conf.TwilioSID, conf.TwilioToken, urlfetch.Client(ctx))
	_, exc, err := client.SendSMS(conf.TwilioPhonenumber, conf.IFTTTPhonenumber, body, "", "")
	if exc != nil {
		return fmt.Errorf("%s %v", exc.Message, exc.Code)
	}

	return err
}
