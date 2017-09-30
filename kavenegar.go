package kavenegar

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type BulkRecipientType string

const (
	BulkRecipientTypePermanentCells BulkRecipientType = "PermanentCells"

	BulkRecipientTypeTemporaryCells BulkRecipientType = "TemporaryCells"

	BulkRecipientTypeAllCells BulkRecipientType = "AllCells"
)

type SendAdvance struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendAdvance"`

	Apikey string `xml:"apikey,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Encoding int32 `xml:"encoding,omitempty"`

	Receptor string `xml:"receptor,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Messagemode int32 `xml:"messagemode,omitempty"`

	Fallbackurl string `xml:"fallbackurl,omitempty"`

	Clientmessageid string `xml:"clientmessageid,omitempty"`
}

type SendAdvanceResponse struct {
	SendAdvanceResult int64 `xml:"SendAdvanceResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type ProvinceInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ProvinceInfo"`
}

type ProvinceInfoResponse struct {
	ProvinceInfoResult *ArrayOfApiProvincies `xml:"ProvinceInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type CityInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ CityInfo"`

	Provinceid int32 `xml:"provinceid,omitempty"`
}

type CityInfoResponse struct {
	CityInfoResult *ArrayOfApiCities `xml:"CityInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendRandomBulkByApiKey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendRandomBulkByApiKey"`

	Apikey string `xml:"apikey,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Provinceid int32 `xml:"provinceid,omitempty"`

	Cityid int32 `xml:"cityid,omitempty"`

	RecipientKind *BulkRecipientType `xml:"recipientKind,omitempty"`

	Randomcount int64 `xml:"randomcount,omitempty"`
}

type SendRandomBulkByApiKeyResponse struct {
	SendRandomBulkByApiKeyResult *ArrayOfLong `xml:"SendRandomBulkByApiKeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendSequentialBulkByApiKey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendSequentialBulkByApiKey"`

	Apikey string `xml:"apikey,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Provinceid int32 `xml:"provinceid,omitempty"`

	Cityid int32 `xml:"cityid,omitempty"`

	RecipientKind *BulkRecipientType `xml:"recipientKind,omitempty"`

	Startindex int64 `xml:"startindex,omitempty"`

	Length int64 `xml:"length,omitempty"`
}

type SendSequentialBulkByApiKeyResponse struct {
	SendSequentialBulkByApiKeyResult *ArrayOfLong `xml:"SendSequentialBulkByApiKeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendSimpleByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendSimpleByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Receptor *ArrayOfString `xml:"receptor,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Msgmode int32 `xml:"msgmode,omitempty"`
}

type SendSimpleByApikeyResponse struct {
	SendSimpleByApikeyResult *ArrayOfLong `xml:"SendSimpleByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendSimpleByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendSimpleByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Receptor *ArrayOfString `xml:"receptor,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Msgmode int32 `xml:"msgmode,omitempty"`
}

type SendSimpleByLoginInfoResponse struct {
	SendSimpleByLoginInfoResult *ArrayOfLong `xml:"SendSimpleByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendArrayByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendArrayByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`

	Sender *ArrayOfString `xml:"sender,omitempty"`

	Message *ArrayOfString `xml:"message,omitempty"`

	Receptor *ArrayOfString `xml:"receptor,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Msgmode *ArrayOfInt `xml:"msgmode,omitempty"`
}

type SendArrayByLoginInfoResponse struct {
	SendArrayByLoginInfoResult *ArrayOfLong `xml:"SendArrayByLoginInfoResult,omitempty"`
}

type SendArrayByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendArrayByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Sender *ArrayOfString `xml:"sender,omitempty"`

	Message *ArrayOfString `xml:"message,omitempty"`

	Receptor *ArrayOfString `xml:"receptor,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Msgmode *ArrayOfInt `xml:"msgmode,omitempty"`
}

