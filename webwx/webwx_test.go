package webwx

import (
	"testing"
)

func Test(t *testing.T) {
	// getuuid
	url := "https://login.wx.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1475625516716"
	s := `
GET /jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1475625516716 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: mm_lang=zh_CN
	`
	s = `
GET /jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1475625920932 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=8288752640; pgv_si=s4345098240; mm_lang=zh_CN
`
	// login
	s = `
GET /cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=YZZ1O_LLMg==&tip=1&r=1842828267&_=1475625920933 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=8288752640; pgv_si=s4345098240; mm_lang=zh_CN
`
	// qrcode
	s = `
GET /qrcode/YZZ1O_LLMg== HTTP/1.1
Host: login.weixin.qq.com
Connection: keep-alive
Cache-Control: max-age=0
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: ts_uid=4736163348; pgv_pvi=8288752640; pgv_si=s4345098240
`

	// waitforlogin
	s = `
GET /cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=YZZ1O_LLMg==&tip=0&r=1842802886&_=1475625920934 HTTP/1.1
Host: login.wx.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: pgv_pvi=8288752640; pgv_si=s4345098240; mm_lang=zh_CN
`

	// redicret
	s = `
GET /cgi-bin/mmwebwx-bin/webwxnewloginpage?ticket=AwV8CbwEyrfjNIANWGiB2gcV@qrticket_0&uuid=YZZ1O_LLMg==&lang=zh_CN&scan=1475626042 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Referer: https://wx.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; wxuin=1377554769; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; pgv_pvi=8288752640; pgv_si=s4345098240
`

	// init
	s = `
POST /cgi-bin/mmwebwx-bin/webwxinit?r=1842684975 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 101
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`
	// notify
	s = `
POST /cgi-bin/mmwebwx-bin/webwxstatusnotify HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 348
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`
	// get icon
	s = `
GET /cgi-bin/mmwebwx-bin/webwxgeticon?seq=1924835613&username=@b1086381738436dbe3bae7a9a80a91c7f2d68a1c09053ef6e6bc2407fb0b73df&skey=@crypt_24c998a9_cd569c96d8f7d3242c84e41d1e31af5d HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: image/webp,image/*,*/*;q=0.8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`
	// get contact
	s = `
GET /cgi-bin/mmwebwx-bin/webwxgetcontact?r=1475626065944&seq=0&skey=@crypt_24c998a9_cd569c96d8f7d3242c84e41d1e31af5d HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Accept: application/json, text/plain, */*
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`
	// synccheck
	s = `
GET /cgi-bin/mmwebwx-bin/synccheck?r=1475626065956&skey=%40crypt_24c998a9_cd569c96d8f7d3242c84e41d1e31af5d&sid=PXmeHy4FQn4Rk3Wa&uin=1377554769&deviceid=e073867393922767&synckey=1_660260314%7C2_660260380%7C3_660260324%7C1000_1475573041&_=1475626064745 HTTP/1.1
Host: webpush.wx2.qq.com
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Accept: */*
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN
`
	s = `
POST /cgi-bin/mmwebwx-bin/webwxbatchgetcontact?type=ex&r=1475626066004 HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 264
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`

	s = `
POST /cgi-bin/mmwebwx-bin/webwxsync?sid=PXmeHy4FQn4Rk3Wa&skey=@crypt_24c998a9_cd569c96d8f7d3242c84e41d1e31af5d HTTP/1.1
Host: wx2.qq.com
Connection: keep-alive
Content-Length: 302
Accept: application/json, text/plain, */*
Origin: https://wx2.qq.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
Content-Type: application/json;charset=UTF-8
Referer: https://wx2.qq.com/
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; wxloadtime=1475626047; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1
`
	// recv msg
	s = `POST /cgi-bin/mmwebwx-bin/webwxstatusnotify HTTP/1.1
	Host: wx2.qq.com
	Connection: keep-alive
	Content-Length: 316
	Accept: application/json, text/plain, */*
	Origin: https://wx2.qq.com
	User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
	Content-Type: application/json;charset=UTF-8
	Referer: https://wx2.qq.com/
	Accept-Encoding: gzip, deflate, br
	Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
	Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxloadtime=1475626047_expired; wxpluginkey=1475625468; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM`

	// send msg
	s = `POST /cgi-bin/mmwebwx-bin/webwxsendmsg HTTP/1.1
	Host: wx2.qq.com
	Connection: keep-alive
	Content-Length: 392
	Accept: application/json, text/plain, */*
	Origin: https://wx2.qq.com
	User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36
	Content-Type: application/json;charset=UTF-8
	Referer: https://wx2.qq.com/
	Accept-Encoding: gzip, deflate, br
	Accept-Language: zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,zh-TW;q=0.2
	Cookie: webwxuvid=3da0e5f13fe04f5a96c54ca0dfccebcc716864e1946b1e94e80fe5828ef73c0e59698818e2d2b7d41a092d881031e9ba; pgv_pvi=8288752640; pgv_si=s4345098240; mm_lang=zh_CN; MM_WX_NOTIFY_STATE=1; MM_WX_SOUND_STATE=1; wxloadtime=1475626047_expired; wxpluginkey=1475626147; wxuin=1377554769; wxsid=PXmeHy4FQn4Rk3Wa; webwx_data_ticket=gSckAlyEqEDG0HqM504f7eNM`

	if len(url) != 0 && len(s) != 0 {
		t.Skip("")
	}
}
