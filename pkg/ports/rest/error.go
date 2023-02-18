package rest

import (
	"context"
	"net/http"
)

// ErrBadRequest error http StatusBadRequest.
func ErrBadRequest(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusBadRequest))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	return err
}

// ErrUnauthorized error http StatusUnauthorized.
func ErrUnauthorized(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusUnauthorized))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusUnauthorized)
	return err
}

// ErrPaymentRequired error http StatusPaymentRequired.
func ErrPaymentRequired(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusPaymentRequired))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusPaymentRequired)
	return err
}

// ErrForbidden error http StatusForbidden.
func ErrForbidden(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusForbidden))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusForbidden)
	return err
}

// ErrMethodNotAllowed error http StatusMethodNotAllowed.
func ErrMethodNotAllowed(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusMethodNotAllowed))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusMethodNotAllowed)
	return err
}

// ErrNotAcceptable error http StatusNotAcceptable.
func ErrNotAcceptable(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusNotAcceptable))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotAcceptable)
	return err
}

// ErrProxyAuthRequired error http StatusProxyAuthRequired.
func ErrProxyAuthRequired(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusProxyAuthRequired))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusProxyAuthRequired)
	return err
}

// ErrRequestTimeout error http StatusRequestTimeout.
func ErrRequestTimeout(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusRequestTimeout))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusRequestTimeout)
	return err
}

// ErrUnsupportedMediaType error http StatusUnsupportedMediaType.
func ErrUnsupportedMediaType(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusUnsupportedMediaType))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	return err
}

// ErrUnprocessableEntity error http StatusUnprocessableEntity.
func ErrUnprocessableEntity(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusUnprocessableEntity))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusUnprocessableEntity)
	return err
}

// ErrInternalServerError error http StatusInternalServerError.
func ErrInternalServerError(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusInternalServerError))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusInternalServerError)
	return err
}

// ErrBadGateway error http StatusBadGateway.
func ErrBadGateway(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusBadGateway))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadGateway)
	return err
}

// ErrServiceUnavailable error http StatusServiceUnavailable.
func ErrServiceUnavailable(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusServiceUnavailable))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusServiceUnavailable)
	return err
}

// ErrGatewayTimeout error http StatusGatewayTimeout.
func ErrGatewayTimeout(w http.ResponseWriter, r *http.Request, err error) error {
	*r = *r.WithContext(context.WithValue(r.Context(), CtxStatusCode, http.StatusGatewayTimeout))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusGatewayTimeout)
	return err
}