type SendArrayByApikeyResponse struct {
	SendArrayByApikeyResult *ArrayOfLong `xml:"SendArrayByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SendPostalCodeByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SendPostalCodeByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Message string `xml:"message,omitempty"`

	Unixdate int64 `xml:"unixdate,omitempty"`

	Msgmode int32 `xml:"msgmode,omitempty"`

	Postalcode int64 `xml:"postalcode,omitempty"`

	Startindex int64 `xml:"startindex,omitempty"`

	Count int64 `xml:"count,omitempty"`

	Random bool `xml:"random,omitempty"`
}

type SendPostalCodeByApikeyResponse struct {
	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type CountPostalCode struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ CountPostalCode"`

	Apikey string `xml:"apikey,omitempty"`

	Postalcode int64 `xml:"postalcode,omitempty"`
}

type CountPostalCodeResponse struct {
	CountPostalCodeResult int64 `xml:"CountPostalCodeResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type GetStatusByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ GetStatusByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type GetStatusByApikeyResponse struct {
	GetStatusByApikeyResult *ArrayOfInt `xml:"GetStatusByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type GetStatusByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ GetStatusByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type GetStatusByLoginInfoResponse struct {
	GetStatusByLoginInfoResult *ArrayOfInt `xml:"GetStatusByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type SelectByApikeyResponse struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectByApikeyResponse"`

	SelectByApikeyResult *ArrayOfApiSelect `xml:"SelectByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectByLoginInfo"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type SelectByLoginInfoResponse struct {
	SelectByLoginInfoResult *ArrayOfApiSelect `xml:"SelectByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectoutboxByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectoutboxByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	StartUnixdate int64 `xml:"startUnixdate,omitempty"`

	EndUnixdate int64 `xml:"endUnixdate,omitempty"`
}

type SelectoutboxByApikeyResponse struct {
	SelectoutboxByApikeyResult *ArrayOfApiSelect `xml:"SelectoutboxByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectoutboxByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectoutboxByLoginInfo"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	StartUnixdate int64 `xml:"startUnixdate,omitempty"`

	EndUnixdate int64 `xml:"endUnixdate,omitempty"`
}

type SelectoutboxByLoginInfoResponse struct {
	SelectoutboxByLoginInfoResult *ArrayOfApiSelect `xml:"SelectoutboxByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectlatestByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectlatestByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Pagesize int32 `xml:"pagesize,omitempty"`
}

type SelectlatestByApikeyResponse struct {
	SelectlatestByApikeyResult *ArrayOfApiSelect `xml:"SelectlatestByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type SelectlatestByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ SelectlatestByLoginInfo"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	Pagesize int32 `xml:"pagesize,omitempty"`
}

type SelectlatestByLoginInfoResponse struct {
	SelectlatestByLoginInfoResult *ArrayOfApiSelect `xml:"SelectlatestByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type CancelByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ CancelByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type CancelByApikeyResponse struct {
	CancelByApikeyResult *ArrayOfInt `xml:"CancelByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type CancelByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ CancelByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`

	Messageid *ArrayOfLong `xml:"messageid,omitempty"`
}

type CancelByLoginInfoResponse struct {
	CancelByLoginInfoResult *ArrayOfInt `xml:"CancelByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type ReceiveByApikey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ReceiveByApikey"`

	Apikey string `xml:"apikey,omitempty"`

	LineNumber string `xml:"lineNumber,omitempty"`

	Isread int16 `xml:"isread,omitempty"`
}

type ReceiveByApikeyResponse struct {
	ReceiveByApikeyResult *ArrayOfApiReceive `xml:"ReceiveByApikeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type ReceiveByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ReceiveByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`

	LineNumber string `xml:"lineNumber,omitempty"`

	Isread int16 `xml:"isread,omitempty"`
}

type ReceiveByLoginInfoResponse struct {
	ReceiveByLoginInfoResult *ArrayOfApiReceive `xml:"ReceiveByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type RemainCreditByApiKey struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ RemainCreditByApiKey"`

	Apikey string `xml:"apikey,omitempty"`
}

type RemainCreditByApiKeyResponse struct {
	RemainCreditByApiKeyResult int64 `xml:"RemainCreditByApiKeyResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type RemainCreditByLoginInfo struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ RemainCreditByLoginInfo"`

	UserName string `xml:"userName,omitempty"`

	Password string `xml:"password,omitempty"`
}

