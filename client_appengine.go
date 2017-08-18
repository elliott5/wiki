// +build appengine

package wiki

import (
	"net/http"
	"sync"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

var mtx sync.Mutex
var currentClient *http.Client

func httpClient() *http.Client {
	return currentClient
}

// ExecuteGAE is a quick fix to use this library within AppEngine.
// Using a mutex in this way serialises every call to this func, which is not good practice.
// TODO remove serialisation - will require altering the Request structure.
func (r *Request) ExecuteGAE(noCheckCert bool, req *http.Request) (*Response, error) {
	mtx.Lock()
	defer mtx.Unlock()
	currentClient = urlfetch.Client(appengine.NewContext(req))
	return r.Execute(false /* noCheckCert==true may not work on GAE & is bad practice anyway */)
}
