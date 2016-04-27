package api

import (
	"goji.io/pat"
	"golang.org/x/net/context"
	"net/http"
)

// Store the connected devices as simple map
var onhubConnections = map[string]bool{}

// How many connected devices are on the network?
func onhubConnectionCount() int {
	count := 0
	for _, v := range onhubConnections {
		if v {
			count = count + 1
		}
	}
	return count
}

func OnhubConnect(c context.Context, w http.ResponseWriter, req *http.Request) {
	device := pat.Param(c, "device")

	// Mark the device as connected
	onhubConnections[device] = true

	renderJSON(w, 200, map[string]interface{}{
		"connected": device,
	})
}

func OnhubDisconnect(c context.Context, w http.ResponseWriter, req *http.Request) {
	device := pat.Param(c, "device")

	// Mark the device as disconnected
	onhubConnections[device] = false

	renderJSON(w, 200, map[string]interface{}{
		"disconnected": device,
	})
}
