package firebaseauth

import (
	"context"

	"github.com/riposo/firebase-auth/internal"
	"github.com/riposo/riposo/pkg/auth"
	"github.com/riposo/riposo/pkg/plugin"
	"github.com/riposo/riposo/pkg/riposo"
)

func init() {
	auth.Register("firebase", func(_ context.Context, hlp riposo.Helpers) (auth.Method, error) {
		cfg := new(internal.Config)
		if err := hlp.ParseConfig(cfg); err != nil {
			return nil, err
		}
		return internal.New(cfg)
	})

	plugin.Register(plugin.New(
		"firebase-auth",
		map[string]interface{}{
			"description": "Authenticate users via Firebase JWT.",
			"url":         "https://github.com/riposo/firebase-auth",
		},
		nil,
		nil,
	))
}
