package megos

import (
	"reflect"
	"testing"
)

func TestGetFrameworkByPrefix_WithFramework(t *testing.T) {
	setup()

	prefix := "Test Framework three"
	frameworks := []Framework{
		Framework{ID: "Framework1", Name: "Test Framework one"},
		Framework{ID: "Framework2", Name: "Test Framework two"},
		Framework{ID: "Framework3", Name: "Test Framework three is really cool"},
	}

	if f, err := client.GetFrameworkByPrefix(frameworks, prefix); !reflect.DeepEqual(f, &frameworks[2]) {
		t.Errorf("Framework is not the one as expected (%s). Expected %+v, got %+v", err, &frameworks[2], f)
	}
}

func TestGetFrameworkByPrefix_WithoutFramework(t *testing.T) {
	setup()

	prefix := "test"
	frameworks := []Framework{
		Framework{ID: "Framework1", Name: "Test Framework one"},
		Framework{ID: "Framework2", Name: "Test Framework two"},
		Framework{ID: "Framework3IsReallyCool", Name: "Test Framework three"},
	}

	f, err := client.GetFrameworkByPrefix(frameworks, prefix)
	if f != nil {
		t.Errorf("Framework is not nil. Expected nil, got %+v", f)
	}
	if err == nil {
		t.Errorf("err is not nil. Expected a string, got %s", err)
	}
}
