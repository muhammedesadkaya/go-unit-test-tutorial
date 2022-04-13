package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_CheckAgeAdult(t *testing.T) {
	//Arrange
	testTime := time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC)

	//Action
	result := CheckAgeAdult(testTime)

	//Assert
	assert.Equal(t, result, 29)

}
