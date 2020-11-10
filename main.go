package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"proxy/v1/proxy"
	"proxy/v1/server"
	"syscall"

	"github.com/Sirupsen/logrus"
)

var (
	addr    string
	irlimit int64

	logger = logrus.New()
)

func init() {
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "proxy REST API address")
	flag.Int64Var(&irlimit, "rlimit", 0, "rlimit")

	flag.Parse()
}

func setRlimit() error {
	rlimit := uint64(irlimit)

	if rlimit > 0 {
		var limit syscall.Rlimit
		if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
			return err
		}

		logger.WithFields(logrus.Fields{
			"current": limit.Cur,
			"max":     limit.Max,
		}).Info("rlimits")

		if limit.Cur < rlimit {
			limit.Cur = rlimit

			if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &limit); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	if err := setRlimit(); err != nil {
		logger.Fatalf("setting rlimit %s", err)
	}
	proxy.InitVpnTable()

	s := server.New(logger)
	go func() {
		sigc := make(chan os.Signal, 10)
		signal.Notify(sigc, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)

		for _ = range sigc {
			if err := s.Close(); err != nil {
				logger.WithField("error", err).Fatal("closing server")
			}
			os.Exit(0)
		}
	}()
	go proxy.CollectStats()
	go func() {

		http.Handle("/", http.FileServer(http.Dir("./assets")))
		http.ListenAndServe("0.0.0.0:8000", nil)
	}()

	if err := http.ListenAndServe(addr, s); err != nil {
		logger.WithField("error", err).Fatal("serving http")
	}
}
