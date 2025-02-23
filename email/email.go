package email

type Email interface {
	SendEmail(args SendEmailInput) (interface{}, error)
}

// SendEmailInput represents the input for SendEmailV2 API
type SendEmailInput struct {
	FromEmailAddress                          string                 `json:"FromEmailAddress"`
	FromEmailAddressIdentityArn               string                 `json:"FromEmailAddressIdentityArn,omitempty"`
	Destination                               *Destination           `json:"Destination"`
	Content                                   *EmailContent          `json:"Content"`
	ReplyToAddresses                          []string               `json:"ReplyToAddresses,omitempty"`
	FeedbackForwardingEmailAddress            string                 `json:"FeedbackForwardingEmailAddress,omitempty"`
	FeedbackForwardingEmailAddressIdentityArn string                 `json:"FeedbackForwardingEmailAddressIdentityArn,omitempty"`
	EmailTags                                 []MessageTag           `json:"EmailTags,omitempty"`
	ConfigurationSetName                      string                 `json:"ConfigurationSetName,omitempty"`
	ListManagementOptions                     *ListManagementOptions `json:"ListManagementOptions,omitempty"`
	EndpointId                                string                 `json:"EndpointId,omitempty"`

	// Response is triggered based no this field
	Scenario string
}

// ListManagementOptions contains options for list management
type ListManagementOptions struct {
	ContactListName string `json:"ContactListName"`
	TopicName       string `json:"TopicName,omitempty"`
}

// Destination contains email recipients
type Destination struct {
	ToAddresses  []string `json:"ToAddresses,omitempty"`
	CcAddresses  []string `json:"CcAddresses,omitempty"`
	BccAddresses []string `json:"BccAddresses,omitempty"`
}

// EmailContent contains the email content in different formats
type EmailContent struct {
	Simple   *SimpleEmailContent   `json:"Simple,omitempty"`
	Raw      *RawMessageContent    `json:"Raw,omitempty"`
	Template *TemplateEmailContent `json:"Template,omitempty"`
}

// SimpleEmailContent represents a simple email format
type SimpleEmailContent struct {
	Subject *Content      `json:"Subject"`
	Body    *Body         `json:"Body"`
	Headers []EmailHeader `json:"Headers,omitempty"`
}

// Body contains different formats of the email body
type Body struct {
	Text *Content `json:"Text,omitempty"`
	Html *Content `json:"Html,omitempty"`
}

// Content represents the content with its charset
type Content struct {
	Data    string `json:"Data"`
	Charset string `json:"Charset,omitempty"`
}

// RawMessageContent represents raw email content
type RawMessageContent struct {
	Data []byte `json:"Data"`
}

// TemplateContent represents the content for a template
type TemplateContent struct {
	Subject string `json:"Subject"`
	Text    string `json:"Text,omitempty"`
	Html    string `json:"Html,omitempty"`
}

// TemplateEmailContent represents a templated email
type TemplateEmailContent struct {
	TemplateArn     string           `json:"TemplateArn,omitempty"`
	TemplateName    string           `json:"TemplateName,omitempty"`
	TemplateData    string           `json:"TemplateData"`
	TemplateContent *TemplateContent `json:"TemplateContent,omitempty"`
	Headers         []EmailHeader    `json:"Headers,omitempty"`
}

// EmailHeader represents an email header
type EmailHeader struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// MessageTag represents an email tag
type MessageTag struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type SendEmailResponse struct {
	MessageId string `json:"MessageId"`
}

type ErrorResponse struct {
	Message   string `json:"Message"`
	Code      string `json:"Code"`
	RequestId string `json:"RequestId"`
}
