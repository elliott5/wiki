// +build !appengine,!js

package wiki

import "net/http"
import "crypto/tls"

func httpClient(noCheckCert bool) *http.Client {
	if noCheckCert {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return &http.Client{Transport: tr}
	}
	return &http.Client{}
}

func (r *Request) ExecuteGAE(noCheckCert bool, unused interface{}) (*Response, error) {
	return r.Execute(noCheckCert)
}
