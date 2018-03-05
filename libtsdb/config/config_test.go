package config

import (
	"github.com/dyweb/gommon/util/testutil"
	"testing"
)

func TestGraphiteClient(t *testing.T) {
	var c GraphiteClientConfig
	testutil.ReadYAMLTo(t, "testdata/graphite.yml", &c)
	// NOTE: time.Duration is supported by go.yaml
	t.Log(c.Timeout)
}
