
###### http service demo
```
package main

import (
	"fmt"
	gw2 "gw"
)

func hi(gw *gw2.GW) (string) {
	fmt.Printf("get name:%v\n",gw.Get("name"))
	fmt.Printf("post name:%v\n",gw.Post("name"))
	return "hello world"
}

func main() {
	gw := gw2.New("hello")
	gw.BindFunc("/", hi)
	gw.SetStaticFileDir("/static/", "static")
	gw.Run(8082)
}
```