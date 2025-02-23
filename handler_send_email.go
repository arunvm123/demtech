package main

import (
	"net/http"

	"github.com/arunvm123/demtech/constants"
	"github.com/arunvm123/demtech/email"
	"github.com/gin-gonic/gin"
)

type handleSendEmailArgs struct {
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

func (server *server) handleSendEmail(c *gin.Context) {

	var args handleSendEmailArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Request body not properly formatted")
		return
	}

	emailInput := mapToSendEmailInput(&args)

	scenario := c.GetHeader("Scenario")

	if len(scenario) == 0 {
		scenario = constants.MockScenarioSuccess
	}

	emailInput.Scenario = scenario

	response, err := server.email.SendEmail(*emailInput)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(getResponseStausCode(response), response)
	return
}

func mapToSendEmailInput(args *handleSendEmailArgs) *email.SendEmailInput {
	if args == nil {
		return nil
	}

	input := &email.SendEmailInput{
		FromEmailAddress:                          args.FromEmailAddress,
		FromEmailAddressIdentityArn:               args.FromEmailAddressIdentityArn,
		ReplyToAddresses:                          args.ReplyToAddresses,
		FeedbackForwardingEmailAddress:            args.FeedbackForwardingEmailAddress,
		FeedbackForwardingEmailAddressIdentityArn: args.FeedbackForwardingEmailAddressIdentityArn,
		ConfigurationSetName:                      args.ConfigurationSetName,
		EndpointId:                                args.EndpointId,
	}

	// Map Destination
	if args.Destination != nil {
		input.Destination = &email.Destination{
			ToAddresses:  args.Destination.ToAddresses,
			CcAddresses:  args.Destination.CcAddresses,
			BccAddresses: args.Destination.BccAddresses,
		}
	}

	// Map ListManagementOptions
	if args.ListManagementOptions != nil {
		input.ListManagementOptions = &email.ListManagementOptions{
			ContactListName: args.ListManagementOptions.ContactListName,
			TopicName:       args.ListManagementOptions.TopicName,
		}
	}

	// Map EmailTags
	if len(args.EmailTags) > 0 {
		input.EmailTags = make([]email.MessageTag, len(args.EmailTags))
		for i, tag := range args.EmailTags {
			input.EmailTags[i] = email.MessageTag{
				Name:  tag.Name,
				Value: tag.Value,
			}
		}
	}

	// Map EmailContent
	if args.Content != nil {
		input.Content = mapEmailContent(args.Content)
	}

	return input
}

func mapEmailContent(content *EmailContent) *email.EmailContent {
	if content == nil {
		return nil
	}

	result := &email.EmailContent{}

	// Map Simple content
	if content.Simple != nil {
		result.Simple = &email.SimpleEmailContent{
			Headers: mapEmailHeaders(content.Simple.Headers),
		}

		// Map Subject
		if content.Simple.Subject != nil {
			result.Simple.Subject = &email.Content{
				Data:    content.Simple.Subject.Data,
				Charset: content.Simple.Subject.Charset,
			}
		}

		// Map Body
		if content.Simple.Body != nil {
			result.Simple.Body = &email.Body{}
			if content.Simple.Body.Text != nil {
				result.Simple.Body.Text = &email.Content{
					Data:    content.Simple.Body.Text.Data,
					Charset: content.Simple.Body.Text.Charset,
				}
			}
			if content.Simple.Body.Html != nil {
				result.Simple.Body.Html = &email.Content{
					Data:    content.Simple.Body.Html.Data,
					Charset: content.Simple.Body.Html.Charset,
				}
			}
		}
	}

	// Map Raw content
	if content.Raw != nil {
		result.Raw = &email.RawMessageContent{
			Data: content.Raw.Data,
		}
	}

	// Map Template content
	if content.Template != nil {
		result.Template = &email.TemplateEmailContent{
			TemplateArn:  content.Template.TemplateArn,
			TemplateName: content.Template.TemplateName,
			TemplateData: content.Template.TemplateData,
			Headers:      mapEmailHeaders(content.Template.Headers),
		}

		if content.Template.TemplateContent != nil {
			result.Template.TemplateContent = &email.TemplateContent{
				Subject: content.Template.TemplateContent.Subject,
				Text:    content.Template.TemplateContent.Text,
				Html:    content.Template.TemplateContent.Html,
			}
		}
	}

	return result
}

func mapEmailHeaders(headers []EmailHeader) []email.EmailHeader {
	if len(headers) == 0 {
		return nil
	}

	result := make([]email.EmailHeader, len(headers))
	for i, header := range headers {
		result[i] = email.EmailHeader{
			Name:  header.Name,
			Value: header.Value,
		}
	}
	return result
}

var errorCodeToStatusCode = map[string]int{
	constants.ErrCodeMessageRejected:           http.StatusBadRequest,
	constants.ErrCodeValidationException:       http.StatusBadRequest,
	constants.ErrCodeBadRequestException:       http.StatusBadRequest,
	constants.ErrCodeMailFromDomainNotVerified: http.StatusBadRequest,
	constants.ErrCodeAccountSuspended:          http.StatusForbidden,
	constants.ErrCodeTooManyRequests:           http.StatusTooManyRequests,
	constants.ErrCodeLimitExceeded:             http.StatusBadRequest,
	constants.ErrCodeInternalServer:            http.StatusInternalServerError,
}

func getResponseStausCode(response interface{}) int {

	var statusCode int

	switch r := response.(type) {
	case email.SendEmailResponse:
		return http.StatusOK
	case email.ErrorResponse:
		return errorCodeToStatusCode[r.Message]
	}

	return statusCode
}
