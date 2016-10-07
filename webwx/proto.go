package webwx

import (
	"net/url"
	"strconv"
	"strings"
)

type Contact struct {
	Alias             string
	AppAccountFlag    int
	AttrStatus        int
	ChatRoomId        int
	City              string
	ContactFlag       int
	DisplayName       string
	EncryChatRoomId   string
	HeadImgFlag       int
	HeadImgUrl        string
	HideInputBarFlag  int
	KeyWord           string
	MemberCount       int
	MemberList        []ChatRoomMember
	NickName          string
	OwnerUin          int
	PYInitial         string
	PYQuanPin         string
	Province          string
	RemarkName        string
	RemarkPYInitial   string
	RemarkPYQuanPin   string
	Sex               int
	Signature         string
	SnsFlag           int
	StarFriend        int
	Statues           int
	Uin               int
	UniFriend         int
	UserName          string
	VerifyFlag        int
	WebWxPluginSwitch int
}

type ChatRoomMember struct {
	Contact
}

type User struct {
	Contact
}

type _Buff struct {
	Buff string
}

type Profile struct {
	Alias             string
	BindEmail         _Buff
	BindMobile        _Buff
	BindUin           int
	BitFlag           int
	HeadImgUpdateFlag int
	HeadImgUrl        string
	NickName          _Buff
	PersonalCard      int
	Sex               int
	Signature         string
	Status            int
	UserName          _Buff
}

type AppInfo struct {
	AppID string
	Type  int
}

type RecommendInfo struct {
	Alias      string
	AttrStatus int
	City       string
	Content    string
	NickName   string
	OpCode     int
	Province   string
	QQNum      int
	Scene      int
	Sex        int
	Signature  string
	Ticket     string
	UserName   string
	VerifyFlag int
}

type Msg struct {
	AppInfo              AppInfo
	AppMsgType           int
	Content              string
	CreateTime           int
	FileName             string
	FileSize             string
	ForwardFlag          int
	FromUserName         string
	HasProductId         int
	ImgHeight            int
	ImgStatus            int
	ImgWidth             int
	MediaId              string
	MsgId                string
	MsgType              int
	NewMsgId             int
	PlayLength           int
	RecommendInfo        RecommendInfo
	Status               int
	StatusNotifyCode     int
	StatusNotifyUserName string
	SubMsgType           int
	Ticket               string
	ToUserName           string
	Url                  string
	VoiceLength          int
}

type MPArticle struct {
	Title  string
	Digest string
	Cover  string
	Url    string
}

type MPSubscribeMsg struct {
	UserName       string
	MPArticleCount int
	MPArticleList  []MPArticle
	Time           int
	NickName       string
}

type SyncKeyItem struct {
	Key int
	Val int
}

type SyncKey struct {
	Count int
	List  []SyncKeyItem
}

func (k SyncKey) format() string {
	tmp := []string{}
	for _, item := range k.List {
		tmp = append(tmp, strconv.Itoa(item.Key)+"_"+strconv.Itoa(item.Val))
	}
	return url.QueryEscape(strings.Join(tmp, "|"))
}

type BaseRequest struct {
	Uin      string
	Sid      string
	Skey     string
	DeviceID string
}

type BaseResponse struct {
	Ret    int
	ErrMsg string
}

type InitRequest struct {
	BaseRequest BaseRequest
}

type InitResponse struct {
	BaseResponse        BaseResponse
	ChatSet             string
	ClickReportInterval int
	ClientVersion       int
	ContactList         []Contact
	Count               int
	GrayScale           int
	InviteStartCount    int
	MPSubscribeMsgCount int
	MPSubscribeMsgList  []MPSubscribeMsg
	SKey                string
	SyncKey             SyncKey
	SystemTime          int
	User                User
}

type StatusNotifyRequest struct {
	BaseRequest  BaseRequest
	ClientMsgId  int
	Code         int
	FromUserName string
	ToUserName   string
}

type StatusNotifyResponse struct {
	BaseResponse BaseResponse
	MsgID        string
}

type GetContactResponse struct {
	BaseResponse BaseResponse
	MemberCount  int
	MemberList   []Contact
	Seq          int
}

type _BatchContactItem struct {
	ChatRoomId string
	UserName   string
}

type BatchGetContactRequest struct {
	BaseRequest BaseRequest
	Count       int
	List        []_BatchContactItem
}

type BatchGetContactResponse struct {
	BaseResponse BaseResponse
	ContactList  []Contact
	Count        int
}

type SyncRequest struct {
	BaseRequest BaseRequest
	SyncKey     SyncKey
	rr          int
}

type SyncResponse struct {
	BaseResponse           BaseResponse
	AddMsgCount            int
	AddMsgList             []Msg
	ContinueFlag           int
	DelContactCount        int
	DelContactList         []Contact
	ModChatRoomMemberCount int
	ModChatRoomMemberList  []Contact
	ModContactCount        int
	ModContactList         []Contact
	Profile                Profile
	SKey                   string
	SyncCheckKey           SyncKey
	SyncKey                SyncKey
}

type UpdateChatRoomAddRequest struct {
	BaseRequest   BaseRequest
	AddMemberList string
	ChatRoomName  string
}

type UpdateChatRoomDelRequest struct {
	BaseRequest   BaseRequest
	DelMemberList string
	ChatRoomName  string
}

type UpdateChatRoomResponse struct {
	BaseResponse BaseResponse
	MemberCount  int
	MemberList   []ChatRoomMember
}

type _VerifyUser struct {
	Value            string
	VerifyUserTicket string
}

type VerifyUserRequest struct {
	BaseRequest        BaseRequest
	Opcode             int
	SceneList          []int
	SceneListCount     int
	VerifyContent      int
	VerifyUserList     []_VerifyUser
	VerifyUserListSize int
	skey               string
}

type SendMsgRequest struct {
}
