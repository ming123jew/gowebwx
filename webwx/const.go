// ref:  https://res.wx.qq.com/zh_CN/htmledition/v2/js/webwxApp308288.js

package webwx

import "strings"

const (
	LANG     = "zh_CN"
	RES_PATH = "/zh_CN/htmledition/v2/"
	APPID    = ""

	API_jsLogin            = "/jslogin?appid=wx782c26e4c19acffb&fun=new" // use w.login_host
	API_redirecturi        = "/cgi-bin/mmwebwx-bin/webwxnewloginpage"    // use w.host, will change host
	API_login              = "/cgi-bin/mmwebwx-bin/login"                // use w.login_host
	API_synccheck          = "/cgi-bin/mmwebwx-bin/synccheck"            // use w.webpush_host
	API_webwxdownloadmedia = "/cgi-bin/mmwebwx-bin/webwxgetmedia"        // use w.upload_host
	API_webwxuploadmedia   = "/cgi-bin/mmwebwx-bin/webwxuploadmedia"     // use w.upload_host

	API_webwxpreview         = "/cgi-bin/mmwebwx-bin/webwxpreview"
	API_webwxinit            = "/cgi-bin/mmwebwx-bin/webwxinit"
	API_webwxgetcontact      = "/cgi-bin/mmwebwx-bin/webwxgetcontact"
	API_webwxsync            = "/cgi-bin/mmwebwx-bin/webwxsync"
	API_webwxbatchgetcontact = "/cgi-bin/mmwebwx-bin/webwxbatchgetcontact"
	API_webwxgeticon         = "/cgi-bin/mmwebwx-bin/webwxgeticon"
	API_webwxsendmsg         = "/cgi-bin/mmwebwx-bin/webwxsendmsg"
	API_webwxsendmsgimg      = "/cgi-bin/mmwebwx-bin/webwxsendmsgimg"
	API_webwxsendmsgvedio    = "/cgi-bin/mmwebwx-bin/webwxsendvideomsg"
	API_webwxsendemoticon    = "/cgi-bin/mmwebwx-bin/webwxsendemoticon"
	API_webwxsendappmsg      = "/cgi-bin/mmwebwx-bin/webwxsendappmsg"
	API_webwxgetheadimg      = "/cgi-bin/mmwebwx-bin/webwxgetheadimg"
	API_webwxgetmsgimg       = "/cgi-bin/mmwebwx-bin/webwxgetmsgimg"
	API_webwxgetmedia        = "/cgi-bin/mmwebwx-bin/webwxgetmedia"
	API_webwxgetvideo        = "/cgi-bin/mmwebwx-bin/webwxgetvideo"
	API_webwxlogout          = "/cgi-bin/mmwebwx-bin/webwxlogout"
	API_webwxgetvoice        = "/cgi-bin/mmwebwx-bin/webwxgetvoice"
	API_webwxupdatechatroom  = "/cgi-bin/mmwebwx-bin/webwxupdatechatroom"
	API_webwxcreatechatroom  = "/cgi-bin/mmwebwx-bin/webwxcreatechatroom"
	API_webwxstatusnotify    = "/cgi-bin/mmwebwx-bin/webwxstatusnotify"
	API_webwxcheckurl        = "/cgi-bin/mmwebwx-bin/webwxcheckurl"
	API_webwxverifyuser      = "/cgi-bin/mmwebwx-bin/webwxverifyuser"
	API_webwxfeedback        = "/cgi-bin/mmwebwx-bin/webwxsendfeedback"
	API_webwxreport          = "/cgi-bin/mmwebwx-bin/webwxstatreport"
	API_webwxsearch          = "/cgi-bin/mmwebwx-bin/webwxsearchcontact"
	API_webwxoplog           = "/cgi-bin/mmwebwx-bin/webwxoplog"
	API_checkupload          = "/cgi-bin/mmwebwx-bin/webwxcheckupload"
)

type _HostMap struct {
	host          []string
	loginPrefix   []string
	uploadPrefix  []string
	webpushPrefix []string
}

func (m _HostMap) UpdateHost(w *webwx) {
	e := w.base_host
	w.login_host = m.loginPrefix[0]
	w.upload_host = m.uploadPrefix[0]
	w.webpush_host = m.webpushPrefix[0]
	for i := 1; i < 6; i++ {
		h := m.host[i]
		if strings.Contains(e, h) {
			w.login_host = m.loginPrefix[i] + h
			w.upload_host = m.uploadPrefix[i] + h
			w.webpush_host = m.webpushPrefix[i] + h
			break
		}
	}
}

