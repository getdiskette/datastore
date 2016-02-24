package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
)

func TestPing(t *testing.T) {
	start()
	// time.Sleep(1 * time.Millisecond)

	errs := frisby.Create("Test ping").
		Get("http://localhost:5025/ping").
		Send().
		ExpectStatus(http.StatusOK).
		ExpectContent("pong").
		AfterText(
		func(F *frisby.Frisby, text string, err error) {
			assert.Nil(t, err)
			assert.Equal(t, "pong", text)
		}).
		Errors()

	for _, err := range errs {
		assert.Nil(t, err)
	}
}
