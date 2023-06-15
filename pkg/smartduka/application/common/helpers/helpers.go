package helpers

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common"
	log "github.com/sirupsen/logrus"
	"go.opencensus.io/trace"
)

// GetEnvVar retrieves the environment variable with the supplied name and fails
// if it is not able to do so
func GetEnvVar(envVarName string) (string, error) {
	envVar := os.Getenv(envVarName)
	if envVar == "" {
		envErrMsg := fmt.Sprintf("the environment variable '%s' is not set", envVarName)
		return "", fmt.Errorf(envErrMsg)
	}
	return envVar, nil
}

// MustGetEnvVar returns the value of the environment variable with the indicated name or panics.
// It is intended to be used in the INTERNALS of the server when we can guarantee (through orderly
// coding) that the environment variable was set at server startup.
// Since the env is required, kill the app if the env is not set. In the event a variable is not super
// required, set a sensible default or don't call this method
func MustGetEnvVar(envVarName string) string {
	val, err := GetEnvVar(envVarName)
	if err != nil {
		msg := fmt.Sprintf("mandatory environment variable %s not found", envVarName)
		log.Panicf(msg)
		os.Exit(1)
	}
	return val
}

// LogStartupError is used to e.g log fatal startup errors.
// It logs, attempts to report the error to StackDriver then panics/crashes.
func LogStartupError(ctx context.Context, err error) {
	errorClient := StackDriver(ctx)
	if err != nil {
		if errorClient != nil {
			errorClient.Report(errorreporting.Entry{Error: err})
		}
		log.WithFields(log.Fields{"error": err}).Error("Server startup error")
	}
}

// StackDriver initializes StackDriver logging, error reporting, profiling etc
func StackDriver(ctx context.Context) *errorreporting.Client {
	// project setup
	projectID, err := GetEnvVar(common.GoogleCloudProjectIDEnvVarName)
	if err != nil {
		log.WithFields(log.Fields{
			"environment variable name": common.GoogleCloudProjectIDEnvVarName,
			"error":                     err,
		}).Error("Unable to determine the Google Cloud Project, can't set up StackDriver")
		return nil
	}

	// logging
	loggingClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.WithFields(log.Fields{
			"project ID": projectID,
			"error":      err,
		}).Error("Unable to initialize logging client")
		return nil
	}
	defer CloseStackDriverLoggingClient(loggingClient)

	// error reporting
	errorClient, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName: common.AppName,
		OnError: func(err error) {
			log.WithFields(log.Fields{
				"project ID":   projectID,
				"service name": common.AppName,
				"error":        err,
			}).Info("Unable to initialize error client")
		},
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Unable to initialize error client")
		return nil
	}
	defer CloseStackDriverErrorClient(errorClient)

	// tracing
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: projectID,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"project ID": projectID,
			"error":      err,
		}).Info("Unable to initialize tracing")
		return errorClient // the error client is already initialized, return it
	}
	trace.RegisterExporter(exporter)

	// profiler
	err = profiler.Start(profiler.Config{
		Service:        common.AppName,
		ServiceVersion: common.AppVersion,
		ProjectID:      projectID,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"project ID":      projectID,
			"service name":    common.AppName,
			"service version": common.AppVersion,
			"error":           err,
		}).Info("Unable to initialize profiling")
		return errorClient // the error client is already initialized, return it
	}

	return errorClient
}

// CloseStackDriverLoggingClient closes a StackDriver logging client and logs any arising error.
//
// It was written to be defer()'d in servrer initialization code.
func CloseStackDriverLoggingClient(loggingClient *logging.Client) {
	err := loggingClient.Close()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Unable to close StackDriver logging client")
	}
}

// CloseStackDriverErrorClient closes a StackDriver error client and logs any arising error.
//
// It was written to be defer()'d in servrer initialization code.
func CloseStackDriverErrorClient(errorClient *errorreporting.Client) {
	err := errorClient.Close()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Unable to close StackDriver error client")
	}
}

// GetUserTokenFromContext retrieves a Firebase *auth.Token from the supplied context
func GetUserTokenFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(common.AuthTokenContextKey)
	if val == nil {
		return "", fmt.Errorf(
			"unable to get auth token from context with key %#v", common.AuthTokenContextKey)
	}

	token := val.(string)

	return token, nil
}

// GetPinExpiryDate returns the expiry date for the given pin
func GetPinExpiryDate() (*time.Time, error) {
	pinExpiryDays := MustGetEnvVar("PIN_EXPIRY_DAYS")
	pinExpiryInt, err := strconv.Atoi(pinExpiryDays)
	if err != nil {
		return nil, fmt.Errorf("failed to convert PIN expiry days to int: %v", err)
	}
	expiryDate := time.Now().AddDate(0, 0, pinExpiryInt)

	return &expiryDate, nil
}
