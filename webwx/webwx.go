package webwx

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type webwx struct {
	uuid         string
	base_host    string //wx.qq.com for default
	login_host   string
	upload_host  string
	webpush_host string
	user_avatar  string
	redirect_uri string
	check_url    string
	uin          string
	sid          string
	skey         string
	pass_ticket  string
	isgrayscale  string
	http_client  *http.Client
	user_agent   string
	sync_time    int
	deviceid     string
	lang         string
	// useful data
	BaseRequest BaseRequest
	SyncKey     SyncKey
	User        User
	Contacts    map[string]Contact
	// status monitoring
	running bool
	err     error
	exit    chan int // controls restart or giveup
	// event hooks
	on_qrcode       func([]byte)
	on_online       func()
	on_offline      func()
	on_exit         func()
	on_session_say  func(Msg)
	on_chatroom_say func(Msg)
}

func (w *webwx) _rawreq(req *http.Request) []byte {
	log.Debug("_rawreq", req.URL)
	for _, c := range w.http_client.Jar.Cookies(req.URL) {
		log.Debug("cookie:", c)
	}
	resp, err := w.http_client.Do(req)
	if err != nil || resp == nil {
		log.Error("err do request", err)
		w.err = err
		return nil
	}
	defer resp.Body.Close()
	h := resp.Request.Host
	if h != w.base_host {
		w.base_host = h
		HostMap.UpdateHost(w)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("err read body", err)
		w.err = err
		return nil
	}
	for _, c := range resp.Cookies() {
		log.Debug("set cookie:", c)
	}
	return body
}

func (w *webwx) _strreq(req *http.Request) string {
	b := w._rawreq(req)
	if b == nil {
		return ""
	}
	return string(b)
}

func (w *webwx) getuuid() {
	r := w.api_getuuid()
	s := w._strreq(r)
	p := regexp.MustCompile(`"([\S]+)"`)
	match := p.FindStringSubmatch(s)
	if len(match) == 2 {
		w.uuid = match[1]
	} else {
		w.err = errors.New("Get uuid failed: " + s)
	}
}

func (w *webwx) getqrcode() {
	w._checkerror()
	r := w.api_getqrcode()
	b := w._rawreq(r)
	if b != nil && w.on_qrcode != nil {
		w.on_qrcode(b)
	}
}

func (w *webwx) checklogin(tip int) int {
	w._checkerror()
	r := w.api_checklogin(tip)
	s := w._strreq(r)
	p := regexp.MustCompile(`window.code=(\d+);`)
	match := p.FindStringSubmatch(s)
	if len(match) == 2 {
		code, err := strconv.Atoi(match[1])
		if err != nil {
			return -1
		}
		if code == 200 { // redirect
			p = regexp.MustCompile(`window.redirect_uri="([\S]+)"`)
			match = p.FindStringSubmatch(s)
			if len(match) > 1 {
				w.redirect_uri = match[1]
			}
		} else if code == 201 { // useAvatar
			p = regexp.MustCompile(`window.userAvatar = '([\S]+)'`)
			match = p.FindStringSubmatch(s)
			if len(match) > 1 {
				w.user_avatar = match[1]
			}
		}
		return code
	}
	return -1
}

func (w *webwx) login_redirect() {
	w._checkerror()
	r := w.api_redirectpage()
	log.Debug("redirect to:", r.URL)
	s := w._strreq(r)
	log.Debug("after redirect:", s)
	p := regexp.MustCompile(`<error><ret>(\d+)</ret><message></message><skey>([\S]+)</skey><wxsid>([\S]+)</wxsid><wxuin>(\d+)</wxuin><pass_ticket>([\S]+)</pass_ticket><isgrayscale>(\d+)</isgrayscale></error>`)
	match := p.FindStringSubmatch(s)
	if len(match) == 7 {
		ret := match[1]
		w.skey = match[2]
		w.sid = match[3]
		w.uin = match[4]
		w.pass_ticket = match[5]
		w.isgrayscale = match[6]
		if ret != "0" {
			log.Warn("ret is not zero", ret)
		}
	} else {
		w.err = errors.New("redirect data error: " + s)
	}
}

