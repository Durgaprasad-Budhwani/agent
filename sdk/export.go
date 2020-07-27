package sdk

// Export is a control interface for an export
type Export interface {
	Control
	// Config is any customer specific configuration for this customer
	Config() Config
	// State is a customer specific state object for this integration and customer
	State() State
	// JobID will return a specific job id for this export which can be used in logs, etc
	JobID() string
	// Pipe should be called to get the pipe for streaming data back to pinpoint
	Pipe() Pipe
	// Historical if true, the integration should perform a full historical export
	Historical() bool
}
