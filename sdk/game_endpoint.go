/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	"github.com/grandcat/zeroconf"
)

// GameEndpoint describes a discovered game endpoint.
type GameEndpoint struct {
	Name    string
	Host    string
	Port    int
	Version string
}

// DiscoverGames attempts to discover ongoing games on the local network.
func DiscoverGames(ctx context.Context, endpoints chan<- *GameEndpoint) error {
	mdnsResolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return err
	}

	mdnsServices := make(chan *zeroconf.ServiceEntry)
	go func(services <-chan *zeroconf.ServiceEntry) {
		for service := range services {
			endpoints <- &GameEndpoint{
				Name: service.Instance,
				Host: service.AddrIPv4[0].String(),
				Port: service.Port,
				//Version: service.Text[0], // TODO
			}
		}
	}(mdnsServices)

	err = mdnsResolver.Browse(ctx, "_conreality._tcp", "local.", mdnsServices)
	if err != nil {
		return err
	}

	return nil
}