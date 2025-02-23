package mockses

import (
	"fmt"

	"github.com/arunvm123/demtech/constants"
	"github.com/arunvm123/demtech/email"
	"golang.org/x/exp/rand"
)

type MockSES struct{}

func New() *MockSES {
	m := MockSES{}
	return &m
}

func (m MockSES) SendEmail(args email.SendEmailInput) (interface{}, error) {
	return mockSendEmail(args.Scenario), nil
}

func mockSendEmail(scenario string) interface{} {
	requestID := fmt.Sprintf("%x", randBytes(16))
	messageID := fmt.Sprintf("0100017272bd51d5-%s-000000", randBytes(12))

	switch scenario {
	case constants.MockScenarioSuccess:
		return email.SendEmailResponse{
			MessageId: messageID,
		}

	case constants.MockScenarioUnverifiedEmail:
		return email.ErrorResponse{
			Code:      constants.ErrCodeMessageRejected,
			Message:   "Email address is not verified.",
			RequestId: requestID,
		}

	case constants.MockScenarioAccountSuspended:
		return email.ErrorResponse{
			Code:      constants.ErrCodeAccountSuspended,
			Message:   "Account suspended. Please contact AWS Support for assistance.",
			RequestId: requestID,
		}

	case constants.MockScenarioRateExceeded:
		return email.ErrorResponse{
			Code:      constants.ErrCodeTooManyRequests,
			Message:   "Rate exceeded. Please reduce your sending rate and try again.",
			RequestId: requestID,
		}

	case constants.MockScenarioMissingFrom:
		return email.ErrorResponse{
			Code:      constants.ErrCodeValidationException,
			Message:   "1 validation error detected: Value null at 'fromEmailAddress' failed to satisfy constraint: Member must not be null",
			RequestId: requestID,
		}

	case constants.MockScenarioDomainNotVerified:
		return email.ErrorResponse{
			Code:      constants.ErrCodeMailFromDomainNotVerified,
			Message:   "The domain you are sending from is not verified in Amazon SES.",
			RequestId: requestID,
		}

	case constants.MockScenarioQuotaExceeded:
		return email.ErrorResponse{
			Code:      constants.ErrCodeLimitExceeded,
			Message:   "Daily message quota exceeded",
			RequestId: requestID,
		}

	default:
		return email.ErrorResponse{
			Code:      constants.ErrCodeInternalServer,
			Message:   "An internal server error occurred.",
			RequestId: requestID,
		}
	}
}

// Helper function to generate random bytes
func randBytes(n int) []byte {
	const letters = "0123456789abcdef"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return bytes
}
