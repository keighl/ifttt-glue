package api

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func NestMotionDetected(c context.Context, w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	soundAlarm := (onhubConnectionCount() == 0)

	// Sound the alarm if nobody's home!
	if soundAlarm {
		deliverSMS(ctx, "#alarm")
	}

	renderJSON(w, 200, map[string]interface{}{
		"soundAlarm": soundAlarm,
	})
}
