package gw

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type GW struct {
	R     *http.Request
	logTitle string
	logId string
}

// create new gw
func New(logTitle string)(* GW){
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
func (gw *GW) BindFunc(path string, fn func() (string)) (*GW) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		gw.initLogId()
		request.ParseForm()
		request.ParseMultipartForm(1024 * 1024 * 2)
		gw.R = request
		gw.logR()
		html := fn()
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

// run a server instance
func (gw *GW) Run(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

// get config content
func (gw *GW) GetCfg(path string) (txt string) {

	return ""
}

// GET query data
func (gw *GW) Get(key string) (string) {
	fmt.Println(key)
	return gw.R.URL.Query().Get(key)
}

// POST query data
func (gw *GW) Post(key string) (string) {
	return gw.R.PostForm.Get(key)
}
