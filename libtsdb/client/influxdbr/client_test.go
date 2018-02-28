package influxdbr

import (
	"fmt"
	"os"
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestClient_CreateDatabase(t *testing.T) {
	t.Skip("requires influxdb running")

	assert := asst.New(t)
	c, err := New(Config{
		Addr:     "http://localhost:8086",
		Database: "libtsdbtest",
	})
	assert.Nil(err)
	res, err := c.CreateDatabase("libtsdbtest")
	if err == nil {
		t.Log(*res)
	} else {
		t.Log(err.Error())
	}
	assert.Nil(err)
}

func TestMain(m *testing.M) {
	fmt.Println("TODO: spin up database")
	code := m.Run()
	fmt.Println("TODO: tear down database")
	os.Exit(code)
}
