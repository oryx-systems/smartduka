package testutils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/oryx-systems/smartduka/pkg/smartduka/application/utils"
)

// CheckIfCurrentDBIsLocal checks whether the database used to run the test is a test/local database. If not, the setup exits
func CheckIfCurrentDBIsLocal() bool {
	isLocal, err := strconv.ParseBool(os.Getenv("IS_LOCAL_DB"))
	if err != nil {
		return false
	}

	return isLocal
}

// PrepareServer is the signature of a function that Knows how to prepare & initialise the server
type PrepareServer func(ctx context.Context, port int, allowedOrigins []string) *http.Server

func randomPort() int {
	rand.Seed(time.Now().Unix())
	min := 32768
	max := 60999
	/* #nosec G404 */
	port := rand.Intn(max-min+1) + min
	return port
}

// StartTestServer starts up test server
func StartTestServer(ctx context.Context, prepareServer PrepareServer, allowedOrigins []string) (*http.Server, string, error) {
	// prepare the server
	port := randomPort()
	srv := prepareServer(ctx, port, allowedOrigins)
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	if srv == nil {
		return nil, "", fmt.Errorf("nil test server")
	}

	// set up the TCP listener
	// this is done early so that we are sure we can connect to the port in
	// the tests; backlogs will be sent to the listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil || l == nil {
		return nil, "", fmt.Errorf("unable to listen on port %d: %w", port, err)
	}

	if utils.IsDebug() {
		log.Printf("LISTENING on port %d", port)
	}

	// start serving
	go func() {
		err := srv.Serve(l)
		if err != nil {
			if utils.IsDebug() {
				log.Printf("serve error: %s", err)
			}
		}
	}()

	// the cleanup of this server (deferred shutdown) needs to occur in the
	// acceptance test that will use this
	return srv, baseURL, nil
}
