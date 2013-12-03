package api

import (
	"net/http"
	"strings"
)

const httpPrefix = "http://"

func CheckHostAlive(hostaddr string) bool {
	hostHttpURL := hostaddr
	if !strings.HasPrefix(hostHttpURL, httpPrefix) {
		hostHttpURL = httpPrefix + hostHttpURL
	}
	_, err := http.Get(hostHttpURL)
	if err != nil {
		return false
	}
	return true
}
