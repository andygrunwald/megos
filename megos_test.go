package megos

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	// client is the Megos client being tested.
	client *Client

	// master is a list of (faked) mesos master nodes
	master []*url.URL

	// mux1 is the HTTP request multiplexer used with the test server.
	mux1 *http.ServeMux

	// server1 is a test HTTP server used to provide mock API responses.
	server1 *httptest.Server

	// mux2 is the HTTP request multiplexer used with the test server.
	mux2 *http.ServeMux

	// server2 is a test HTTP server used to provide mock API responses.
	server2 *httptest.Server

	// mux3 is the HTTP request multiplexer used with the test server.
	mux3 *http.ServeMux

	// server3 is a test HTTP server used to provide mock API responses.
	server3 *httptest.Server
)

// setup sets up a test HTTP server along with a megos.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the http call being tested.
func setup() {
	// test server: 1
	mux1 = http.NewServeMux()
	server1 = httptest.NewServer(mux1)

	// test server: 2
	mux2 = http.NewServeMux()
	server2 = httptest.NewServer(mux2)

	// test server: 3
	mux3 = http.NewServeMux()
	server3 = httptest.NewServer(mux3)

	m1, _ := url.Parse(server1.URL)
	m2, _ := url.Parse(server2.URL)
	m3, _ := url.Parse(server3.URL)
	master = []*url.URL{m1, m2, m3}

	client = NewClient(master)
}

// teardown closes the test HTTP server.
func teardown() {
	server1.Close()
	server2.Close()
	server3.Close()
}

// testMethod is a utility function to test the request method provided in want
func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	setup()
	defer teardown()

	if client == nil {
		t.Error("Megos client is nil. Expected megos.Client structure")
	}

	if !reflect.DeepEqual(client.Master, master) {
		t.Error("Megos master are not the same as initialized.")
	}
}

func TestParsePidInformation_WithPort(t *testing.T) {
	role := "master"
	host := "10.1.1.12"
	port := 5555
	pid := fmt.Sprintf("%s@%s:%d", role, host, port)
	parsedPid := client.ParsePidInformation(pid)

	if parsedPid.Role != role {
		t.Errorf("Role is not equal. Expected %s, got %s", role, parsedPid.Role)
	}
	if parsedPid.Host != host {
		t.Errorf("Host is not equal. Expected %s, got %s", host, parsedPid.Host)
	}
	if parsedPid.Port != port {
		t.Errorf("Port is not equal. Expected %d, got %d", port, parsedPid.Port)
	}
}

func TestParsePidInformation_WithoutPort(t *testing.T) {
	role := "master"
	host := "10.1.1.12"
	port := 5050
	pid := fmt.Sprintf("%s@%s", role, host)
	parsedPid := client.ParsePidInformation(pid)

	if parsedPid.Role != role {
		t.Errorf("Role is not equal. Expected %s, got %s", role, parsedPid.Role)
	}
	if parsedPid.Host != host {
		t.Errorf("Host is not equal. Expected %s, got %s", host, parsedPid.Host)
	}
	if parsedPid.Port != port {
		t.Errorf("Port is not equal. Expected %d, got %d", port, parsedPid.Port)
	}
}

func TestParsePidInformation_String(t *testing.T) {
	role := "master"
	host := "10.1.1.12"
	port := 5555
	pid := fmt.Sprintf("%s@%s:%d", role, host, port)
	parsedPid := client.ParsePidInformation(pid)

	if s := parsedPid.String(); s != pid {
		t.Errorf("Stringer of pid is not equal. Expected %s, got %s", pid, s)
	}
}

// Missing tests:
// * DetermineLeader
