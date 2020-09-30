package utils

// ErrorCode godoc
// Definition of errorcode
type ErrorCode struct {
	Code    int
	Message string
}

var (
	// UnknownError godoc
	// Unable to connect to database or redis and other errors
	UnknownError = &ErrorCode{Code: 10001, Message: "Unknown error."}
	// AccessDenied godoc
	// User not logged in
	AccessDenied = &ErrorCode{Code: 10002, Message: "Access denied."}
	// PermissionDenied godoc
	// User not logged in
	PermissionDenied = &ErrorCode{Code: 10003, Message: "Permission denied."}
	// InvalidCredentials godoc
	// The user name or password incorrect.
	InvalidCredentials = &ErrorCode{Code: 20001, Message: "Invalid credentials."}
	// TooManyRetry godoc
	TooManyRetry = &ErrorCode{Code: 20002, Message: "Too many retrying accounts are locked. Please try again later."}
	// TheNameRequired godoc
	TheNameRequired = &ErrorCode{Code: 20002, Message: "The name is required."}
	// ThePasswordRequired godoc
	ThePasswordRequired = &ErrorCode{Code: 20003, Message: "The password is required."}
	// TheTaskDoesNotExist godoc
	TheTaskDoesNotExist = &ErrorCode{Code: 30001, Message: "The task does not exist."}
	// TheColumnDoesNotExist godoc
	TheColumnDoesNotExist = &ErrorCode{Code: 30002, Message: "The column does not exist."}
	// TheColumnRequired godoc
	TheColumnRequired = &ErrorCode{Code: 30003, Message: "The column id is required."}
	// TheTaskNameRequired godoc
	TheTaskNameRequired = &ErrorCode{Code: 30004, Message: "The task name is required."}
	// TheTaskNameLengthExceed godoc
	TheTaskNameLengthExceed = &ErrorCode{Code: 30005, Message: "Task name cannot exceed 100 characters in length."}
)