type RemainCreditByLoginInfoResponse struct {
	RemainCreditByLoginInfoResult int64 `xml:"RemainCreditByLoginInfoResult,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

type ArrayOfApiProvincies struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ArrayOfApiProvincies"`

	ApiProvincies []*ApiProvincies `xml:"ApiProvincies,omitempty"`
}

type ApiProvincies struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ApiProvincies"`

	Provinceid int32 `xml:"provinceid,omitempty"`

	Name string `xml:"name,omitempty"`

	Permanentcellcount int64 `xml:"permanentcellcount,omitempty"`

	Temporarycellcount int64 `xml:"temporarycellcount,omitempty"`
}

type ArrayOfApiCities struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ArrayOfApiCities"`

	ApiCities []*ApiCities `xml:"ApiCities,omitempty"`
}

type ApiCities struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ApiCities"`

	Cityid int32 `xml:"cityid,omitempty"`

	Name string `xml:"name,omitempty"`

	Permanentcellcount int64 `xml:"permanentcellcount,omitempty"`

	Temporarycellcount int64 `xml:"temporarycellcount,omitempty"`
}

type ArrayOfLong struct {
	Long []int64 `xml:"long,omitempty"`
}

type ArrayOfString struct {
	String []string `xml:"string,omitempty"`
}

type ArrayOfInt struct {
	Int []int32 `xml:"int,omitempty"`
}

type ArrayOfApiSelect struct {
	ApiSelect []*ApiSelect `xml:"ApiSelect,omitempty"`
}

type ApiSelect struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ApiSelect"`

	Messageid int64 `xml:"messageid,omitempty"`

	Message string `xml:"message,omitempty"`

	Status int32 `xml:"status,omitempty"`

	Statustext string `xml:"statustext,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Receptor string `xml:"receptor,omitempty"`

	Date int64 `xml:"date,omitempty"`

	Cost int32 `xml:"cost,omitempty"`
}

type ArrayOfApiReceive struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ArrayOfApiReceive"`

	ApiReceive []*ApiReceive `xml:"ApiReceive,omitempty"`
}

type ApiReceive struct {
	XMLName xml.Name `xml:"http://api.kavenegar.com/ ApiReceive"`

	Messageid int64 `xml:"messageid,omitempty"`

	Message string `xml:"message,omitempty"`

	Sender string `xml:"sender,omitempty"`

	Receptor string `xml:"receptor,omitempty"`

	Date int64 `xml:"date,omitempty"`
}

type v1Soap struct {
	client *SOAPClient
}

func Newv1Soap(url string, tls bool, auth *BasicAuth) *v1Soap {
	if url == "" {
		url = "http://api.kavenegar.com/soap/v1.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &v1Soap{
		client: client,
	}
}

func (service *v1Soap) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *v1Soap) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

