package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/oryx-systems/smartduka/pkg/smartduka/application/common/helpers"
	"github.com/pkg/errors"
	"github.com/pquerna/otp/totp"
)

const (
	// DebugEnvVarName is used to determine if we should print extended tracing / logging (debugging aids)
	// to the console
	DebugEnvVarName = "DEBUG"

	issuer      = "Oryx Developers"
	accountName = "oreondevelopers@gmail.com"
)

// WriteJSONResponse writes the content supplied via the `source` parameter to
// the supplied http ResponseWriter. The response is returned with the indicated
// status.
func WriteJSONResponse(w http.ResponseWriter, source interface{}, status int) {
	w.WriteHeader(status) // must come first...otherwise the first call to Write... sets an implicit 200
	content, errMap := json.Marshal(source)
	if errMap != nil {
		msg := fmt.Sprintf("error when marshalling %#v to JSON bytes: %#v", source, errMap)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, errMap = w.Write(content)
	if errMap != nil {
		msg := fmt.Sprintf(
			"error when writing JSON %s to http.ResponseWriter: %#v", string(content), errMap)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

// DecodeJSONToTargetStruct maps JSON from a HTTP request to a struct.
func DecodeJSONToTargetStruct(w http.ResponseWriter, r *http.Request, targetStruct interface{}) {
	err := json.NewDecoder(r.Body).Decode(targetStruct)
	if err != nil {
		WriteJSONResponse(w, ErrorMap(err), http.StatusBadRequest)
		return
	}
}

// BoolEnv gets and parses a boolean environment variable
func BoolEnv(envVarName string) bool {
	envVar, err := helpers.GetEnvVar(envVarName)
	if err != nil {
		return false
	}
	val, err := strconv.ParseBool(envVar)
	if err != nil {
		return false
	}
	return val
}

// IsDebug returns true if debug has been turned on in the environment
func IsDebug() bool {
	return BoolEnv(DebugEnvVarName)
}

// RequestDebugMiddleware dumps the incoming HTTP request to the log for inspection
func RequestDebugMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					log.Errorf("Unable to read request body for debugging: error %#v", err)
				}
				if IsDebug() {
					req, err := httputil.DumpRequest(r, true)
					if err != nil {
						log.Errorf("Unable to dump cloned request for debugging: error %#v", err)
					}
					log.Printf("Raw request: %v", string(req))
				}
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			},
		)
	}
}

// ReportErr writes the indicated error to supplied response writer and also logs it
func ReportErr(w http.ResponseWriter, err error, status int) {
	if IsDebug() {
		log.Printf("%s", err)
	}
	WriteJSONResponse(w, ErrorMap(err), status)
}

// ErrorMap turns the supplied error into a map with "error" as the key
func ErrorMap(err error) map[string]string {
	errMap := make(map[string]string)
	errMap["error"] = err.Error()
	return errMap
}

// GenerateOTP is used to generate a one time password
func GenerateOTP() (string, error) {
	opts := totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	}
	key, err := totp.Generate(opts)
	if err != nil {
		return "", errors.Wrap(err, "generateOTP")
	}

	code, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		return "", errors.Wrap(err, "generateOTP > GenerateCode")
	}

	return code, nil
}
