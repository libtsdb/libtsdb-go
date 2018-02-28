package genericw

import (
	"net/http"

	"github.com/libtsdb/libtsdb-go/libtsdb/common"
)

// Client is a generic HTTP based client
type Client struct {
	enc common.Encoder
	h   *http.Client
}
