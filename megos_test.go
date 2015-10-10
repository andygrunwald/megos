package megos

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

var (
	// client is the Megos client being tested.
	client *Client
)

func setup() {
	client = NewClient(getSetupMaster())
}

func getSetupMaster() []*url.URL {
	m1, _ := url.Parse("http://mesos-master1:5050/")
	m2, _ := url.Parse("http://mesos-master2:5050/")
	m3, _ := url.Parse("http://mesos-master3:5050/")
	master := []*url.URL{m1, m2, m3}

	return master
}

func TestNewClient(t *testing.T) {
	setup()
	if client == nil {
		t.Error("Megos client is nil. Expected megos.Client structure")
	}

	if !reflect.DeepEqual(client.Master, getSetupMaster()) {
		t.Error("Megos master are not the same as initialized.")
	}
}

func TestParsePidInformation_WithPort(t *testing.T) {
	setup()
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
	setup()
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
