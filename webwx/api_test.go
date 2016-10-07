package webwx

import (
	"bytes"
	"net/http"
	"testing"
)

var w = NewWebWx()

func _testApi(t *testing.T, s string, r *http.Request, debug bool) {
	buf := new(bytes.Buffer)
	r.Write(buf)
	if debug {
		t.Log(s)
		t.Log("\n" + buf.String())
	}
}

func TestApiGetUUID(t *testing.T) {
	s := `
GET /jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1475646433988 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
`
	HostMap.UpdateHost(w)
	r := w.api_getuuid()
	_testApi(t, s, r, false)
}

func TestApiGetQRcode(t *testing.T) {
	s := `
GET /qrcode/4Zok23Gt9g== HTTP/1.1
Host: login.weixin.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: ts_uid=4736163348; pgv_pvi=941454336; pgv_si=s7109208064
`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_getqrcode()
	_testApi(t, s, r, false)
}

func TestApiCheckLogin(t *testing.T) {
	s := `
GET /cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=AbdJkHwA7A==&tip=0&r=1779066869&_=1475689655464 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; pgv_si=s7109208064; mm_lang=zh_CN
`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_checklogin(0)
	_testApi(t, s, r, false)
}

func TestApiLogout(t *testing.T) {
	HostMap.UpdateHost(w)
	w.BaseRequest = BaseRequest{Uin: "1377554769", Sid: "JkpTzkoC7p3Qgo+n", Skey: "", DeviceID: "e972742953532257"}
	w.skey = w.BaseRequest.Skey
	w.uin = w.BaseRequest.Uin
	w.sid = w.BaseRequest.Sid
	r := w.api_webwxlogout(0)
	if r == nil {
		return
	}
	_testApi(t, "", r, true)
}

func TestApiRedirectPage(t *testing.T) {
	s := `
GET /cgi-bin/mmwebwx-bin/webwxnewloginpage?ticket=A7vxoLJ5L2tQzOhHQneveB-O@qrticket_0&uuid=wdgZ_eNqIQ==&lang=zh_CN&scan=1475689970 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; wxloadtime=1475651408_expired; wxpluginkey=1475629201; wxuin=1377554769; mm_lang=zh_CN; pgv_si=s7109208064
`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_redirectpage()
	_testApi(t, s, r, false)
}

func TestApiWebwxInit(t *testing.T) {
	s := `
POST /cgi-bin/mmwebwx-bin/webwxinit?r=1778764354 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 101
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; wxpluginkey=1475629201; pgv_si=s7109208064; wxuin=1377554769; wxsid=3pQAasI5n1aX+/yt; wxloadtime=1475689984; webwx_data_ticket=gScZp3KTjahPPGJV+b5rwbfp; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1


{"BaseRequest":{"Uin":"1377554769","Sid":"3pQAasI5n1aX+/yt","Skey":"","DeviceID":"e750820437849651"}}
`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_webwxinit()
	_testApi(t, s, r, false)
}

func TestApiSynccheck(t *testing.T) {
	s := `
GET /cgi-bin/mmwebwx-bin/synccheck?r=1475717651164&skey=%40crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853&sid=Wsl9pv7GwTkXPwbs&uin=1377554769&deviceid=e282022489800533&synckey=1_660260314%7C2_660260644%7C3_660260635%7C1000_1475716321&_=1475717632048 HTTP/1.1
Host: webpush.wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: */*
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_synccheck()
	_testApi(t, s, r, false)
}

func TestApiSync(t *testing.T) {
	s := `
POST /cgi-bin/mmwebwx-bin/webwxsync?sid=Wsl9pv7GwTkXPwbs&skey=@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853&lang=zh_CN&pass_ticket=AtB7R79rH3dlBIea%252Fw3EpKKWwNItjWAG13WrfofHIuTJgrbL0H%252BOQu2azV7HVrAB HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 302
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
	_ = `{"BaseRequest":{"Uin":1377554769,"Sid":"gAumbqNzbguPuoV7","Skey":"@crypt_24c998a9_cbe09e42a6302994a960a9a61c1a5ee7","DeviceID":"e247297082389616"},"SyncKey":{"Count":4,"List":[{"Key":1,"Val":660260712},{"Key":2,"Val":660260710},{"Key":3,"Val":660260635},{"Key":1000,"Val":1475716321}]},"rr":1736813075}`
	HostMap.UpdateHost(w)
	w.uuid = "4Zok23Gt9g=="
	r := w.api_webwxsync()
	_testApi(t, s, r, false)
}

