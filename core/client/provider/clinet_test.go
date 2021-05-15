package provider

import "testing"

func TestNewProviderClient(t *testing.T) {

	provider, err := NewProviderClient(nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_ = provider
}
