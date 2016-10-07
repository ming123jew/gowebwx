package webwx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func updateRequestHeader(w *webwx, r *http.Request, accept string) {
	// pity go/http.Client does not support these encodings..
	// GET  r.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// POST r.Header.Set("Accept-Encoding", "gzip, deflate, sdch, br")
	r.Header.Set("User-Agent", w.user_agent)
	r.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2")
	r.Header.Set("Referer", "https://"+w.base_host+"/")
	if r.Method == "POST" {
		r.Header.Set("Origin", "https://"+w.base_host)
	}
	if accept == "json" {
		r.Header.Set("Accept", "application/json, text/plain, */*")
	} else if accept == "image" {
		r.Header.Set("Accept", "image/webp,image/*,*/*;q=0.8") //changes!
	} else {
		r.Header.Set("Accept", "*/*")
	}
}

func (w *webwx) api_getuuid() *http.Request {
	redirect := url.QueryEscape(fmt.Sprintf("https://%s%s", w.base_host, API_redirecturi))
	u := fmt.Sprintf("https://%s%s&redirect_uri=%s&lang=%s&_=%v", w.login_host, API_jsLogin, redirect, w.lang, TimeMs())
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "*")
	return r
}

func (w *webwx) api_getqrcode() *http.Request {
	u := fmt.Sprintf("https://%s/qrcode/%s", w.login_host, w.uuid)
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "image")
	return r
}

func (w *webwx) api_checklogin(tip int) *http.Request {
	if w.sync_time == 0 {
		w.sync_time = TimeMs()
	} else {
		w.sync_time += 1
	}
	u := fmt.Sprintf("https://%s%s?loginicon=true&uuid=%s&tip=%v&r=%v&_=%v", w.login_host, API_login, w.uuid, tip, revTimeMs(), w.sync_time)
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "*")
	return r
}

func (w *webwx) api_redirectpage() *http.Request {
	u := w.redirect_uri + "&fun=new&version=v2"
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "*")
	r.Header.Set("Upgrade-Insecure-Requests", "1")
	r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	return r
}

func (w *webwx) api_webwxinit() *http.Request {
	u := fmt.Sprintf("https://%s%s?r=%v", w.base_host, API_webwxinit, revTimeMs())
	if w.pass_ticket != "" {
		u += "&pass_ticket=" + url.QueryEscape(w.pass_ticket)
	}
	b, _ := json.Marshal(InitRequest{w.BaseRequest})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Upgrade-Insecure-Requests", "1")
	return r
}

func (w *webwx) api_synccheck() *http.Request {
	if w.sync_time == 0 {
		w.sync_time = TimeMs()
	} else {
		w.sync_time += 1
	}
	u := fmt.Sprintf("https://%s%s?r=%v&skey=%s&sid=%s&uin=%s&deviceid=%s&synckey=%s&_=%v",
		w.webpush_host, API_synccheck, TimeMs(), url.QueryEscape(w.skey), w.sid, w.uin, w.deviceid, w.SyncKey.format(), w.sync_time)
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "*")
	return r
}

func (w *webwx) api_webwxsync() *http.Request {
	u := fmt.Sprintf("https://%s%s?", w.base_host, API_webwxsync)
	b, _ := json.Marshal(SyncRequest{
		BaseRequest: w.BaseRequest,
		SyncKey:     w.SyncKey,
		rr:          revTimeMs(),
	})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxstatusnotify(to string, code int) *http.Request {
	//https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxstatusnotify
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxstatusnotify, w.pass_ticket)
	b, _ := json.Marshal(StatusNotifyRequest{
		BaseRequest:  w.BaseRequest,
		ClientMsgId:  TimeMs(),
		Code:         code,
		FromUserName: w.User.UserName,
		ToUserName:   to,
	})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxgetcontact(seq int) *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s&r=%v&seq=%v&skey=%s", w.base_host, API_webwxgetcontact, w.pass_ticket, TimeMs(), seq, url.QueryEscape(w.skey))
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "*")
	return r
}

