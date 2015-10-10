package megos

import (
	"fmt"
	"testing"
)

func TestGetURLForStateFile(t *testing.T) {
	setup()
	defer teardown()

	masterNode := client.Master[1]
	u := client.GetURLForStateFile(*masterNode)

	ms := masterNode.String() + "/master/state.json"
	if us := u.String(); ms != us {
		t.Errorf("URLs not the same. Expected %s, got %s", ms, us)
	}
}

func TestGetURLForStateFilePid(t *testing.T) {
	setup()
	defer teardown()

	p := Pid{
		Role: "master",
		Host: "192.168.1.6",
		Port: 5050,
	}

	u := client.GetURLForStateFilePid(p)

	ps := fmt.Sprintf("http://%s:%d/master/state.json", p.Host, p.Port)
	if us := u.String(); ps != us {
		t.Errorf("URLs not the same. Expected %s, got %s", ps, us)
	}
}

// Missing tests:
// * GetStateFromCluster
// * GetStateFromLeader
// * GetStateFromPid
