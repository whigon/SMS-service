package aliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"sms/model"
)

type AliyunConfig struct {
	RegionId  string `yaml:"region_id"`
	APIKey    string `yaml:"api_key"`
	APISecret string `yaml:"api_secret"`
}

type Aliyun struct {
	client *dysmsapi.Client
}

func NewAliyun() (*Aliyun, error) {
	// 获取客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "", "")

	if err != nil {
		return nil, err
	}

	return &Aliyun{client: client}, nil
}

func (aliyun *Aliyun) Send(message model.Message) error {
	for _, phoneNumber := range message.To {
		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"

		request.SignName = message.From
		request.PhoneNumbers = phoneNumber
		request.TemplateCode = message.TemplateCode
		request.TemplateParam = "{"
		for k, v := range message.TemplateParam {
			request.TemplateParam = request.TemplateParam + k + ":" + v + ","
		}
		request.TemplateParam += "}"

		log.Println(request)

		response, err := aliyun.client.SendSms(request)

		if err != nil {
			log.Println(err)
			return err
		}
		if response.Message != "OK" {
			log.Println(response)
			return fmt.Errorf(response.Message)
		}
	}

	return nil
}
