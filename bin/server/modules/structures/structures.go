package structures

import (
	"encoding/json"
	"net/http"
)

//LoginRequest req
type LoginRequest struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

//LoginResponse res
type LoginResponse struct {
	Message string `json:"message,omitempty"`
}

//RegisterRequest req
type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Class    string `json:"class,omitempty"`
}

//SessionRequest req
type SessionRequest struct {
	SessionID string `json:"sessionid,omitempty"`
}

//DecodeLogin method for decoding structure from request
func DecodeLogin(r *http.Request) (LoginRequest, error) {
	var out LoginRequest
	err := json.NewDecoder(r.Body).Decode(&out)
	return out, err
}

//DecodeSession method for decoding session structure from request
func DecodeSession(r *http.Request) (SessionRequest, error) {
	var out SessionRequest
	err := json.NewDecoder(r.Body).Decode(&out)
	return out, err
}
