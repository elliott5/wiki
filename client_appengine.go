// +build appengine

package wiki

import (
	"context"
	"net/http"
	"sync"
	"time"

	"google.golang.org/appengine/urlfetch"
)

var mtx sync.Mutex
var currentClient *http.Client

func httpClient(noCheckCert bool) *http.Client {
	return currentClient
}

// ExecuteGAE is a quick fix to use this library within AppEngine.
// Using a mutex in this way serialises every call to this func, which is not good practice.
// TODO remove serialisation - will require altering the Request structure.
func (r *Request) ExecuteGAE(noCheckCert bool, ctxt context.Context) (*Response, error) {
	mtx.Lock()
	defer mtx.Unlock()
	ctxTO, _ := context.WithTimeout(ctxt, 60*time.Second)
	currentClient = urlfetch.Client(ctxTO)
	return r.Execute(false /* noCheckCert==true may not work on GAE & is bad practice anyway */)
}
