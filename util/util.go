package util

import (
	"net/http"
	"strings"

	"github.com/itang/gotang"
	"github.com/nu7hatch/gouuid"
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

func UUID() string {
	u4, err := uuid.NewV4()
	gotang.AssertNoError(err, "")
	return u4.String()
}
