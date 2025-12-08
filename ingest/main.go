package main

import (
	"flag"

	logger "github.com/sirupsen/logrus"

	"shoong.com/ingest/config"
	r "shoong.com/ingest/rtmp"
)

var configFlag = flag.String("config", "config.toml", "Path to config file")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)

	s := r.NewServer(cfg)

	listener, server := s.Server()

	if err := server.Serve(listener); err != nil {
		logger.Panicf("Failed to serve: %+v", err)
	}
}
