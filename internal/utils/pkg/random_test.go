package pkg_test

import (
	"reflect"
	"testing"

	"github.com/nurcahyaari/kite/internal/utils/pkg"
)

func TestRandomString(t *testing.T) {
	returnString := pkg.RandomString(32)

	if reflect.TypeOf(returnString).String() != "string" {
		t.Errorf("The return is not a string")
	}
}
