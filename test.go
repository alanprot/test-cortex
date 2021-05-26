package main

import (
	"flag"
	"github.com/weaveworks/common/middleware"
	"io/ioutil"
	"os"

	"github.com/cortexproject/cortex/pkg/cortex"
	"github.com/cortexproject/cortex/pkg/util/flagext"
	util_log "github.com/cortexproject/cortex/pkg/util/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	configFileOption = "config.file"
)

func main() {
	var (
		cfg cortex.Config
	)

	configFile := parseConfigFileParameter(os.Args[1:])

	flagext.RegisterFlags(&cfg)

	LoadConfig(configFile, &cfg)

	flagext.IgnoredFlag(flag.CommandLine, configFileOption, "Configuration file to load.")
	flag.CommandLine.Init(flag.CommandLine.Name(), flag.ContinueOnError)
	flag.CommandLine.Parse(os.Args[1:])

	util_log.InitLogger(&cfg.Server)

	cfg.Server.HTTPMiddleware = []middleware.Interface {
		TestMiddleware{},
	}

	t, err := cortex.New(cfg)

	util_log.CheckFatal("running cortex", err)

	err = t.Run()

	util_log.CheckFatal("running cortex", err)
}

func LoadConfig(filename string, cfg *cortex.Config) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}

	err = yaml.UnmarshalStrict(buf, cfg)
	if err != nil {
		return errors.Wrap(err, "Error parsing config file")
	}

	return nil
}

// Parse -config.file and -config.expand-env option via separate flag set, to avoid polluting default one and calling flag.Parse on it twice.
func parseConfigFileParameter(args []string) (configFile string) {
	// ignore errors and any output here. Any flag errors will be reported by main flag.Parse() call.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)

	// usage not used in these functions.
	fs.StringVar(&configFile, configFileOption, "", "")

	// Try to find -config.file and -config.expand-env option in the flags. As Parsing stops on the first error, eg. unknown flag, we simply
	// try remaining parameters until we find config flag, or there are no params left.
	// (ContinueOnError just means that flag.Parse doesn't call panic or os.Exit, but it returns error, which we ignore)
	for len(args) > 0 {
		_ = fs.Parse(args)
		args = args[1:]
	}

	return
}
