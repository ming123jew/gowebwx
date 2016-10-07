package webwx

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestProtoEncodeBlank(t *testing.T) {
	r := SyncRequest{}
	b, err := json.Marshal(r)
	if err != nil {
		t.Error("json Marshal failed", err)
	}
	t.Log(string(b))
}

func TestProtoDecodePartial(t *testing.T) {
	// TODO
}

func _testProtocol(t *testing.T, filename string, r interface{}, debug bool) {
	f, err := os.Open(filename)
	if err != nil {
		t.Skip("open file failed", err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Skip("read file failed", err)
	}

	// Unmarshal
	rv := reflect.ValueOf(r)
	if debug {
		t.Log("type", rv.Kind())
	}
	if rv.Kind() == reflect.Ptr {
		err = json.Unmarshal(data, rv.Interface())
	} else {
		err = json.Unmarshal(data, &r)
	}
	if err != nil {
		t.Error("Unmarshall failed", err)
	}
	// Marshal agiain
	b, err := json.Marshal(r)
	if err != nil {
		t.Error("json Marshal failed", err)
	}
	if debug {
		//t.Log(string(data))
		t.Log(string(b))
	}
}

func TestProtoInitRequest(t *testing.T) {
	filename := "testdata/init.json"
	var r InitRequest
	_testProtocol(t, filename, &r, false)
	//t.Log("Uin", r.BaseRequest.Uin)
}

func TestProtoInitResponse(t *testing.T) {
	filename := "testdata/initResp.json"
	var r InitResponse
	_testProtocol(t, filename, &r, false)
	//t.Log("SyncKey", r.SyncKey.format())
}

func TestProtoInitResponse2(t *testing.T) {
	filename := "testdata/initResp2.json"
	var r InitResponse
	_testProtocol(t, filename, &r, true)
	t.Log("SyncKey", r.SyncKey.format())
}

func TestProtoStatusNotifyRequest(t *testing.T) {
}

func TestProtoStatusNotifyResponse(t *testing.T) {
}

func TestProtoGetContactResponse(t *testing.T) {
	filename := "testdata/getcontactResp.json"
	var r GetContactResponse
	_testProtocol(t, filename, r, false)
}

func TestProtoBatchGetContactRequest(t *testing.T) {
	filename := "testdata/batchgetcontact.json"
	var r BatchGetContactRequest
	_testProtocol(t, filename, r, false)
}

func TestProtoBatchGetContactResponse(t *testing.T) {
	filename := "testdata/batchgetcontactResp.json"
	var r BatchGetContactResponse
	_testProtocol(t, filename, r, false)
}

func TestProtoSyncRequest(t *testing.T) {
	filename := "testdata/sync.json"
	var r SyncRequest
	_testProtocol(t, filename, &r, false)
	//t.Log("SyncKey", r.SyncKey.format())
}

func TestProtoSyncResponse(t *testing.T) {
	filename := "testdata/syncResp.json"
	var r SyncResponse
	_testProtocol(t, filename, &r, false)
	//t.Log("SyncKey", r.SyncKey.format())
}

func TestProtoUpdateChatRoomAddRequest(t *testing.T) {
}
func TestProtoUpdateChatRoomDelRequest(t *testing.T) {
}
func TestProtoUpdateChatRoomResponse(t *testing.T) {
}
func TestProtoVerifyUserRequest(t *testing.T) {
}
func TestProtoSendMsgRequest(t *testing.T) {
}
