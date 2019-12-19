package styles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/tumasgiu/go-mapbox/lib/base"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Styles api wrapper instance
type Styles struct {
	base *base.Base
}

// NewStyles Create a new Styles API wrapper
func NewStyles(base *base.Base) *Styles {
	return &Styles{base}
}

// ListOpts request options for directions api
type ListOpts struct {
	// The maximum number of styles to return.
	Limit string `url:"limit,omitempty"`
	//The ID of the style after which to start the listing. The style ID is found in the Link header of a response.
	Start string `url:"start,omitempty"`
}

func (s *Styles) List(username string, opts *ListOpts) (styles []Style, err error) {
	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	resp, err := s.base.Request(http.MethodGet, fmt.Sprintf("styles/v1/%s", username), &v, nil)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &styles)
	if err != nil {
		return
	}

	return
}

func (s *Styles) Retrieve(username, styleId string) (style Style, err error) {
	v := url.Values{}
	resp, err := s.base.Request(http.MethodGet, fmt.Sprintf("styles/v1/%s/%s", username, styleId), &v, nil)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &style)
	if err != nil {
		return
	}

	return
}

func (s *Styles) Update(username string, input Style) (style Style, err error) {
	data, err := json.Marshal(input)
	if err != nil {
		return
	}

	buf := bytes.NewReader(data)

	resp, err := s.base.Request(http.MethodPatch, fmt.Sprintf("styles/v1/%s/%s", username, input.Id), nil, buf)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &style)
	if err != nil {
		return
	}

	return
}

func (s *Styles) Create(username string, input Style) (style Style, err error) {
	data, err := json.Marshal(input)
	if err != nil {
		return
	}

	buf := bytes.NewReader(data)

	resp, err := s.base.Request(http.MethodPost, fmt.Sprintf("styles/v1/%s", username), nil, buf)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &style)
	if err != nil {
		return
	}

	return
}