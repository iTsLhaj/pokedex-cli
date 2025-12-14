package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestExit(t *testing.T) {
	/*
		there's nothing to test in exit
		it should be tested manually
	*/
	t.Logf("meow!")
}

func TestHelp(t *testing.T) {
	t.Run("is it printing the right help message", func(t *testing.T) {
		buf := new(bytes.Buffer)

		err := commandHelpMock(buf)
		if err != nil {
			t.Fatal(err)
		}
		if buf.String() != helpMessage {
			t.Errorf("help message does not match the expected output")
		}
	})
}

func TestMap(t *testing.T) {
	client := http.Client{}

	t.Run("test API status", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, baseUrl, nil)

		res, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("status code does not match the expected")
		}
	})

}
