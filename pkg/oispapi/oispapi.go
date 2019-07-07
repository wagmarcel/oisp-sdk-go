package oispapi

// Oispapi is managing an API session
type Oispapi struct {
	token string
	url   string
}

// NewOispAPIFromToken is initiating the Oispapi struct from a user token
func NewOispAPIFromToken(token string, url string) *Oispapi {

	return &Oispapi{
		token: token,
		url:   url,
	}
}

// NewOispAPIFromUser is initiating the Oispapi struct from username/password
func NewOispAPIFromUser(username string, password string, url string) (*Oispapi, error) {
	o := &Oispapi{
		url: url,
	}
	err := o.getUserToken(username, password)
	return o, err
}