func (w *webwx) webwxinit() {
	w._checkerror()
	w.BaseRequest = BaseRequest{
		Uin:      w.uin,
		Sid:      w.sid,
		Skey:     w.skey,
		DeviceID: w.deviceid,
	}
	r := w.api_webwxinit()
	b := w._rawreq(r)
	var resp InitResponse
	err := json.Unmarshal(b, &resp)
	if err != nil {
		log.Error("webwxinit unmarshal error:", err)
		ioutil.WriteFile("_debug_init.json", b, 400)
		w.err = err
		return
	}

	b, _ = json.Marshal(resp)
	log.Debug("webwxinit resp:", string(b))
	switch resp.BaseResponse.Ret {
	case 0:
	case 1100, 1101, 1102:
		w.err = errors.New("init Login on other device!")
		return
	case 1205:
		w.err = errors.New("init 1205")
		return
	}

	w.User = resp.User
	w.skey = resp.SKey
	w.SyncKey = resp.SyncKey

	// add Contacts
	// ignore u.addContact(n.User),
	// w._addcontact(resp.User)
	for _, c := range resp.ContactList {
		w._addcontact(c)
	}
	// initChatList, only the part where chatroom memebers are added is useful
	for _, chat := range strings.Split(resp.ChatSet, ",") {
		if isChatRoomContact(chat) {
			log.Debug("Get Chatroom members:", chat)
		}
	}
	// getChatList, sort chatlist, no use

	// notifyMobile StatusNotifyCode_INITED
	w.webwxstatusnotify(w.User.UserName, StatusNotifyCode_INITED)

	// 6050 MPSubscribeMsgList init, seems webwx never gets MPSubscripteMsg?
	// v.init(n.MPSubscribeMsgList),

	// setCheckUrl no use
	// w.check_url = fmt.Sprintf("&skey=%s&deviceid=%s&pass_ticket=%s&opcode=2&scene=1&username=%s", w.skey, w.deviceid, w.pass_ticket, w.User.UserName)

	// init contact 0
	w.webwxgetcontact(0)
	// 16 times until o.Seq is 0

	// start synccheck

}

