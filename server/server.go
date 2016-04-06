package server

import (
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/alyoshka/calc-server/gen-go/calculate"
)

// Run - starts new server on specified address
func Run(addr string) error {
	protocolFactory := thrift.NewTJSONProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	handler := newCalculateHandler()
	processor := calculate.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	log.Println("Starting the server on ", addr)
	return server.Serve()
}
