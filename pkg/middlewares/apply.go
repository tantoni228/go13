package middlewares

import "net/http"

// HTTP middleware
type Middleware func(next http.Handler) http.Handler

// Apply applies middlewares to handler.
//
// Usage example:
//
//	handler = middlewares.Apply(
//		handler,
//		middlewares.LoggerProvider(l),
//		middlewares.Logging(),
//	)
func Apply(handler http.Handler, mwares ...Middleware) http.Handler {
	for _, mw := range mwares {
		handler = mw(handler)
	}
	return handler
}
