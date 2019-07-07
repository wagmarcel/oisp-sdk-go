package oispapi

import "net/http"

// Account contains the detail of one account
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Oispapi is managing an API session
type Oispapi struct {
	token    string
	url      string
	accounts []Account
	userID   string
}

func setAuthHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

// NewOispAPIFromToken is initiating the Oispapi struct from a user token
func NewOispAPIFromToken(token string, url string) (*Oispapi, error) {

	o := &Oispapi{
		token: token,
		url:   url,
	}

	err := o.getAuthTokenInfo()
	if err != nil {
		return nil, err
	}
	return o, nil
}

// NewOispAPIFromUser is initiating the Oispapi struct from username/password
func NewOispAPIFromUser(username string, password string, url string) (*Oispapi, error) {
	o := &Oispapi{
		url: url,
	}
	err := o.getUserToken(username, password)
	if err != nil {
		return nil, err
	}
	err = o.getAuthTokenInfo()
	if err != nil {
		return nil, err
	}
	return o, nil
}

// GetToken returns the access token value
func (o *Oispapi) GetToken() string {
	return o.token
}
