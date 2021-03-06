package gw

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

type GW struct {
	R        *http.Request
	logTitle string
	logId    string
}

// create new gw
func New(logTitle string) (*GW) {
	gw := new(GW)
	gw.logTitle = logTitle
	return gw
}

// initilize log id
func (gw *GW) initLogId() {
	gw.logId = fmt.Sprintf("%v %v %v",
		time.Now().Format("2006-01-02 15:04:05"),
		rand.Int(),
		gw.logTitle,
	)
}

// set handler func
func (gw *GW) BindFunc(path string, fn func(gw *GW) (string)) (*GW) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		gw.initLogId()
		request.ParseForm()
		request.ParseMultipartForm(1024 * 1024 * 2)
		gw.R = request
		gw.logR()
		html := fn(gw)
		gw.logW(html)
		fmt.Fprint(writer, html)
	})
	return gw
}

// write log for input
func (gw *GW) logR() {
	fmt.Printf("%v input % #v \n", gw.logId, gw.R)
}

// write log
func (gw *GW) Log(log string) {
	fmt.Printf("%v log %v \n", gw.logId, log)
}

// write log output log
func (gw *GW) logW(log string) {
	fmt.Printf("%v output %v \n", gw.logId, log)
}

// set static file service
func (gw *GW) SetStaticFileDir(url string, dir string) {
	fs := http.FileServer(http.Dir(dir))
	http.Handle(url, http.StripPrefix(url, fs))
}

// run a server instance
func (gw *GW) Run(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

// get config content
func (gw *GW) GetCfg(path string) ([]byte) {
	if filename, err := filepath.Abs(path); nil == err {
		if bytes, err := ioutil.ReadFile(filename); nil == err {
			return bytes
		}
	}
	return []byte("")
}

// GET query data
func (gw *GW) Get(key string) (string) {
	return gw.R.URL.Query().Get(key)
}

// POST query data
func (gw *GW) Post(key string) (string) {
	return gw.R.PostForm.Get(key)
}

// Http Get
func (gw *GW) HttpGet(url string) ([]byte) {
	res, err := http.Get(url)
	if nil == err {
		bytes, err := ioutil.ReadAll(res.Body)
		if nil == err {
			return bytes
		}
	}
	return []byte("")
}

// Curl Post
func (gw *GW) HttpPost(url string, data url.Values) ([]byte) {
	if res, err := http.PostForm(url, data); nil == err {
		if bytes, err := ioutil.ReadAll(res.Body); nil == err {
			return bytes;
		}
	}
	return []byte("");
}
