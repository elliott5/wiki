// +build !appengine

package wiki

import "net/http"

func httpClient() *http.Client {
	return &http.Client{}
}

func (r *Request) ExecuteGAE(noCheckCert bool, unused interface{}) (*Response, error) {
	return r.Execute(noCheckCert)
}
