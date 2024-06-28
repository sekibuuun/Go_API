package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/sekibuuun/go_api/apperrors"
	"github.com/sekibuuun/go_api/common"
	"google.golang.org/api/idtoken"
)

const (
	googleClientID = "569348598464-u0tcrp5iembgt9j2ptkkvci89u0931ts.apps.googleusercontent.com"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid req header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		bearer, idToken := authHeaders[0], authHeaders[1]

		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid req header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid req header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
