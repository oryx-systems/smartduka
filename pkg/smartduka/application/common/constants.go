package common

// ContextKey is used as a type for the UID key for the Firebase *auth.Token on context.Context.
// It is a custom type in order to minimize context key collissions on the context
// (.and to shut up golint).
type ContextKey string

const (
	// DefaultRegion is the system default region
	DefaultRegion = "KE"

	// PortEnvVarName is the name of the environment variable that defines the
	// server port
	PortEnvVarName = "PORT"

	// GoogleCloudProjectIDEnvVarName is used to determine the ID of the GCP project e.g for setting up StackDriver client
	GoogleCloudProjectIDEnvVarName = "GOOGLE_CLOUD_PROJECT_ID"

	// AppName is the name of "this server"
	AppName = "api-gateway"

	// AppVersion is the app version (used for StackDriver error reporting)
	AppVersion = "0.0.1"

	// AuthTokenContextKey is the key used to store the auth token on the context.Context
	AuthTokenContextKey = ContextKey("UID")
)
