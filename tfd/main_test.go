package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	InputData int
	Answer    int
	Expected  int
}

var Cases []TestCase = []TestCase{
	{
		InputData: 0,
		Expected:  1,
	},
	{
		InputData: 1,
		Expected:  1,
	},
	{
		InputData: 3,
		Expected:  6,
	},
	{
		InputData: 5,
		Expected:  120,
	},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if test.Answer = factorial(test.InputData); test.Answer != test.Expected {
			t.Errorf("Test case %d failed: input data %v, result %v expected %v", id, test.InputData, test.Answer, test.Expected)
		}
	}
}

type HttpTestCase struct {
	Name     string
	Numeric  int
	Expected []byte
}

var HttpCases = []HttpTestCase{
	{
		Name:     "first case",
		Numeric:  1,
		Expected: []byte("1"),
	},
	{
		Name:     "second test",
		Numeric:  3,
		Expected: []byte("6"),
	},
	{
		Name:     "third test",
		Numeric:  5,
		Expected: []byte("120"),
	},
}

func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandlerFactorial)
	for _, test := range HttpCases {
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil)
			if err != nil {
				t.Error(err)
			}
			handler.ServeHTTP(recorder, request)
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input %v, result %v, expected %v",
					test.Name,
					test.Numeric,
					string(recorder.Body.Bytes()),
					string(test.Expected))
			}
		})
	}
}
