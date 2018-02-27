// Package client contains client for multiple time series databases, note write and read clients are separated to
// reduce dependencies in write client, and it's common when using time series database, some clients always write
// (i.e. monitor agents), some clients always read (i.e. visualization dashboards)
package client
