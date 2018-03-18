package genericw

import (
	"net"
	"time"

	"github.com/dyweb/gommon/errors"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/common"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
)

var _ libtsdb.TSDBClient = (*TcpClient)(nil)
var _ libtsdb.WriteClient = (*TcpClient)(nil)

// TcpClient encode points using encoder and write to raw tcp connection
// TODO: re connect when fail
// TODO: allow insecure
type TcpClient struct {
	// tsdb
	enc  common.Encoder
	meta libtsdb.Meta

	// tcp
	addr    string
	timeout time.Duration
	conn    net.Conn
	// TODO: support reconnect
	reconCount int

	// TODO: stat
	bytesSend          uint64
	bytesSendSuccess   uint64
	intPointWritten    uint64
	doublePointWritten uint64
}

// TODO: we should allow dial later
func NewTcp(meta libtsdb.Meta, encoder common.Encoder, addr string, timeout time.Duration) (*TcpClient, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "can't dial tcp")
	}
	return &TcpClient{
		meta:    meta,
		enc:     encoder,
		addr:    addr,
		timeout: timeout,
		conn:    conn,
	}, nil
}

func (c *TcpClient) Meta() libtsdb.Meta {
	return c.meta
}

// WriteIntPoint only writes to encoder, but does not flush it
func (c *TcpClient) WriteIntPoint(p *pb.PointIntTagged) {
	c.intPointWritten += 1
	c.enc.WritePointIntTagged(p)
}

// WriteDoublePoint only writes to encoder, but does not flush it
func (c *TcpClient) WriteDoublePoint(p *pb.PointDoubleTagged) {
	c.doublePointWritten += 1
	c.enc.WritePointDoubleTagged(p)
}

// WriteSeriesIntTagged only writes to encoder, but does not flush it
func (c *TcpClient) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	c.intPointWritten += uint64(len(p.Points))
	c.enc.WriteSeriesIntTagged(p)
}

// WriteSeriesIntTagged only writes to encoder, but does not flush it
func (c *TcpClient) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	c.doublePointWritten += uint64(len(p.Points))
	c.enc.WriteSeriesDoubleTagged(p)
}

func (c *TcpClient) Close() error {
	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "error closing tcp connection")
	}
	return nil
}

func (c *TcpClient) Flush() error {
	return c.send()
}

func (c *TcpClient) send() error {
	_, err := c.conn.Write(c.enc.Bytes())
	c.enc.Reset()
	// TODO: retry
	// TODO: keep stats
	if err != nil {
		return errors.Wrap(err, "error send http request")
	}
	return nil
}