func TestApiStatusNotify(t *testing.T) {
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxstatusnotify?lang=zh_CN&pass_ticket=AtB7R79rH3dlBIea%252Fw3EpKKWwNItjWAG13WrfofHIuTJgrbL0H%252BOQu2azV7HVrAB HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 348
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
	_ = `{"BaseRequest":{"Uin":1377554769,"Sid":"gAumbqNzbguPuoV7","Skey":"@crypt_24c998a9_cbe09e42a6302994a960a9a61c1a5ee7","DeviceID":"e564833755593153"},"Code":3,"FromUserName":"@f477f4eb35d71ae4da97bf92292465f23452593f4c444f4b743b476b32eec94d","ToUserName":"@f477f4eb35d71ae4da97bf92292465f23452593f4c444f4b743b476b32eec94d","ClientMsgId":1475731936539}`
}
func TestApiGetContact(t *testing.T) {
	_ = `
GET /cgi-bin/mmwebwx-bin/webwxgetcontact?lang=zh_CN&pass_ticket=AtB7R79rH3dlBIea%252Fw3EpKKWwNItjWAG13WrfofHIuTJgrbL0H%252BOQu2azV7HVrAB&r=1475717651159&seq=0&skey=@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Accept: application/json, text/plain, */*
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
}
func TestApiGetChatroomMember(t *testing.T) {
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxbatchgetcontact?type=ex&r=1475717651287&lang=zh_CN&pass_ticket=AtB7R79rH3dlBIea%252Fw3EpKKWwNItjWAG13WrfofHIuTJgrbL0H%252BOQu2azV7HVrAB HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 362
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
	_ = `{"BaseRequest":{"Uin":1377554769,"Sid":"Wsl9pv7GwTkXPwbs","Skey":"@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853","DeviceID":"e092970549368470"},"Count":2,"List":[{"UserName":"@@da560f7dbff80e23f6ccbbdfcc974bfe49ec9f68b44df6760d4c3586a40ca276","ChatRoomId":""},{"UserName":"@@77f72bfd2dcea21314f75a418a7dbcfc1a4e6f06b9ba6cc1f75e800bf62839be","ChatRoomId":""}]}`
}

func TestApiGeticon(t *testing.T) {
	_ = `
GET /cgi-bin/mmwebwx-bin/webwxgeticon?seq=660260034&username=@6fccf38a103a51e3f3ded5d235896891&skey=@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxpluginkey=1475659982; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; wxloadtime=1475717650; mm_lang=zh_CN; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD
`
}

func TestApiGetHeadimg(t *testing.T) {
	_ = `
GET /cgi-bin/mmwebwx-bin/webwxgetheadimg?seq=0&username=@@77f72bfd2dcea21314f75a418a7dbcfc1a4e6f06b9ba6cc1f75e800bf62839be&skey= HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; mm_lang=zh_CN; wxloadtime=1475717650_expired; wxpluginkey=1475716321; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD`
}

func TestApiSendmsg(t *testing.T) {
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxsendmsg?lang=zh_CN&pass_ticket=AtB7R79rH3dlBIea%252Fw3EpKKWwNItjWAG13WrfofHIuTJgrbL0H%252BOQu2azV7HVrAB HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 406
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/?&lang=zh_CN
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; mm_lang=zh_CN; wxloadtime=1475717650_expired; wxpluginkey=1475716321; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD`
	_ = `{"BaseRequest":{"Uin":1377554769,"Sid":"Wsl9pv7GwTkXPwbs","Skey":"@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853","DeviceID":"e299263287798257"},"Msg":{"Type":1,"Content":"[疑问]在哪儿呢？","FromUserName":"@344d8745f9120f842da8eee2c3acbdcaec7ff94e2503ddbc3c8ee30078a517d5","ToUserName":"@6fccf38a103a51e3f3ded5d235896891","LocalID":"14757186177790621","ClientMsgId":"14757186177790621"},"Scene":0}`
	//https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxsendappmsg?fun=async&f=json
	_ = `
{
	"BaseResponse": {
		"Ret": 0,
		"ErrMsg": ""
	}
	,
	"MsgID": "5678434366221810387",
	"LocalID": "14756424267470017"
}
`
}

