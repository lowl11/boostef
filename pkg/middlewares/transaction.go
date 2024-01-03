package middlewares

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boostef/ef"
	"time"
)

func Transaction(duration ...time.Duration) types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		var txDuration *time.Duration
		if len(duration) > 0 {
			txDuration = &duration[0]
		}

		nativeCtx := ctx.Context()

		nativeCtx = ef.MustBeginTransaction(nativeCtx, txDuration)
		ctx.SetContext(nativeCtx)
		defer ef.MustRollbackTransaction(nativeCtx)

		err := ctx.Next()
		if err != nil {
			ef.MustRollbackTransaction(nativeCtx)
			return err
		}

		ef.MustCommitTransaction(nativeCtx)
		return nil
	}
}
