package mandrilltest

import (
	"errors"
	"fmt"

	"github.com/Kasita-Inc/mandrill"
)

// TestClient implements the EmailClient interface for testing
type TestClient struct {
	APIKey       string
	Templates    map[string]*mandrill.Template
	Subaccounts  map[string]*mandrill.Subaccount
	Message      *mandrill.Message
	OK           bool
	TemplateName string
	Contents     interface{}
}

// NewClient returns a TestClient
func NewClient(apiKey string) (client *TestClient) {
	return &TestClient{
		OK:          true,
		APIKey:      apiKey,
		Templates:   make(map[string]*mandrill.Template, 0),
		Subaccounts: make(map[string]*mandrill.Subaccount, 0),
	}
}

// Ping mocks a ping to the Email Service
func (client *TestClient) Ping() (pong string, err error) {
	if client.OK {
		return "pong", nil
	}
	return "", errors.New("Ping failed")
}

// MessagesSend mocks sending an email through the Email Service
func (client *TestClient) MessagesSend(message *mandrill.Message) (responses []*mandrill.MessagesResponse, err error) {
	if !client.OK {
		return nil, errors.New("MessageSend failed")
	}
	client.Message = message
	for _, to := range message.To {
		response := &mandrill.MessagesResponse{Email: to.Email}
		responses = append(responses, response)
	}
	return responses, nil
}

// MessagesSendTemplate mocks sending a templated email through the Email Service
func (client *TestClient) MessagesSendTemplate(message *mandrill.Message, templateName string, contents interface{}) (
	responses []*mandrill.MessagesResponse, err error) {
	if !client.OK {
		return nil, errors.New("MessagesSendTemplate failed")
	}
	if _, ok := client.Templates[templateName]; !ok {
		return nil, fmt.Errorf("no template (%s) found", templateName)
	}
	client.TemplateName = templateName
	client.Contents = contents
	return client.MessagesSend(message)
}

// SubaccountInfo mocks getting SubaccountInfo from the Email Service
func (client *TestClient) SubaccountInfo(subaccountID string) (response *mandrill.Subaccount, err error) {
	if !client.OK {
		return nil, errors.New("SubaccountInfo failed")
	}
	subaccount, ok := client.Subaccounts[subaccountID]
	if !ok {
		return nil, fmt.Errorf("no subaccount (%s) found", subaccountID)
	}
	return subaccount, nil
}

// DeleteSubaccount mocks deleting a subaccount from the Email Service
func (client *TestClient) DeleteSubaccount(subaccountID string) (response *mandrill.Subaccount, err error) {
	if !client.OK {
		return nil, errors.New("DeleteSubaccount failed")
	}
	subaccount, ok := client.Subaccounts[subaccountID]
	if !ok {
		return nil, fmt.Errorf("no subaccount (%s) found", subaccountID)
	}
	delete(client.Subaccounts, subaccountID)
	return subaccount, nil
}

// UpdateSubaccount mocks updating a subaccount on the Email Service
func (client *TestClient) UpdateSubaccount(subaccount *mandrill.Subaccount) (response *mandrill.Subaccount, err error) {
	if !client.OK {
		return nil, errors.New("UpdateSubaccount failed")
	}
	if _, ok := client.Subaccounts[subaccount.ID]; !ok {
		return nil, fmt.Errorf("no subaccount (%s) found", subaccount.ID)
	}
	client.Subaccounts[subaccount.ID] = subaccount
	return subaccount, nil
}

// AddSubaccount mocks adding a subaccount to the Email Service
func (client *TestClient) AddSubaccount(subaccount *mandrill.Subaccount) (response *mandrill.Subaccount, err error) {
	if !client.OK {
		return nil, errors.New("AddSubaccount failed")
	}
	if _, ok := client.Subaccounts[subaccount.ID]; ok {
		return nil, fmt.Errorf("subaccount (%s) already exists", subaccount.ID)
	}
	client.Subaccounts[subaccount.ID] = subaccount
	return subaccount, nil
}

// TemplateInfo mocks getting a template from the Email Service
func (client *TestClient) TemplateInfo(templateName string) (response *mandrill.Template, err error) {
	if !client.OK {
		return nil, errors.New("TemplateInfo failed")
	}
	template, ok := client.Templates[templateName]
	if !ok {
		return nil, fmt.Errorf("no template (%s) found", templateName)
	}
	return template, nil
}

// DeleteTemplate mocks deleting a template from the Email Service
func (client *TestClient) DeleteTemplate(templateName string) (response *mandrill.Template, err error) {
	if !client.OK {
		return nil, errors.New("DeleteTemplate failed")
	}
	template, ok := client.Templates[templateName]
	if !ok {
		return nil, fmt.Errorf("no template (%s) found", templateName)
	}
	delete(client.Templates, templateName)
	return template, nil
}

// UpdateTemplate mocks updating a template on the Email Service
func (client *TestClient) UpdateTemplate(template *mandrill.Template) (response *mandrill.Template, err error) {
	if !client.OK {
		return nil, errors.New("UpdateTemplate failed")
	}

	if _, ok := client.Templates[template.Key]; !ok {
		return nil, fmt.Errorf("no template (%s) found", template.Key)
	}
	client.Templates[template.Key] = template
	return template, nil
}

// AddTemplate mocks adding a template to the Email Service
func (client *TestClient) AddTemplate(template *mandrill.Template) (response *mandrill.Template, err error) {
	if !client.OK {
		return nil, errors.New("AddTemplate failed")
	}

	if _, ok := client.Templates[template.Key]; ok {
		return nil, fmt.Errorf("template (%s) already exists", template.Key)
	}
	client.Templates[template.Key] = template
	return template, nil
}
