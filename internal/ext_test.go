package internal

import (
	"fmt"

	"github.com/bsm/firejwt"
	"github.com/riposo/riposo/pkg/auth"
)

type mockClaims firejwt.Claims

func (*mockClaims) Stop() {}
func (m *mockClaims) Decode(s string) (*firejwt.Claims, error) {
	if s == "VALID.TOKEN" {
		return (*firejwt.Claims)(m), nil
	}
	return nil, fmt.Errorf("some error")
}

// Mock returns an auth method with a mock JWT validator.
func Mock(claims *firejwt.Claims, cfg *Config) auth.Method {
	if cfg == nil {
		cfg = new(Config)
	}
	return newMethod((*mockClaims)(claims), cfg)
}
