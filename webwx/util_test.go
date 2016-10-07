package webwx

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestUtilGenDeviceId(t *testing.T) {
	if s := genDeviceId(); len(s) != 16 {
		t.Error("Invalid deviceId:", s, len(s))
	}
}

func TestUtilGetUserAgent(t *testing.T) {
	t.Log(getUserAgent())
}

func TestRegexp(t *testing.T) {
	s := `<error><ret>0</ret><message></message><skey>@crypt_24c998a9_f90604f28a133860a55b9954db925a2a</skey><wxsid>Hp6P6xBuGnZJS+j8</wxsid><wxuin>1377554769</wxuin><pass_ticket>COk9zr04ui3%2BiirzqXU8XP5x5OJf7p2Cb8H2qt2%2BtpwARpT2doI6WJqncBf3HjkO</pass_ticket><isgrayscale>1</isgrayscale></error>`
	p := regexp.MustCompile(`<error><ret>(\d+)</ret><message></message><skey>([\S]+)</skey><wxsid>([\S]+)</wxsid><wxuin>(\d+)</wxuin><pass_ticket>([\S]+)</pass_ticket><isgrayscale>(\d+)</isgrayscale></error>`)
	match := p.FindStringSubmatch(s)
	t.Log("match len:", len(match))
	for _, v := range match {
		t.Log(v)
	}
	s = `window.synccheck={retcode:"0",selector:"2"}`
	p = regexp.MustCompile(`\{retcode:"(\d+)"\s*,\s*selector:"(\d+)"}`)
	match = p.FindStringSubmatch(s)
	t.Log("match len:", len(match))
	for _, v := range match {
		t.Log(v)
	}

}

func TestEmptyStruct(t *testing.T) {
	r := BaseRequest{}
	s := fmt.Sprintf("%s, %s, %s, %s", r.Uin, r.Sid, r.Skey, r.DeviceID)
	t.Log(s)
}

func TestJsLoginResp(t *testing.T) {
	s := []string{
		`window.QRLogin.code = 200; window.QRLogin.uuid = "Qbdd7BUqXQ==";`,
		`window.code=408;`,
	}
	t.Skip(s[1])
}

func TestLoggerReflect(t *testing.T) {
	if log == nil {
		t.Skip("[WebWx] logger not inited, ignore..")
		return
	}
	// get access
	v := reflect.ValueOf(log)
	for {
		t.Log(v.Type(), v.Kind(), v.CanSet(), v.CanInterface(), v.CanAddr())
		//t := v.Type()
		//if t.String() == "seelog.logLevelConstraints" {
		//	break
		//}
		k := v.Kind()
		if k == reflect.Ptr || k == reflect.Interface {
			v = v.Elem()
		} else if k == reflect.Struct {
			v = v.Field(0)
		} else {
			break
		}
	}
	t.Log(v.Type(), v.Kind(), v.CanSet(), v.CanInterface(), v.CanAddr())
}
