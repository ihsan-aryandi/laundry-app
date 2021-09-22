package middleware

import (
	"github.com/bandungrhapsody/rhaprouter"
	"os"
)

func HandleCORS(next rhaprouter.Handler) rhaprouter.Handler {
	return func(ctx *rhaprouter.Context) error {
		ctx.SetHeader("Access-Control-Allow-Origin", os.Getenv("CLIENT_HOST"))
		ctx.SetHeader("Access-Control-Allow-Credentials", "true")
		ctx.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.SetHeader("Access-Control-Max-Age", "1209600")
		return next(ctx)
	}
}
