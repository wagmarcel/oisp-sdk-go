package oispapi

import (
	"net/http"
	"regexp"
)

// Oispapi is managing an API session
type Oispapi struct {
	token        string
	url          string
	accounts     []Account
	userID       string
	activAccount int
}

// Account contains the detail of one account
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Device contains the device details
type Device struct {
	DeviceID    string            `json:"deviceId"`
	Name        string            `json:"name"`
	GatewayID   string            `json:"gatewayId"`
	DomainID    string            `json:"domainId,omitempty"`
	Status      string            `json:"status,omitempty"`
	Created     int64             `json:"created,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Components  []Component       `json:"components,omitempty"`
	Contact     string            `json:"contact,omitempty"`
	Loc         []float64         `json:"loc,omitempty"`
	Description string            `json:"desription,omitempty"`
}

// Component reprecents a single device component
type Component struct {
	CID             string        `json:"cid"`
	Name            string        `json:"name"`
	ComponentTypeID string        `json:"componentTypeId"`
	Type            string        `json:"type"`
	ComponentType   ComponentType `json:"componentType"`
}

// ComponentType describe the details of a component type
type ComponentType struct {
	CTID        string `json:"_id"`
	ID          string `json:"id"`
	DomainID    string `json:"domainID"`
	Dimension   string `json:"dimension"`
	Default     string `json:"default"`
	Display     string `json:"display"`
	Format      string `json:"format"`
	Measureunit string `json:"measureunit"`
	Version     string `json:"version"`
	Type        string `json:"type"`
	DataType    string `json:"dataType"`
	Min         string `json:"min"`
	Max         string `json:"max"`
}

func setHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
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

func makeURL(url string, path string, replacements map[string]string) string {
	replaced := path
	for k, v := range replacements {
		re := regexp.MustCompile(`(\{` + k + `\})`)
		replaced = re.ReplaceAllString(replaced, v)
	}
	return url + replaced
}

// GetToken returns the access token value
func (o *Oispapi) GetToken() string {
	return o.token
}
