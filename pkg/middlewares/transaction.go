package middlewares

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boostef/ef"
)

func Transaction() types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		nativeCtx := ctx.Context()
		nativeCtx = ef.MustBeginTransaction(nativeCtx)
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
