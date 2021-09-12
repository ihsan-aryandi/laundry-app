package middleware

import (
	"github.com/bandungrhapsody/rhaprouter"
	"os"
)

func HandleCORS(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		ctx.SetHeader("Access-Control-Allow-Origin", os.Getenv("CLIENT_HOST"))
		ctx.SetHeader("Access-Control-Allow-Credentials", "true")

		return next(ctx)
	}
}
