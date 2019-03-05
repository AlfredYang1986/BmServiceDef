package BmSms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BmSms struct {
	RegionId     string
	AccessKeyId  string
	AccessSecret string
	Domain       string
	Version      string
	SignName     string
	TemplateCode string
}

func (s BmSms) NewSmsDaemon(args map[string]string) *BmSms {
	return &BmSms{
		RegionId:     args["regionId"],
		AccessKeyId:  args["accessKeyId"],
		AccessSecret: args["accessSecret"],
		Domain:       args["domain"],
		Version:      args["version"],
		SignName:     args["signName"],
		TemplateCode: args["templateCode"],
	}
}

func (s BmSms) GetSmsClient() *sdk.Client {
	client, err := sdk.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessSecret)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func (s BmSms) SendMsg(phoneNumber string, rcode string) (error, *responses.CommonResponse) {
	client := s.GetSmsClient()
	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = s.Domain
	req.Version = s.Version
	req.ApiName = "SendSms"
	req.QueryParams["RegionId"] = s.RegionId
	req.QueryParams["PhoneNumbers"] = phoneNumber
	req.QueryParams["SignName"] = s.SignName
	req.QueryParams["TemplateCode"] = s.TemplateCode
	req.QueryParams["TemplateParam"] = "{\"code\":\"" + rcode + "\"}"
	res, err := client.ProcessCommonRequest(req)

	if err != nil {
		panic(err.Error())
	}
	return err, res
}

func (s BmSms) VerifyCode(bizId string, phoneNumber string) (error, *responses.CommonResponse) {
	client := s.GetSmsClient()
	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = s.Domain
	req.Version = s.Version
	req.ApiName = "QuerySendDetails"
	req.QueryParams["RegionId"] = s.RegionId
	req.QueryParams["PhoneNumber"] = phoneNumber
	req.QueryParams["BizId"] = bizId

	res, err := client.ProcessCommonRequest(req)
	if err != nil {
		panic(err.Error())
	}
	return err, res
}
