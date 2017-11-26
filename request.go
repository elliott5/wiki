package wiki

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Request sets up the request against the api with the correct parameters
// and has functionality to fetch the data and convert it to a response.
type Request struct {
	*url.URL
}

// ThumbSize is the maximum size of the returned image, as a string in px.
var ThumbSize = "100"

// NewRequest creates a new request against baseURL for language.
// Language is interpolated in the baseURL if asked, if not it is ignored.
// Query is the title of the page to fetch.
// Returns an error if the URL can not be parsed.
func NewRequest(baseURL, query, language string) (*Request, error) {
	if strings.Contains(baseURL, "%s") {
		baseURL = fmt.Sprintf(baseURL, language)
	}
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	v := url.Query()
	v.Set("action", "query")
	v.Set("prop", "extracts|info|pageimages")
	v.Set("pithumbsize", ThumbSize)
	v.Set("format", "json")
	v.Set("exintro", "")
	v.Set("explaintext", "")
	v.Set("inprop", "url")
	v.Set("redirects", "")
	v.Set("converttitles", "")
	v.Set("titles", query)
	url.RawQuery = v.Encode()

	return &Request{url}, nil
}

// Execute fetches the data and decodes it into a Response.
// Returns an error if the data could not be retrived or the decoding fails.
func (r *Request) Execute(noCheckCert bool) (*Response, error) {
	client := httpClient(noCheckCert)

	data, err := client.Get(r.String())
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	d := json.NewDecoder(data.Body)
	resp := &Response{}
	err = d.Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
