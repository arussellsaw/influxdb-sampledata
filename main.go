package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/arussellsaw/telemetry"
	"github.com/arussellsaw/telemetry/reporters"
)

func main() {
	host := flag.String("h", "http://127.0.0.1:8086", "url to influxdb http api")
	dev := flag.Int("v", 100, "amount to modify previous value each step")
	database := flag.String("d", "testdata", "influxdb database")
	prefix := flag.String("p", "sampledata_", "metric prefix")
	metric := flag.String("m", "random-data", "name of metric")
	flag.Parse()

	var tel = telemetry.New(*prefix, (10 * time.Second))
	var reporter = reporters.InfluxReporter{
		Host:     *host,
		Interval: (30 * time.Second),
		Tel:      tel,
		Database: *database,
	}
	reporter.Report()

	var sample = telemetry.NewAverage(tel, *metric, (60 * time.Second))
	var val = 100
	var add = 0
	for {
		add = rand.Intn(*dev) - (*dev / 2)
		if val+add > 0 {
			val = val + add
		}
		sample.Add(tel, float64(val))
		fmt.Printf("added point %v \n", val)
		time.Sleep(3 * time.Second)
	}
}
