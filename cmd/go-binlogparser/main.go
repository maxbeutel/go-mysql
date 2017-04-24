package main

import (
	"flag"
	"os"

	"github.com/siddontang/go-mysql/replication"
)

var name = flag.String("name", "", "binlog file name")

func main() {
	flag.Parse()

	p := replication.NewBinlogParser()

	f := func(e *replication.BinlogEvent) error {
		e.Dump(os.Stdout)
		return nil
	}

	err := p.ParseFile(*name, f)

	if err != nil {
		println(err)
	}
}
