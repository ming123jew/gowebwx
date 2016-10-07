package webwx

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func TimeS() int {
	return int(time.Now().Unix())
}

func TimeMs() int {
	return int(time.Now().UnixNano() / 1e6)
}

func revTimeMs() int {
	t := int(time.Now().UnixNano() / 1e6)
	t = (^t) & (2<<32 - 1)
	return t
}

func genDeviceId() string {
	rand.Seed(time.Now().Unix())
	return "e" + strconv.Itoa(rand.Int())[2:17]
}

func getUserAgent() string {
	// os: android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows
	/*
		switch runtime.GOOS {
		case "android":
		case "darwin":
		case "freebsd", "linux":
		case "windows":
		}
	*/
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36"
}

func defaultQRCodeHandler(b []byte) {
	path := "qrcode.png"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Create qrcode.png failed:", err)
		return
	}
	_, err = io.Copy(file, bytes.NewReader(b))
	if err != nil {
		fmt.Println("Copy to qrcode.png failed:", err)
		return
	}

	if runtime.GOOS == "darwin" {
		exec.Command("open", path).Run()
	} else {

	}
}

func isChatRoomContact(name string) bool {
	return strings.HasPrefix(name, "@@") || strings.HasSuffix(name, "@chatroom")
}

func defaultMsgHandler(m Msg) {
	log.Infof("Message:\n\tfrom: %s\n\tto: %s\n\ttype: %v\n\tcontent: %s\n\ttime: %v\n\tnotify: %s",
		m.FromUserName, m.ToUserName, m.MsgType, m.Content, m.CreateTime, m.StatusNotifyUserName)
}
