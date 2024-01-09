package validate

import (
	"testing"
)

func TestCheckValidateCode(t *testing.T) {
	testval := checkValidateCode(1)
	if testval != "test" {
		t.Errorf("not valude = test")
	}
}