func (w *webwx) api_webwxbatchgetcontact(usernames []string) *http.Request {
	u := fmt.Sprintf("https://%s%s?type=ex&r=%v&pass_ticket=%s", w.base_host, API_webwxbatchgetcontact, TimeMs(), w.pass_ticket)
	b, _ := json.Marshal(BatchGetContactRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxgeticon(seq int, username string) *http.Request {
	u := fmt.Sprintf("https://%s%s?seq=%v&username=%s&skey=%s", w.base_host, API_webwxgeticon, seq, username, url.QueryEscape(w.skey))
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "image")
	return r
}

func (w *webwx) api_webwxgetheadimg(seq int, username string) *http.Request {
	u := fmt.Sprintf("https://%s%s?seq=%v&username=%s&skey=%s", w.base_host, API_webwxgetheadimg, seq, username, url.QueryEscape(w.skey))
	r, _ := http.NewRequest("GET", u, nil)
	updateRequestHeader(w, r, "image")
	return r
}

func (w *webwx) api_webwxupdatechatroom(mod, room, members string) *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxupdatechatroom, w.pass_ticket)
	var b []byte
	if mod == "add" {
		b, _ = json.Marshal(UpdateChatRoomAddRequest{BaseRequest: w.BaseRequest, ChatRoomName: room, AddMemberList: members})
	} else {
		b, _ = json.Marshal(UpdateChatRoomDelRequest{BaseRequest: w.BaseRequest, ChatRoomName: room, DelMemberList: members})
	}
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

// TODO API_webwxdownloadmedia
func (w *webwx) api_webwxdownloadmedia() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.upload_host, API_webwxdownloadmedia, w.pass_ticket)
	r, _ := http.NewRequest("GET", u, nil)
	return r
}

// TODO API_webwxuploadmedia
// need first OPTION then POST
func (w *webwx) api_webwxuploadmedia(method string) *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.upload_host, API_webwxuploadmedia, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "*")
	if method == "OPTIONS" {
		r.Header.Set("Access-Control-Request-Method", "POST")
		r.Header.Set("Access-Control-Request-Headers", "")
	}
	// TODO wrong content type
	// r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	//Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryCg0ZCaBdxraXFCtS
	return r
}

func (w *webwx) api_webwxsendmsg() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxsendmsg, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxsendmsgimg() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxsendmsgimg, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxsendmsgvedio() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxsendmsgvedio, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

func (w *webwx) api_webwxsendemoticon() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxsendemoticon, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}
func (w *webwx) api_webwxsendappmsg() *http.Request {
	u := fmt.Sprintf("https://%s%s?pass_ticket=%s", w.base_host, API_webwxsendappmsg, w.pass_ticket)
	b, _ := json.Marshal(SendMsgRequest{})
	rd := strings.NewReader(string(b))
	r, _ := http.NewRequest("POST", u, rd)
	updateRequestHeader(w, r, "json")
	r.Header.Set("Content-Type", "application/json;charset=UTF-8")
	return r
}

// TODO API_webwxlogout
func (w *webwx) api_webwxlogout(typ int) *http.Request {
	u := fmt.Sprintf("https://%s%s?redirect=1&type=%v&skey=%s", w.base_host, API_webwxlogout, typ, url.QueryEscape(w.skey))
	formdata := fmt.Sprintf("sid=%s&uin=%s", w.sid, w.uin)
	rd := strings.NewReader(formdata)
	r, err := http.NewRequest("POST", u, rd)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	updateRequestHeader(w, r, "*")
	r.Header.Set("Upgrade-Insecure-Requests", "1")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	return r
}

// TODO API_webwxpreview
// TODO API_webwxgetmsgimg
// TODO API_webwxgetmedia
// TODO API_webwxgetvideo
// TODO API_webwxgetvoice
// TODO API_webwxcreatechatroom
// TODO API_webwxcheckurl
// TODO API_webwxverifyuser
// TODO API_webwxfeedback
// TODO API_webwxreport
// TODO API_webwxsearch
// TODO API_webwxoplog
// TODO API_checkupload
