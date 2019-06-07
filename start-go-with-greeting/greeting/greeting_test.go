package greeting_test

import (
	"karnwad-x-internship/greeting"
	"testing"
)

func TestGreetingShuoldReturnHelloWorld(t *testing.T) {
	//arrange
	expectedResult := "Good morning , karnawat"
	//action
	actualResult := greeting.BuildGreetingMessage("karnawat")
	//assert
	if actualResult != expectedResult {
		t.Errorf(" %v but got %v ", expectedResult, actualResult)
	}
}
