package mockses

import "github.com/arunvm123/demtech/email"

type MockSES struct{}

func New() *MockSES {
	m := MockSES{}
	return &m
}

func (m MockSES) SendEmail(args email.SendEmailInput) (interface{}, error) {
	return email.SendEmailResponse{
		MessageId: "mock-message-id-123",
		RequestId: "mock-request-id-456",
	}, nil
}
