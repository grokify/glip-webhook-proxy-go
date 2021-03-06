package config

import (
	"testing"
)

var ConfigurationTests = []struct {
	v    int
	want string
}{
	{8080, ":8080"}}

func TestConfigurationAddress(t *testing.T) {
	for _, tt := range ConfigurationTests {
		cfg := Configuration{
			Port: tt.v}

		addr := cfg.Address()
		if tt.want != addr {
			t.Errorf("Configuration.Address(%v): want %v, got %v", tt.v, tt.want, addr)
		}
	}
}
