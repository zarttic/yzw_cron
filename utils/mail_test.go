package utils

import "testing"

func TestMail(t *testing.T) {
	err := SendEmail("test")
	if err != nil {
		t.Error(err)
	}
}
