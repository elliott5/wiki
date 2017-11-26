// +build js

package wiki

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var jsonpMagicMtx sync.Mutex

// see https://en.wikipedia.org/wiki/JSONP
// example https://siongui.github.io/2016/01/23/go-jsonp-example-cors-by-gopherjs/
func jsonpMagic(url string) chan string {
	ch := make(chan string)
	go func() {
		jsonpMagicMtx.Lock()
		defer jsonpMagicMtx.Unlock()

		js.Global.Set("jsonpWikiCallback", func(jsonData map[string]interface{}) {
			go func() {
				ch <- js.Global.Get("JSON").Call("stringify", jsonData).String()
			}()
		})

		targetUrl := url + `&callback=jsonpWikiCallback`

		doc := js.Global.Get("document")
		head := doc.Call("getElementsByTagName", "head").Index(0)
		ele := doc.Call("createElement", "script")
		ele.Call("setAttribute", "src", targetUrl)
		head.Call("appendChild", ele)
	}()
	return ch
}

type jsWikiClient struct{}

type jsResp struct {
	Body io.ReadCloser
}

func (*jsWikiClient) Get(url string) (*jsResp, error) {
	select {
	case res := <-jsonpMagic(url):
		return &jsResp{Body: ioutil.NopCloser(strings.NewReader(res))}, nil
	case <-time.After(time.Second * 2): // TIMEOUT
		return nil, errors.New("timeout fetching wikipedia data")
	}
}

func httpClient(noCheckCert bool) *jsWikiClient {
	return &jsWikiClient{}
}

func (r *Request) ExecuteGAE(noCheckCert bool, unused interface{}) (*Response, error) {
	return r.Execute(noCheckCert)
}
