package main

import (
	"github.com/riposo/firebase-auth/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/auth"
	"github.com/riposo/riposo/pkg/plugin"
	"github.com/riposo/riposo/pkg/riposo"
)

func init() {
	auth.Register("firebase", func(_ *riposo.Helpers) (auth.Method, error) {
		cfg := new(internal.Config)
		if err := riposo.ParseEnv(cfg); err != nil {
			return nil, err
		}
		return internal.New(cfg)
	})
}

// Plugin export definition.
func Plugin(_ *api.Routes) (plugin.Plugin, error) {
	return plugin.New(
		"firebase_auth",
		map[string]interface{}{
			"description": "Authenticate users via Firebase JWT.",
			"url":         "https://github.com/riposo/firebase-auth",
		},
		nil,
	), nil
}

func main() {}
