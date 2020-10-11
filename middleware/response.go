package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afief/mockidi/entity"
)

// ResponseMiddleware ...
func ResponseMiddleware(handle entity.Handler) entity.Handler {
	return func(w http.ResponseWriter, r *http.Request) (context.Context, error) {
		resCtx, err := handle(w, r)

		var resp *entity.HTTPResponse

		if err != nil {
			resp = &entity.HTTPResponse{
				Status: 400,
				Body: map[string]string{
					"errorMessage": err.Error(),
				},
			}
		}

		w.Header().Set("Content-Type", "application/json")
		for k, v := range resp.Headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(resp.Status)

		if strBody, ok := resp.Body.(string); ok {
			fmt.Fprint(w, strBody)
			return
		}

		json.NewEncoder(w).Encode(resp.Body)
		return resp, nil
	}
}
