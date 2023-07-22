package lin_error

type linError interface {
	// Error transfer to string, satisfied the Error interface
	Error() string

	// Code return error code
	Code() ErrorCode

	// Stack return stack info
	Stack() string

	// WithMessage wrap
	WithMessage(format string, a ...any) linError
}
