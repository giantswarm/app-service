package flag

import (
	"github.com/giantswarm/microkit/flag"

	"github.com/giantswarm/app-service/flag/service"
)

type Flag struct {
	Service service.Service
}

func New() *Flag {
	f := &Flag{}
	flag.Init(f)
	return f
}
