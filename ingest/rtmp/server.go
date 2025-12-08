package rtmp

import (
	"io"
	"net"

	logger "github.com/sirupsen/logrus"
	"github.com/yutopp/go-rtmp"

	"shoong.com/ingest/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	s := new(Server)

	s.cfg = cfg

	return s
}

func (s *Server) Server() (*net.TCPListener, *rtmp.Server) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")

	if err != nil {
		logger.Panicf("Failed to resolve TCP address: %+v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		logger.Panicf("Failed to listen TCP: %+v", err)
	}

	server := rtmp.NewServer(&rtmp.ServerConfig{
		OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
			l := logger.StandardLogger()

			logger.Tracef("Client connected to Client: %v", conn.RemoteAddr().String())

			h := &Handler{}

			return conn, &rtmp.ConnConfig{
				Logger:  l,
				Handler: h,
				ControlState: rtmp.StreamControlStateConfig{
					DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
				},
			}
		},
	})

	return listener, server
}