var HostMap = _HostMap{
	[]string{"", "wx2.qq.com", "wx8.qq.com", "qq.com", "web2.wechat.com", "wechat.com"},
	[]string{"login.weixin.qq.com", "login.", "login.", "login.wx.", "login.", "login.web."},
	[]string{"file.wx.qq.com", "file.", "file.", "file.wx.", "file.", "file.web."},
	[]string{"webpush.weixin.qq.com", "webpush.", "webpush.", "webpush.wx.", "webpush.", "webpush.web."},
}

const (
	EMOTICON_REG = `img\\sclass="(qq)?emoji (qq)?emoji([\\da-f]*?)"\\s(text="[^<>(\\s]*")?\\s?src="[^<>(\\s]*"\\s*`

	SP_CONTACT_FILE_HELPER      = "filehelper"
	SP_CONTACT_NEWSAPP          = "newsapp"
	SP_CONTACT_RECOMMEND_HELPER = "fmessage"

	OpLogCmdId_TOPCONTACT    = 3
	OpLogCmdId_MODREMARKNAME = 2

	CONTACTFLAG_CONTACT             = 1
	CONTACTFLAG_CHATCONTACT         = 2
	CONTACTFLAG_CHATROOMCONTACT     = 4
	CONTACTFLAG_BLACKLISTCONTACT    = 8
	CONTACTFLAG_DOMAINCONTACT       = 16
	CONTACTFLAG_HIDECONTACT         = 32
	CONTACTFLAG_FAVOURCONTACT       = 64
	CONTACTFLAG_3RDAPPCONTACT       = 128
	CONTACTFLAG_SNSBLACKLISTCONTACT = 256
	CONTACTFLAG_NOTIFYCLOSECONTACT  = 512
	CONTACTFLAG_TOPCONTACT          = 2048

	MM_USERATTRVERIFYFALG_BIZ          = 1
	MM_USERATTRVERIFYFALG_FAMOUS       = 2
	MM_USERATTRVERIFYFALG_BIZ_BIG      = 4
	MM_USERATTRVERIFYFALG_BIZ_BRAND    = 8
	MM_USERATTRVERIFYFALG_BIZ_VERIFIED = 16

	MM_DATA_TEXT                = 1
	MM_DATA_HTML                = 2
	MM_DATA_IMG                 = 3
	MM_DATA_PRIVATEMSG_TEXT     = 11
	MM_DATA_PRIVATEMSG_HTML     = 12
	MM_DATA_PRIVATEMSG_IMG      = 13
	MM_DATA_VOICEMSG            = 34
	MM_DATA_PUSHMAIL            = 35
	MM_DATA_QMSG                = 36
	MM_DATA_VERIFYMSG           = 37
	MM_DATA_PUSHSYSTEMMSG       = 38
	MM_DATA_QQLIXIANMSG_IMG     = 39
	MM_DATA_POSSIBLEFRIEND_MSG  = 40
	MM_DATA_SHARECARD           = 42
	MM_DATA_VIDEO               = 43
	MM_DATA_VIDEO_IPHONE_EXPORT = 44
	MM_DATA_EMOJI               = 47
	MM_DATA_LOCATION            = 48
	MM_DATA_APPMSG              = 49
	MM_DATA_VOIPMSG             = 50
	MM_DATA_STATUSNOTIFY        = 51
	MM_DATA_VOIPNOTIFY          = 52
	MM_DATA_VOIPINVITE          = 53
	MM_DATA_MICROVIDEO          = 62
	MM_DATA_SYSNOTICE           = 9999
	MM_DATA_SYS                 = 1e4
	MM_DATA_RECALLED            = 10002

	MSGTYPE_TEXT               = 1
	MSGTYPE_IMAGE              = 3
	MSGTYPE_VOICE              = 34
	MSGTYPE_VIDEO              = 43
	MSGTYPE_MICROVIDEO         = 62
	MSGTYPE_EMOTICON           = 47
	MSGTYPE_APP                = 49
	MSGTYPE_VOIPMSG            = 50
	MSGTYPE_VOIPNOTIFY         = 52
	MSGTYPE_VOIPINVITE         = 53
	MSGTYPE_LOCATION           = 48
	MSGTYPE_STATUSNOTIFY       = 51
	MSGTYPE_SYSNOTICE          = 9999
	MSGTYPE_POSSIBLEFRIEND_MSG = 40
	MSGTYPE_VERIFYMSG          = 37
	MSGTYPE_SHARECARD          = 42
	MSGTYPE_SYS                = 1e4
	MSGTYPE_RECALLED           = 10002

	MSG_SEND_STATUS_READY   = 0
	MSG_SEND_STATUS_SENDING = 1
	MSG_SEND_STATUS_SUCC    = 2
	MSG_SEND_STATUS_FAIL    = 5

	APPMSGTYPE_TEXT                    = 1
	APPMSGTYPE_IMG                     = 2
	APPMSGTYPE_AUDIO                   = 3
	APPMSGTYPE_VIDEO                   = 4
	APPMSGTYPE_URL                     = 5
	APPMSGTYPE_ATTACH                  = 6
	APPMSGTYPE_OPEN                    = 7
	APPMSGTYPE_EMOJI                   = 8
	APPMSGTYPE_VOICE_REMIND            = 9
	APPMSGTYPE_SCAN_GOOD               = 10
	APPMSGTYPE_GOOD                    = 13
	APPMSGTYPE_EMOTION                 = 15
	APPMSGTYPE_CARD_TICKET             = 16
	APPMSGTYPE_REALTIME_SHARE_LOCATION = 17
	APPMSGTYPE_TRANSFERS               = 2e3
	APPMSGTYPE_RED_ENVELOPES           = 2001
	APPMSGTYPE_READER_TYPE             = 100001

	UPLOAD_MEDIA_TYPE_IMAGE      = 1
	UPLOAD_MEDIA_TYPE_VIDEO      = 2
	UPLOAD_MEDIA_TYPE_AUDIO      = 3
	UPLOAD_MEDIA_TYPE_ATTACHMENT = 4

	PROFILE_BITFLAG_NOCHANGE = 0
	PROFILE_BITFLAG_CHANGE   = 190

	CHATROOM_NOTIFY_OPEN  = 1
	CHATROOM_NOTIFY_CLOSE = 0

	StatusNotifyCode_READED        = 1
	StatusNotifyCode_ENTER_SESSION = 2
	StatusNotifyCode_INITED        = 3
	StatusNotifyCode_SYNC_CONV     = 4
	StatusNotifyCode_QUIT_SESSION  = 5

	VERIFYUSER_OPCODE_ADDCONTACT   = 1
	VERIFYUSER_OPCODE_SENDREQUEST  = 2
	VERIFYUSER_OPCODE_VERIFYOK     = 3
	VERIFYUSER_OPCODE_VERIFYREJECT = 4
	VERIFYUSER_OPCODE_SENDERREPLY  = 5
	VERIFYUSER_OPCODE_RECVERREPLY  = 6

	ADDSCENE_PF_QQ      = 4
	ADDSCENE_PF_EMAIL   = 5
	ADDSCENE_PF_CONTACT = 6
	ADDSCENE_PF_WEIXIN  = 7
	ADDSCENE_PF_GROUP   = 8
	ADDSCENE_PF_UNKNOWN = 9
	ADDSCENE_PF_MOBILE  = 10
	ADDSCENE_PF_WEB     = 33

	TIMEOUT_SYNC_CHECK = 0

	EMOJI_FLAG_GIF = 2

	KEYCODE_BACKSPACE   = 8
	KEYCODE_ENTER       = 13
	KEYCODE_SHIFT       = 16
	KEYCODE_ESC         = 27
	KEYCODE_DELETE      = 34
	KEYCODE_ARROW_LEFT  = 37
	KEYCODE_ARROW_UP    = 38
	KEYCODE_ARROW_RIGHT = 39
	KEYCODE_ARROW_DOWN  = 40
	KEYCODE_NUM2        = 50
	KEYCODE_AT          = 64
	KEYCODE_NUM_ADD     = 107
	KEYCODE_NUM_MINUS   = 109
	KEYCODE_ADD         = 187
	KEYCODE_MINUS       = 189

	MM_NOTIFY_CLOSE = 0
	MM_NOTIFY_OPEN  = 1
	MM_SOUND_CLOSE  = 0
	MM_SOUND_OPEN   = 1

	MM_SEND_FILE_STATUS_QUEUED  = 0
	MM_SEND_FILE_STATUS_SENDING = 1
	MM_SEND_FILE_STATUS_SUCCESS = 2
	MM_SEND_FILE_STATUS_FAIL    = 3
	MM_SEND_FILE_STATUS_CANCEL  = 4

	MM_EMOTICON_WEB = "_web"

	RES_IMG_DEFAULT       = RES_PATH + "images/img.gif"
	RES_IMG_PLACEHOLDER   = RES_PATH + "images/spacer.gif"
	RES_SOUND_RECEIVE_MSG = RES_PATH + "sound/msg.mp3"
	RES_SOUND_SEND_MSG    = RES_PATH + "sound/text.mp3"
)
