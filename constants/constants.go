package constants

// Error Codes
const (
	// Validation and Request Errors
	ErrCodeMessageRejected           = "MessageRejected"
	ErrCodeValidationException       = "ValidationException"
	ErrCodeBadRequestException       = "BadRequestException"
	ErrCodeMailFromDomainNotVerified = "MailFromDomainNotVerifiedException"

	// Authentication and Authorization Errors
	ErrCodeAccountSuspended = "AccountSuspendedException"

	// Rate and Quota Errors
	ErrCodeTooManyRequests = "TooManyRequestsException"
	ErrCodeLimitExceeded   = "LimitExceededException"

	// Server Errors
	ErrCodeInternalServer = "InternalServerError"
)

// Mock scenarios for SES email responses
const (
	// Success scenario
	MockScenarioSuccess = "success"

	// Error scenarios
	MockScenarioUnverifiedEmail   = "unverified_email"
	MockScenarioAccountSuspended  = "account_suspended"
	MockScenarioRateExceeded      = "rate_exceeded"
	MockScenarioMissingFrom       = "missing_from"
	MockScenarioDomainNotVerified = "domain_not_verified"
	MockScenarioQuotaExceeded     = "daily_quota_exceeded"
)
