package model

type Provider interface {
	Send(message Message) error
}

type Message struct {
	Provider      string            `json:"provider"`
	To            []string          `json:"phone_number"`
	From          string            `json:"sign_name"`
	TemplateCode  string            `json:"template_code"`
	TemplateParam map[string]string `json:"template_param"`
}
