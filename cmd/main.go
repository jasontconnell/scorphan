package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jasontconnell/scorphan/conf"
	"github.com/jasontconnell/scorphan/process"
)

func main() {
	c := flag.String("c", "config.json", "config filename")
	path := flag.String("path", "", "sitecore path of items to check for orphans")
	output := flag.String("output", "out.csv", "output filename")
	flag.Parse()

	start := time.Now()

	if *c == "" || *path == "" {
		flag.PrintDefaults()
		return
	}

	cfg := conf.LoadConfig(*c)

	log.Println("loading items from sitecore")
	m, err := process.LoadItems(cfg.ConnectionString, cfg.ProtobufLocation, *path)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("loaded", len(m), "items")

	log.Println("loading potential values from sitecore database")
	values, err := process.GetValues(cfg.ConnectionString)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("loaded", len(values), "values")

	log.Println("finding orphans")
	orphans := process.FindOrpahs(m, values)

	buf := bytes.NewBufferString("")
	for _, orphan := range orphans {
		buf.WriteString(fmt.Sprintf("%s, %s\n", "{"+strings.ToUpper(orphan.GetId().String())+"}", orphan.GetPath()))
	}
	err = os.WriteFile(*output, buf.Bytes(), os.ModePerm)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(len(orphans), "orphans out of ", len(m))
	log.Println("finished.", time.Since(start))
}
