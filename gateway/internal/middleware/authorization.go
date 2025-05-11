package middleware

import (
	"net/http"
	"strings"

	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.gatewaySettings.NeedAuth(r) {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Vary", "Authorization")

		authorozationHeader := r.Header.Get("Authorization")

		if authorozationHeader == "" {
			util.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		headerParts := strings.Split(authorozationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			util.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]
		payload, err := m.tokenVerifier.VerifyToken(token)
		if err != nil {
			util.InvalidAuthenticationTokenResponse(w, r)
			return
		}

		if !m.gatewaySettings.HasPermission(payload.Role) {
			util.UnauthorizedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
