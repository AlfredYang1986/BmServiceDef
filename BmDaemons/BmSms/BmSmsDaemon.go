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
	ApiName      string
	SignName     string
}

func (s BmSms) NewSmsDaemon(args map[string]string) *BmSms {
	return &BmSms{
		RegionId:     args["regionId"],
		AccessKeyId:  args["accessKeyId"],
		AccessSecret: args["accessSecret"],
		Domain:       args["domain"],
		Version:      args["version"],
		ApiName:      args["apiName"],
		SignName:     args["signName"],
	}
}

func (s BmSms) GetSmsClient() *sdk.Client {
	client, err := sdk.NewClientWithAccessKey(s.RegionId, s.AccessKeyId, s.AccessSecret)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func (s BmSms) SendMsg(phoneNumber string) (error, *responses.CommonResponse) {
	client := s.GetSmsClient()
	req := requests.NewCommonRequest()
	req.Method = "POST"
	req.Scheme = "https"
	req.Domain = s.Domain
	req.Version = s.Version
	req.ApiName = s.ApiName
	req.ApiName = s.ApiName
	req.QueryParams["RegionId"] = s.RegionId
	req.QueryParams["PhoneNumbers"] = phoneNumber
	req.QueryParams["SignName"] = s.SignName

	res, err := client.ProcessCommonRequest(req)
	if err != nil {
		panic(err.Error())
	}
	return err, res
}
