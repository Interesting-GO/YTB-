package test

import (
	"github.com/dollarkillerx/easyutils"
	"testing"
)

func TestTime(t *testing.T) {
	t.Log(easyutils.TimeGetTimeToString(easyutils.TimeGetNowTimeStr()))
}
