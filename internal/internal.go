package internal

import (
	"net/http"
	"strings"

	"github.com/bsm/firejwt"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/auth"
)

// Config supports custom configuration.
type Config struct {
	Auth struct {
		Firebase struct {
			ProjectID       string   `yaml:"project_id"`
			OnlyDomains     []string `yaml:"only_domains"`
			AllowUnverified bool     `yaml:"allow_unverified"`
		}
	}
}

// New inits a new auth Method.
func New(cfg *Config) (auth.Method, error) {
	dec, err := firejwt.New(cfg.Auth.Firebase.ProjectID)
	if err != nil {
		return nil, err
	}

	return newMethod(dec, cfg), nil
}

func newMethod(dec decoder, cfg *Config) auth.Method {
	w := &wrapper{decoder: dec, cfg: cfg}
	if allow := cfg.Auth.Firebase.OnlyDomains; len(allow) != 0 {
		w.onlyDomains = make(map[string]struct{}, len(allow))
		for _, s := range allow {
			w.onlyDomains[s] = struct{}{}
		}
	}
	return w
}

type decoder interface {
	Stop()
	Decode(string) (*firejwt.Claims, error)
}

type wrapper struct {
	decoder
	onlyDomains map[string]struct{}
	cfg         *Config
}

// Authenticate implements auth.Method interface.
func (w *wrapper) Authenticate(r *http.Request) (*api.User, error) {
	claims, err := w.extract(r)
	if err != nil {
		return nil, err
	}

	// ensure email is verified
	if !claims.EmailVerified && !w.cfg.Auth.Firebase.AllowUnverified {
		return nil, auth.Errorf("email is not verified")
	}

	// check if email (domain) is permitted
	if !w.isPermitted(claims.Email) {
		return nil, auth.Errorf("email not permitted")
	}

	return &api.User{ID: claims.Email}, nil
}

// Close implements io.Closer interface.
func (w *wrapper) Close() error {
	w.Stop()
	return nil
}

func (w *wrapper) isPermitted(email string) bool {
	if len(w.onlyDomains) == 0 {
		return true
	}

	pos := strings.IndexByte(email, '@')
	if pos < 0 {
		return false
	}

	_, ok := w.onlyDomains[email[pos+1:]]
	return ok
}

func (w *wrapper) extract(r *http.Request) (*firejwt.Claims, error) {
	// extract header
	header := r.Header.Get("Authorization")
	if header == "" {
		return nil, auth.Errorf("no authorization header received")
	}

	// validate prefix
	const prefix = "Bearer "
	if len(header) < len(prefix) || !strings.EqualFold(header[:len(prefix)], prefix) {
		return nil, auth.Errorf("not a bearer token")
	}

	// decode claims
	claims, err := w.Decode(header[len(prefix):])
	if err != nil {
		return nil, auth.WrapError(err)
	}
	return claims, nil
}
