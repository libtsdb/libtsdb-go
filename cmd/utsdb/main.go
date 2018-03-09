package main

import (
	"fmt"
	"runtime"
	"os"

	dlog "github.com/dyweb/gommon/log"
	icli "github.com/at15/go.ice/ice/cli"
	goicelog "github.com/at15/go.ice/ice/util/logutil"

	_ "github.com/libtsdb/libtsdb-go/libtsdb/common/graphite"
	_ "github.com/libtsdb/libtsdb-go/libtsdb/common/influxdb"
	_ "github.com/libtsdb/libtsdb-go/libtsdb/common/kairosdb"
	_ "github.com/libtsdb/libtsdb-go/libtsdb/common/xephonk"
)

const (
	myname = "utsdb"
)

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

var cli *icli.Root
var log = dlog.NewApplicationLogger()

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("Universal Time Series Database Shell"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(listCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	log.AddChild(goicelog.Registry)
}
