package main

import (
	// standard
	"fmt"

	// local
	"github.com/jheise/go-cymon/api"

	// external
	"github.com/mitchellh/go-homedir"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apipath = kingpin.Flag("apikey", "Path to API key").Default("~/.cycmd").String()
	section = kingpin.Flag("section", "Section to query, default lookup").Default("lookup").String()
	report  = kingpin.Arg("report", "What to report on [ipv4, ipv6, domain, hostname]").Required().String()
	target  = kingpin.Arg("target", "What to report on").Required().String()
	client  *api.Client
)

func init() {
	kingpin.Parse()
	var err error

	if *apipath == "~/.cycmd" {
		keypath, err := homedir.Expand(*apipath)
		if err != nil {
			panic(err)
		}

		client, err = api.NewClientFromFile(keypath)
		if err != nil {
			panic(err)
		}

	} else {
		client, err = api.NewClientFromFile(*apipath)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	switch *report {
	case "ip":
		processIP(*target, *section)
	case "domain":
		processDomain(*target, *section)
	case "url":
		processUrl(*target, *section)
	case "dga":
		processDga(*target, *section)
	default:
		fmt.Printf("Report %s not implemented\n", *report)
	}
}
