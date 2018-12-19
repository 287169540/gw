
###### http service demo
```$xslt
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

###### http get
```$xslt
package main

import (
	"fmt"
	gw2 "gw"
)

func main() {
	gw := gw2.New("hehe")
	rs := gw.HttpGet("http://www.baidu.com")
	fmt.Println(string(rs))
}

```

###### http post
```cassandraql
package main

import (
	"fmt"
	gw2 "gw"
	"net/url"
)

func main() {
	gw := gw2.New("hehe")
	data:= url.Values{}
	data.Set("name", "rj")
	data.Set("age", "12")
	rs:=gw.HttpPost("http://www.baidu.com",data)
	fmt.Println(string(rs))
}
```