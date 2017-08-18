// +build !appengine

package wiki

import "net/http"

func httpClient() *http.Client {
	return &http.Client{}
}

func (r *Request) ExecuteGAE(noCheckCert bool, req *http.Request) (*Response, error) {
	return r.Execute(noCheckCert)
}
