/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"

	rpc "github.com/conreality/conreality.go/rpc"
	"github.com/pkg/errors"
)

// StopGame TODO...
func (session *Session) StopGame(ctx context.Context) error {
	request := &rpc.Nothing{}
	_, err := session.client.session.StopGame(ctx, request)
	if err != nil {
		return errors.Wrap(err, "StopGame failed")
	}
	return nil
}
