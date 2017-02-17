package model

import (
	"testing"
)

func TestCreateApiKey(t *testing.T) {
	apiKey := newApiKey()

	t.Log("%+v", apiKey)
}