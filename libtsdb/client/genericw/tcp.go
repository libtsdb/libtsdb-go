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
var _ libtsdb.TracedTcpClient = (*TcpClient)(nil)

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

	// stat for single request
	trace     libtsdb.TcpTrace
	prevTrace libtsdb.TcpTrace

	// stat for accumulated counters
	// NOTE: we maintain counter in generic clients so encoder don't need to worry about it
	totalPayloadSize        int
	totalRawSize            int
	totalRawMetaSize        int
	totalIntPointWritten    int
	totalDoublePointWritten int
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
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointIntTagged(p)
}

// WriteDoublePoint only writes to encoder, but does not flush it
func (c *TcpClient) WriteDoublePoint(p *pb.PointDoubleTagged) {
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointDoubleTagged(p)
}

func (c *TcpClient) WriteSeriesIntTagged(p *pb.SeriesIntTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WriteSeriesIntTagged(p)
}

func (c *TcpClient) WriteSeriesDoubleTagged(p *pb.SeriesDoubleTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

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

func (c *TcpClient) Trace() libtsdb.Trace {
	// make a copy, otherwise when the trace is used, the pointer might be pointing to a changed trace
	cp := c.prevTrace
	return &cp
}

func (c *TcpClient) TcpTrace() libtsdb.TcpTrace {
	return c.prevTrace
}

func (c *TcpClient) send() error {
	c.totalPayloadSize += c.enc.Len()
	c.trace.PayloadSize = c.enc.Len()
	c.trace.Start = time.Now().UnixNano()
	_, err := c.conn.Write(c.enc.Bytes())
	c.trace.End = time.Now().UnixNano()
	// reset
	c.enc.Reset()
	c.prevTrace = c.trace
	c.trace.Reset()
	// TODO: retry
	if err != nil {
		c.prevTrace.Error = true
		c.prevTrace.ErrorMessage = err.Error()
		return errors.Wrap(err, "error send tcp request")
	}
	return nil
}
