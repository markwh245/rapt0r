package vulnerabilities

import (
	"testing"
)

func TestSqlInjection(t *testing.T) {
	host := "http://testphp.vulnweb.com/listproducts.php?cat=1'"

	if !CheckSqlInjection(host) {
		t.Errorf("Check SQL Injection failed")
	}
}
