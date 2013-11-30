package api

import (
	"net/http"
)

func CheckHostAlive(hostaddr string) bool {
	_, err := http.Get("http://" + hostaddr)
	if err != nil {
		return false
	}
	return true
}
