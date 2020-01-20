package protocol

import (
	"net"
	"time"

	"github.com/dyweb/gommon/errors"
	"github.com/libtsdb/libtsdb-go/tspb"
)

// TCPClient encode points using encoder and write to raw tcp connection
// TODO: re connect when fail
// TODO: allow insecure
type TCPClient struct {
	// tsdb
	enc Encoder

	// tcp
	addr    string
	timeout time.Duration
	conn    net.Conn
	// TODO: support reconnect
	reconCount int

	// stat for single request
	trace     TcpTrace
	prevTrace TcpTrace

	// stat for accumulated counters
	// NOTE: we maintain counter in generic clients so encoder don't need to worry about it
	totalPayloadSize        int
	totalRawSize            int
	totalRawMetaSize        int
	totalIntPointWritten    int
	totalDoublePointWritten int
}

// TODO: we should allow dial when send
func NewTCPClient(encoder Encoder, addr string, timeout time.Duration) (*TCPClient, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "can't dial tcp")
	}
	return &TCPClient{
		enc:     encoder,
		addr:    addr,
		timeout: timeout,
		conn:    conn,
	}, nil
}

// WriteIntPoint only writes to encoder, but does not flush it
func (c *TCPClient) WriteIntPoint(p *tspb.PointIntTagged) {
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointIntTagged(p)
}

// WriteDoublePoint only writes to encoder, but does not flush it
func (c *TCPClient) WriteDoublePoint(p *tspb.PointDoubleTagged) {
	c.trace.Points += 1
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += 1
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WritePointDoubleTagged(p)
}

func (c *TCPClient) WriteSeriesIntTagged(p *tspb.SeriesIntTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalIntPointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WriteSeriesIntTagged(p)
}

func (c *TCPClient) WriteSeriesDoubleTagged(p *tspb.SeriesDoubleTagged) {
	c.trace.Points += len(p.Points)
	c.trace.RawSize += p.RawSize()
	c.trace.RawMetaSize += p.RawMetaSize()

	c.totalDoublePointWritten += len(p.Points)
	c.totalRawSize += p.RawSize()
	c.totalRawMetaSize += p.RawMetaSize()

	c.enc.WriteSeriesDoubleTagged(p)
}

func (c *TCPClient) Close() error {
	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "error closing tcp connection")
	}
	return nil
}

func (c *TCPClient) Flush() error {
	return c.send()
}

func (c *TCPClient) Trace() Trace {
	// make a copy, otherwise when the trace is used, the pointer might be pointing to a changed trace
	cp := c.prevTrace
	return &cp
}

func (c *TCPClient) TcpTrace() TcpTrace {
	return c.prevTrace
}

func (c *TCPClient) send() error {
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
