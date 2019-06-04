package models

import (
	"encoding/json"

	"github.com/lifei6671/mindoc/utils/requests"
)

const (
	BaseURL       = "https://open.fxiaoke.com"
	AppID         = "FSAID_13182c4"
	AppSecret     = "4e288fc4c091490f94817fe32d8eaddf"
	PermanentCode = "A61881643BBB424D9AE02D8764D4F0BF"
)

var (
	AppAccessToken  string
	CorpAccessToken string
	CorpID          string
)

type FxiaokeSdk struct {
}

func NewFxiaokeSdk() *FxiaokeSdk {
	result := new(FxiaokeSdk)
	appTokenResp, _ := result.CgiAppAccessTokenGet()
	AppAccessToken = appTokenResp.AppAccessToken

	cropTokenResp, _ := result.CgiCorpAccessTokenGet()
	CorpAccessToken = cropTokenResp.CorpAccessToken
	CorpID = cropTokenResp.CorpID
	return result
}

type CgiAppAccessTokenGetRequest struct {
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type CgiAppAccessTokenGetResponse struct {
	AppAccessToken string `json:"appAccessToken"`
	ExpiresIn      int    `json:"expiresIn"`
	ErrorCode      int    `json:"errorCode"`
	ErrorMessage   string `json:"errorMessage"`
}

type CgiCorpAccessTokenGetRequest struct {
	AppID         string `json:"appId"`
	AppSecret     string `json:"appSecret"`
	PermanentCode string `json:"permanentCode"`
}

type CgiCorpAccessTokenGetResponse struct {
	CorpAccessToken string `json:"corpAccessToken"`
	CorpID          string `json:"corpId"`
	ExpiresIn       int    `json:"expiresIn"`
	ErrorCode       int    `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
}

type Oauth2OpenUserIdGetRequest struct {
	AppAccessToken string `json:"appAccessToken"`
	Code           string `json:"code"`
}

type Oauth2OpenUserIdGetResponse struct {
	CorpID       string `json:"corpId"`
	OpenUserID   string `json:"openUserId"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type CgiUserGetRequest struct {
	CorpAccessToken string `json:"corpAccessToken"`
	CorpID          string `json:"corpId"`
	OpenUserID      string `json:"openUserId"`
}

type CgiUserGetResponse struct {
	OpenUserID      string `json:"openUserId"`
	Account         string `json:"account"`
	Name            string `json:"name"`
	NickName        string `json:"nickName"`
	Email           string `json:"email"`
	Mobile          string `json:"mobile"`
	Gender          string `json:"gender"`
	IsStop          string `json:"isStop"`
	ProfileImageURL string `json:"profileImageUrl"`
	DepartmentIds   []int  `json:"departmentIds"`
	ErrorCode       int    `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
}

type CgiMessageSendTextContent struct {
	Content string `json:"content"`
}

type CgiMessageSendTextRequest struct {
	CorpAccessToken string                    `json:"corpAccessToken"`
	CorpID          string                    `json:"corpId"`
	ToUser          string                    `json:"toUser"`
	MsgType         string                    `json:"msgType"`
	Text            CgiMessageSendTextContent `json:"text"`
}

type CgiMessageSendTextResponse struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type CgiJsapiticketGetRequest struct {
	CorpAccessToken string `json:"corpAccessToken"`
	CorpID          string `json:"corpId"`
}

type CgiJsapiticketGetResponse struct {
	Ticket       string `json:"ticket"`
	ExpiresIn    int    `json:"expiresIn"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func (f *FxiaokeSdk) CgiAppAccessTokenGet() (CgiAppAccessTokenGetResponse, error) {
	req, _ := json.Marshal(CgiAppAccessTokenGetRequest{AppID, AppSecret})
	result, err := requests.HttpPostWithJson(BaseURL+"/cgi/appAccessToken/get", string(req))
	response := CgiAppAccessTokenGetResponse{}
	json.Unmarshal(result, &response)
	return response, err
}

func (f *FxiaokeSdk) CgiCorpAccessTokenGet() (CgiCorpAccessTokenGetResponse, error) {
	req, _ := json.Marshal(CgiCorpAccessTokenGetRequest{AppID, AppSecret, PermanentCode})
	result, err := requests.HttpPostWithJson(BaseURL+"/cgi/corpAccessToken/get/V2", string(req))
	response := CgiCorpAccessTokenGetResponse{}
	json.Unmarshal(result, &response)
	return response, err
}

func (f *FxiaokeSdk) Oauth2OpenUserIdGet(code string) (Oauth2OpenUserIdGetResponse, error) {
	req, _ := json.Marshal(Oauth2OpenUserIdGetRequest{AppAccessToken, code})
	result, err := requests.HttpPostWithJson(BaseURL+"/oauth2/openUserId/get", string(req))
	response := Oauth2OpenUserIdGetResponse{}
	json.Unmarshal(result, &response)
	return response, err
}

func (f *FxiaokeSdk) CgiUserGet(openUserID string) (CgiUserGetResponse, error) {
	req, _ := json.Marshal(CgiUserGetRequest{CorpAccessToken, CorpID, openUserID})
	result, err := requests.HttpPostWithJson(BaseURL+"/cgi/user/get", string(req))
	response := CgiUserGetResponse{}
	json.Unmarshal(result, &response)
	return response, err
}

func (f *FxiaokeSdk) CgiMessageSendText(toUser, content string) (CgiMessageSendTextResponse, error) {
	req, _ := json.Marshal(CgiMessageSendTextRequest{CorpAccessToken, CorpID, toUser, "text", CgiMessageSendTextContent{content}})
	result, err := requests.HttpPostWithJson(BaseURL+"/cgi/message/send", string(req))
	response := CgiMessageSendTextResponse{}
	json.Unmarshal(result, &response)
	return response, err
}

func (f *FxiaokeSdk) CgiJsapiticketGet() (CgiJsapiticketGetResponse, error) {
	req, _ := json.Marshal(CgiJsapiticketGetRequest{CorpAccessToken, CorpID})
	result, err := requests.HttpPostWithJson(BaseURL+"/cgi/jsApiTicket/get", string(req))
	response := CgiJsapiticketGetResponse{}
	json.Unmarshal(result, &response)
	return response, err
}
