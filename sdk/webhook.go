package sdk

// WebHookScope is the scope of the webhook
type WebHookScope string

const (
	// WebHookScopeOrg is the org scope for a webhook
	WebHookScopeOrg WebHookScope = "org"
	// WebHookScopeRepo is the repo scope for a webhook
	WebHookScopeRepo WebHookScope = "repo"
	// WebHookScopeProject is the project scope for a webhook
	WebHookScopeProject WebHookScope = "project"
)

// WebHook is a control interafce for web hook data received by pinpoint on behalf of the integration
type WebHook interface {
	Control
	// Config is any customer specific configuration for this customer
	Config() Config
	// State is a customer specific state object for this integration and customer
	State() State
	// CustomerID will return the customer id for the web hook
	CustomerID() string
	// IntegrationInstanceID will return the unique instance id for this integration for a customer
	IntegrationInstanceID() string
	// RefID will return the ref id from when the hook was created
	RefID() string
	// Pipe returns a pipe for sending data back to pinpoint from the web hook data
	Pipe() Pipe
	// Data returns the payload of a webhook decoded from json into a map
	Data() (map[string]interface{}, error)
	// Bytes will return the underlying data as bytes
	Bytes() []byte
	// Headers are the headers that came from the web hook
	Headers() map[string]string
	// Scope is the registered webhook scope
	Scope() WebHookScope
}

// WebHookManager is the manager for dealing with WebHooks
type WebHookManager interface {
	// Create is used by the integration to create a webhook on behalf of the integration for a given customer, reftype and refid
	// the result will be a fully qualified URL to the webhook endpoint that should be registered with the integration
	Create(customerID string, integrationInstanceID string, refType string, refID string, scope WebHookScope, params ...string) (string, error)
	// Delete will remove the webhook from the entity based on scope
	Delete(customerID string, integrationInstanceID string, refType string, refID string, scope WebHookScope) error
	// Exists returns true if the webhook is registered for the given entity based on ref_id and scope
	Exists(customerID string, integrationInstanceID string, refType string, refID string, scope WebHookScope) bool
	// Errored will set the errored state on the webhook and the message will be the Error() value of the error
	Errored(customerID string, integrationInstanceID string, refType string, refID string, scope WebHookScope, err error)
	// HookURL will return the webhook url
	HookURL(customerID string, integrationInstanceID string, refType string, refID string, scope WebHookScope) (string, error)
}