func (service *v1Soap) SendAdvance(request *SendAdvance) (*SendAdvanceResponse, error) {
	response := new(SendAdvanceResponse)
	err := service.client.Call("http://api.kavenegar.com/SendAdvance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) ProvinceInfo(request *ProvinceInfo) (*ProvinceInfoResponse, error) {
	response := new(ProvinceInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/ProvinceInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) CityInfo(request *CityInfo) (*CityInfoResponse, error) {
	response := new(CityInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/CityInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendRandomBulkByApiKey(request *SendRandomBulkByApiKey) (*SendRandomBulkByApiKeyResponse, error) {
	response := new(SendRandomBulkByApiKeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SendRandomBulkByApiKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendSequentialBulkByApiKey(request *SendSequentialBulkByApiKey) (*SendSequentialBulkByApiKeyResponse, error) {
	response := new(SendSequentialBulkByApiKeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SendSequentialBulkByApiKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendSimpleByApikey(request *SendSimpleByApikey) (*SendSimpleByApikeyResponse, error) {
	response := new(SendSimpleByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SendSimpleByApikey", request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *v1Soap) SendSimpleByLoginInfo(request *SendSimpleByLoginInfo) (*SendSimpleByLoginInfoResponse, error) {
	response := new(SendSimpleByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/SendSimpleByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendArrayByLoginInfo(request *SendArrayByLoginInfo) (*SendArrayByLoginInfoResponse, error) {
	response := new(SendArrayByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/SendArrayByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendArrayByApikey(request *SendArrayByApikey) (*SendArrayByApikeyResponse, error) {
	response := new(SendArrayByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SendArrayByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SendPostalCodeByApikey(request *SendPostalCodeByApikey) (*SendPostalCodeByApikeyResponse, error) {
	response := new(SendPostalCodeByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SendPostalCodeByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) CountPostalCode(request *CountPostalCode) (*CountPostalCodeResponse, error) {
	response := new(CountPostalCodeResponse)
	err := service.client.Call("http://api.kavenegar.com/CountPostalCode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) GetStatusByApikey(request *GetStatusByApikey) (*GetStatusByApikeyResponse, error) {
	response := new(GetStatusByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/GetStatusByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) GetStatusByLoginInfo(request *GetStatusByLoginInfo) (*GetStatusByLoginInfoResponse, error) {
	response := new(GetStatusByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/GetStatusByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectByApikey(request *SelectByApikey) (*SelectByApikeyResponse, error) {
	response := new(SelectByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectByLoginInfo(request *SelectByLoginInfo) (*SelectByLoginInfoResponse, error) {
	response := new(SelectByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectoutboxByApikey(request *SelectoutboxByApikey) (*SelectoutboxByApikeyResponse, error) {
	response := new(SelectoutboxByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectoutboxByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectoutboxByLoginInfo(request *SelectoutboxByLoginInfo) (*SelectoutboxByLoginInfoResponse, error) {
	response := new(SelectoutboxByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectoutboxByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectlatestByApikey(request *SelectlatestByApikey) (*SelectlatestByApikeyResponse, error) {
	response := new(SelectlatestByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectlatestByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) SelectlatestByLoginInfo(request *SelectlatestByLoginInfo) (*SelectlatestByLoginInfoResponse, error) {
	response := new(SelectlatestByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/SelectlatestByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) CancelByApikey(request *CancelByApikey) (*CancelByApikeyResponse, error) {
	response := new(CancelByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/CancelByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) CancelByLoginInfo(request *CancelByLoginInfo) (*CancelByLoginInfoResponse, error) {
	response := new(CancelByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/CancelByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) ReceiveByApikey(request *ReceiveByApikey) (*ReceiveByApikeyResponse, error) {
	response := new(ReceiveByApikeyResponse)
	err := service.client.Call("http://api.kavenegar.com/ReceiveByApikey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) ReceiveByLoginInfo(request *ReceiveByLoginInfo) (*ReceiveByLoginInfoResponse, error) {
	response := new(ReceiveByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/ReceiveByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) RemainCreditByApiKey(request *RemainCreditByApiKey) (*RemainCreditByApiKeyResponse, error) {
	response := new(RemainCreditByApiKeyResponse)
	err := service.client.Call("http://api.kavenegar.com/RemainCreditByApiKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *v1Soap) RemainCreditByLoginInfo(request *RemainCreditByLoginInfo) (*RemainCreditByLoginInfoResponse, error) {
	response := new(RemainCreditByLoginInfoResponse)
	err := service.client.Call("http://api.kavenegar.com/RemainCreditByLoginInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Items []interface{} `xml:",omitempty"`
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tls     bool
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	//log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	//log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	respError := new(SOAPError)
	respEnvelope = new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: respError}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}
	if respError.Status != 200 {
		return fmt.Errorf("%d : %s", respError.Status, respError.Statusmessage)
	}

	return nil
}

//SOAPError ...
type SOAPError struct {
	Status int32 `xml:"status,omitempty"`

	Statusmessage string `xml:"statusmessage,omitempty"`
}

//New...
func New() *v1Soap {
	return Newv1Soap("", false, nil)
}
