package middleware

import (
	"context"
	"net/http"
	"sslchecker/jwtutil"
	"sslchecker/response"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

const (
	accountIDContextKey = "aid"
)

// Auth middleware checks if requester is authorized to access the given resource
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// for regular secure routes, the "typ" claim can safely be ignored
		sub, _, err := jwtutil.ValidateTokenFromRequest(r)
		if err != nil {
			response.RespondWithError(w, app_error.Unauthorized)
			return
		}

		// the "sub" claim should contain the account ID
		next(w, r.WithContext(context.WithValue(r.Context(), accountIDContextKey, sub)), ps)
	}
}

func AccountIDFromContext(ctx context.Context) (uuid.UUID, error) {
	value, ok := ctx.Value(accountIDContextKey).(string)
	if !ok {
		return uuid.Nil, app_error.Unauthorized
	}

	guid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, app_error.Unauthorized
	}

	return guid, nil
}
