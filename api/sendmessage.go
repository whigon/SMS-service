package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"sms/model"
	"sms/provider/aliyun"
)

func SendSMS(w http.ResponseWriter, r *http.Request) error {
	message := model.Message{}
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(message)

	// 获取provider
	provider, err := providerByName(message.Provider)
	if err != nil {
		return err
	}

	// 发送短信
	if err = provider.Send(message); err != nil {
		return err
	}

	return sendOK(w, http.StatusOK)
}

func providerByName(name string) (model.Provider, error) {
	//fmt.Println(name)
	switch name {
	case "aliyun":
		return aliyun.NewAliyun()
	}

	return nil, fmt.Errorf("%s: Unknown provider", name)
}

func sendOK(w http.ResponseWriter, data interface{}) error {
	err := sendJSON(w, http.StatusOK, map[string]interface{}{
		"data": data,
	})
	return err
}

func sendJSON(w http.ResponseWriter, status int, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error encoding json response: %v", obj))
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	return err
}
