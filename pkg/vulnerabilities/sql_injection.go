package vulnerabilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dtr0x80/jowfuzz/pkg/requests"
)

func CheckSqlInjection(host string) bool {
	payload := ""
	vullString := "mysql_fetch_array"

	hostWithPayload := host + payload

	req, err := requests.Get(hostWithPayload)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	if strings.Contains(string(body), vullString) {
		return true
	}

	return false
}
