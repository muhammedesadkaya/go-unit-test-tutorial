package helper

import (
	"testing"
	"time"
)

func Test_CheckAgeAdult(t *testing.T) {

	testTime := time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC)

	result := CheckAgeAdult(testTime)

	if result != 29 {
		t.Errorf("CheckAgeAdult(1993-08-26) FAILED. Excepted %d, got %d", 29, result)
	} else {
		t.Logf("CheckAgeAdult(1993-08-26) PASSED. Excepted %d, got %d", 29, result)
	}

}