func (w *webwx) webwxlogout(typ int) {
	return
	w._checkerror()
	r := w.api_webwxlogout(typ)
	log.Debug("logout url", r.URL)
	w.http_client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	defer func() {
		w.http_client.CheckRedirect = nil
	}()
	resp, err := w.http_client.Do(r)
	if err != nil {
		log.Error("webwxlogout error", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("logout read body error", err)
		w.err = err
		return
	}
	log.Debug("webwxlogout", string(body))
}

func (w *webwx) _addcontact(c Contact) {
	if w.Contacts == nil {
		w.Contacts = make(map[string]Contact)
	}
	w.Contacts[c.UserName] = c
}

func (w *webwx) webwxstatusnotify(to string, code int) {
	w._checkerror()
	r := w.api_webwxstatusnotify(to, code)
	_ = w._strreq(r)
	// we don't need the response
	// var resp StatusNotifyResponse
	// err := json.Unmarshal(b, &resp)
}

func (w *webwx) webwxgetcontact(seq int) {
	w._checkerror()
	r := w.api_webwxgetcontact(seq)
	// add memberlist addcontacts
	b := w._rawreq(r)
	var resp GetContactResponse
	err := json.Unmarshal(b, &resp)
	if err != nil {
		w.err = err
		return
	}
	for _, c := range resp.MemberList {
		log.Debug(c)
	}
}

func (w *webwx) webwxbatchgetcontact(usernames []string) {
	w._checkerror()
	r := w.api_webwxbatchgetcontact(usernames)
	b := w._rawreq(r)
	var resp BatchGetContactResponse
	err := json.Unmarshal(b, &resp)
	if err != nil {
		w.err = err
		return
	}
	for _, c := range resp.ContactList {
		log.Debug(c)
	}
}

func (w *webwx) synccheck() (retcode, selector string) {
	w._checkerror()
	log.Info("webwx checking..")
	r := w.api_synccheck()
	s := w._strreq(r)
	log.Debug("Synccheck:", s)
	p := regexp.MustCompile(`\{retcode:"(\d+)"\s*,\s*selector:"(\d+)"}`)
	match := p.FindStringSubmatch(s)
	if len(match) == 3 {
		retcode = match[1]
		selector = match[2]
	} else {
		w.err = errors.New("sync check failed: " + s)
		retcode = ""
		selector = ""
	}
	return
}

func (w *webwx) webwxsync() {
	w._checkerror()
	r := w.api_webwxsync()
	b := w._rawreq(r)
	var resp SyncResponse
	err := json.Unmarshal(b, &resp)
	if err != nil {
		log.Error("webwxsync unmarshal failed.", err)
		w.err = err
		return
	}
	// f.setSyncKey(e.SyncKey),
	if resp.SyncKey.Count > 0 {
		w.SyncKey = resp.SyncKey
	}
	// f.updateUserInfo(e.Profile, function() {}),
	p := resp.Profile
	if p.BitFlag == PROFILE_BITFLAG_CHANGE {
		if p.HeadImgUpdateFlag != 0 {
			w.User.HeadImgUrl = p.HeadImgUrl
		}
		if p.NickName.Buff != "" {
			w.User.NickName = p.NickName.Buff
		}
	}

	// handle Contact change
	for _, c := range resp.DelContactList {
		log.Debug("Del Contact:", c)
		delete(w.Contacts, c.UserName)
	}
	for _, c := range resp.ModContactList {
		log.Debug("Add Contact:", c)
		//w._addcontact(c)
	}
	// handle ChatRoomMember change
	for _, c := range resp.ModChatRoomMemberList {
		log.Debug("ChatRoomUpdate:", c)
	}
	// update Profile
	// handle Msg
	for _, m := range resp.AddMsgList {
		if isChatRoomContact(m.ToUserName) {
			w.on_chatroom_say(m)
		} else {
			w.on_session_say(m)
		}
	}
}

func (w *webwx) webwxsendmsg() {
}

func (w *webwx) _checkerror() {
	if w.err != nil {
		log.Critical("err detected! stopping..", w.err)
		w.exit <- 1
		runtime.Goexit()
	}
}

// events and api
func NewWebWx() *webwx {
	jar, _ := cookiejar.New(nil)
	client := http.Client{
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       35 * time.Second,
	}
	w := &webwx{
		base_host:       "wx.qq.com",
		lang:            "zh_CN",
		deviceid:        genDeviceId(),
		user_agent:      getUserAgent(),
		http_client:     &client,
		running:         false,
		exit:            make(chan int),
		on_qrcode:       defaultQRCodeHandler,
		on_session_say:  defaultMsgHandler,
		on_chatroom_say: defaultMsgHandler,
	}
	HostMap.UpdateHost(w)
	return w
}

func (w *webwx) reset() {
	w.err = nil
	w.sync_time = 0
}

func (w *webwx) _start() {
	defer func() {
		log.Info("defered Done.")
	}()

	w.reset()
	w.getuuid()
	w.getqrcode()

	login_retries := 3
	code := w.checklogin(1)

CHECK_LOGIN_LOOP:
	for {
		switch code {
		case 200:
			w.sync_time = 0
			break CHECK_LOGIN_LOOP
		case 201, 408:
			code = w.checklogin(0)
		case 400, 500, 0:
			if login_retries < 0 {
				break CHECK_LOGIN_LOOP
			}
			login_retries -= 1
			time.Sleep(5 * time.Second)
		case -1:
			break CHECK_LOGIN_LOOP
		}
	}

	w.login_redirect()
	w.webwxinit()

	if w.on_online != nil {
		w.on_online()
	}

CHECK_SYNC_LOOP:
	for {
		retcode, selector := w.synccheck()
		switch retcode {
		case "0":
			if selector != "0" {
				w.webwxsync()
			}
		case "1100":
			log.Debug("loop Login on other device!")
			w.webwxlogout(0)
			break CHECK_SYNC_LOOP
		case "1101", "1102":
			w.webwxlogout(1)
			break CHECK_SYNC_LOOP
		default:
			time.Sleep(25 * time.Second)
		}
	}

	if w.on_offline != nil {
		w.on_offline()
	}

	// no error occurred till now
	// tell main goroutine to stop
	w.exit <- 1
}

func (w *webwx) Start() {
	go w._start()
	<-w.exit
	log.Info("exit")
	// TODO: maybe do some cleanup

	if w.on_exit != nil {
		w.on_exit()
	}
}

func (w *webwx) Stop() {
	log.Info("Stop!")
}

func (w *webwx) Offline() {
	// TODO
}

func (w *webwx) SendSessionSay() {
	// TODO
}

func (w *webwx) SendChatRoomSay() {
	// TODO
}

func (w *webwx) OnQRCode(cb func([]byte)) {
	if cb == nil {
		return
	}
	w.on_qrcode = cb
}

func (w *webwx) OnOnline(cb func()) {
	if cb == nil {
		return
	}
	w.on_online = cb
}

func (w *webwx) OnOffline(cb func()) {
	if cb == nil {
		return
	}
	w.on_offline = cb
}

func (w *webwx) OnExit(cb func()) {
	if cb == nil {
		return
	}
	w.on_exit = cb
}

func (w *webwx) OnSessionSay(cb func(msg Msg)) {
	if cb == nil {
		return
	}
	w.on_session_say = cb
}

func (w *webwx) OnChatRoomSay(cb func(msg Msg)) {
	if cb == nil {
		return
	}
	w.on_chatroom_say = cb
}