func TestApiUpdateChatroom(t *testing.T) {
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxupdatechatroom?fun=addmember HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 283
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; wxloadtime=1475717650_expired; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxpluginkey=1475716321; wxuin=1377554769; wxsid=Wsl9pv7GwTkXPwbs; webwx_data_ticket=gSdbDVCp2AMFzDZ8Uk6oDSdD`
	_ = `{"AddMemberList":"@6fccf38a103a51e3f3ded5d235896891","ChatRoomName":"@@77f72bfd2dcea21314f75a418a7dbcfc1a4e6f06b9ba6cc1f75e800bf62839be","BaseRequest":{"Uin":1377554769,"Sid":"Wsl9pv7GwTkXPwbs","Skey":"@crypt_24c998a9_d8290312784e2bc0324dd6bf9a1d2853","DeviceID":"e346749108079273"}}`
}

func TestApiUploadMedia(t *testing.T) {
	//https://file.wx2.qq.com/cgi-bin/mmwebwx-bin/webwxuploadmedia?f=json
	_ = `
OPTIONS /cgi-bin/mmwebwx-bin/webwxuploadmedia?f=json HTTP/1.1
Host: file.wx2.qq.com
Connection: keep-alive
Access-Control-Request-Method: POST
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Access-Control-Request-Headers:
Accept: */*
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
`
	_ = `{"BaseResponse": {"Ret": 1,"ErrMsg": ""},"MediaId": "","StartPos": 0,"CDNThumbImgHeight": 0,"CDNThumbImgWidth": 0}`
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxuploadmedia?f=json HTTP/1.1
Host: file.wx2.qq.com
Connection: keep-alive
Content-Length: 2067
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryCg0ZCaBdxraXFCtS
Accept: */*
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
`
	_ = `
{
	"BaseResponse": {
		"Ret": 0,
		"ErrMsg": ""
	}
	,
	"MediaId": "@crypt_e702b36_9cb04f38a65571eadfa5d5f9c294127e59305bd4df1c855d9ef82a1ca84a54147c3d47f469d668174ecf5e035b04afc62b260bb69b50244648d8153b61ca0e4a9bbe007bd989ef32b07a98d8eb8f0c26",
	"StartPos": 473,
	"CDNThumbImgHeight": 0,
	"CDNThumbImgWidth": 0
}
`
}

func TestSendEmoticon(t *testing.T) {
	_ = `
POST /cgi-bin/mmwebwx-bin/webwxsendemoticon?fun=sys&pass_ticket=EXMR%252F9mNqksyQiBwNC0%252FpLJrKH1IsfIHokYQekM81NoSYyUvhJI%252FPQvpK9SweJAP HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 434
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; mm_lang=zh_CN; wxloadtime=1475731936_expired; wxpluginkey=1475716321; wxuin=1377554769; wxsid=gAumbqNzbguPuoV7; webwx_data_ticket=gSe1f4/QL//3cUyR8ziPCklx
`
	_ = `{"BaseRequest":{"Uin":1377554769,"Sid":"gAumbqNzbguPuoV7","Skey":"@crypt_24c998a9_cbe09e42a6302994a960a9a61c1a5ee7","DeviceID":"e871417385623582"},"Msg":{"Type":47,"EmojiFlag":2,"EMoticonMd5":"846f30447c5c4c9beefeb5a61bec0ba3","FromUserName":"@f477f4eb35d71ae4da97bf92292465f23452593f4c444f4b743b476b32eec94d","ToUserName":"@b7f9e7ed16591730d6d1f8ec9599aee4","LocalID":"14757353172040475","ClientMsgId":"14757353172040475"},"Scene":0}`
}

func TestGetMsgimg(t *testing.T) {
	_ = `
GET /cgi-bin/mmwebwx-bin/webwxgetmsgimg?&MsgID=739166147427210376&skey=%40crypt_24c998a9_cbe09e42a6302994a960a9a61c1a5ee7&type=big HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=941454336; webwxuvid=53115eb0931978ccfafdb4f55d5d382ace2255f59b34b74c057e4403890a492044ed5bb5de0fc4b996ff24260c951927; pgv_si=s7109208064; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; mm_lang=zh_CN; wxloadtime=1475731936_expired; wxpluginkey=1475716321; wxuin=1377554769; wxsid=gAumbqNzbguPuoV7; webwx_data_ticket=gSe1f4/QL//3cUyR8ziPCklx
`
}
