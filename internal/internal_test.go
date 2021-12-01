package internal_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bsm/firejwt"
	. "github.com/bsm/ginkgo"
	. "github.com/bsm/gomega"
	"github.com/riposo/firebase-auth/internal"
	"github.com/riposo/riposo/pkg/api"
	"github.com/riposo/riposo/pkg/auth"
)

var _ = Describe("Firebase Auth Method", func() {
	var subject auth.Method
	var req *http.Request
	var claims *firejwt.Claims

	BeforeEach(func() {
		req = httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer VALID.TOKEN")

		claims = &firejwt.Claims{
			Email:         "alice@test.host",
			EmailVerified: true,
		}
		subject = internal.Mock(claims, nil)
	})

	AfterEach(func() {
		Expect(subject.(io.Closer).Close()).To(Succeed())
	})

	It("should authenticate", func() {
		Expect(subject.Authenticate(req)).To(Equal(&api.User{ID: "alice@test.host"}))
	})

	It("should fail gracefully", func() {
		req.Header.Del("Authorization")
		_, err := subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))
		Expect(err).To(MatchError(`no authorization header received`))

		req.Header.Set("Authorization", "Other VALID.TOKEN")
		_, err = subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))
		Expect(err).To(MatchError(`not a bearer token`))

		req.Header.Set("Authorization", "Bearer BAD.TOKEN")
		_, err = subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))

		subject = internal.Mock(&firejwt.Claims{}, nil)
		_, err = subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))
		Expect(err).To(MatchError(`some error`))
	})

	It("should reject unverified", func() {
		claims.EmailVerified = false
		subject = internal.Mock(claims, nil)
		_, err := subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))
		Expect(err).To(MatchError(`email is not verified`))

		config := new(internal.Config)
		config.Auth.Firebase.AllowUnverified = true
		subject = internal.Mock(claims, config)
		Expect(subject.Authenticate(req)).To(Equal(&api.User{ID: "alice@test.host"}))
	})

	It("may limit domains", func() {
		config := new(internal.Config)
		config.Auth.Firebase.OnlyDomains = []string{"other.host", "test.host"}
		subject = internal.Mock(claims, config)
		Expect(subject.Authenticate(req)).To(Equal(&api.User{ID: "alice@test.host"}))

		claims.Email = "alice@external.host"
		subject = internal.Mock(claims, config)
		_, err := subject.Authenticate(req)
		Expect(err).To(MatchError(auth.ErrUnauthenticated))
		Expect(err).To(MatchError(`email not permitted`))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal")
}
