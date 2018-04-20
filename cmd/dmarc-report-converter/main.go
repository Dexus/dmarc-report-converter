package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var version string

func main() {
	var (
		flagVersion bool
		//	flagConfig  string
		flagInputDir     string
		flagOutputDir    string
		flagOutputFormat string
		flagLookupAddr   bool
		flagDelete       bool
	)

	flag.BoolVar(&flagVersion, "version", false, "show version and exit")
	flag.StringVar(&flagInputDir, "in", "", "path to input directory")
	flag.StringVar(&flagOutputDir, "out", "", "path to output directory (if is not set, print to stdout)")
	flag.StringVar(&flagOutputFormat, "format", "txt", "output format (text, html, json)")
	flag.BoolVar(&flagLookupAddr, "lookup", false, "performs a reverse lookups")
	flag.BoolVar(&flagDelete, "delete", false, "delete source file after conversation")
	//flag.StringVar(&flagConfig, "config", "./config.yaml", "Path to config file")
	flag.Parse()

	if flagVersion {
		fmt.Printf("Version: %v\n", version)
		os.Exit(0)
	}

	if flagInputDir == "" {
		log.Fatal("-in is not set")
	}

	cfg, err := newConfig(flagOutputDir, flagOutputFormat, flagLookupAddr)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	inFiles, err := filepath.Glob(filepath.Join(flagInputDir, "*.*"))
	if err != nil {
		log.Fatal(err)
	}

	if flagLookupAddr {
		log.Printf("performs a reverse lookups, this may take some time")
	}

	log.Printf("found input files: %v", inFiles)
	for _, f := range inFiles {
		err = convertFile(f, cfg)
		if err != nil {
			log.Printf("[ERROR] %v, skip", err)
			continue
		}

		if flagDelete {
			log.Printf("delete %v after converting", f)
			err = os.Remove(f)
			if err != nil {
				log.Printf("[ERROR] %v, skip", err)
				continue
			}
		}
	}
}
