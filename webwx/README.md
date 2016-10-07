### GoWebWx

**!! IN PROGRESS !!**

This is a wechat cmd client written with go.

It utilize the web wechat apis. (https://wx.qq.com)

current apis refs:  [webwxApp308288](https://res.wx.qq.com/zh_CN/htmledition/v2/js/webwxApp308288.js)

### Design (Not Completed!)

1. Can run in main goroutine as ```w.Start(0)```, or in seperated goroutine as ```go w.Start(0)```, the argument means retry times on fail.
2. Hooks :
    - w.OnOnline(func())
    - w.OnOffline(func())
    - w.OnQRCode(func([]byte)) the argument is the image data
    - w.OnOnline(func())
    - w.OnSessionSay(func(msg Msg))
    - w.OnChatRoomSay(func(msg Msg))
3. Actions:
    - w.SendSessionSay(msg Msg)
    - w.SendChatRoomSay(msg Msg)

### Example

Currently default qrcode hook can run on mac only.

```go run example/main.go```

```go
package main

import (
	"fmt"
	"github.com/wuxc/gowebwx/webwx"
)

func main() {
	w := webwx.NewWebWx()
	w.OnOnline(func() {
		fmt.Println("Online!")
	})
	w.OnOffline(func() {
		fmt.Println("Offline!")
	})
	w.Start(0)
}
```

### Refs

- java https://github.com/biezhi/wechat-robot
- python https://github.com/lyyyuna/wechat_robot
- python https://github.com/Urinx/WeixinBot
- python https://github.com/liuwons/wxBot
- go https://github.com/ManiacMike/gorobot
