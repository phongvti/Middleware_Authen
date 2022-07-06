package models

import "github.com/kataras/iris/v12/context"


type Route struct{
	Method string
	Pattern string
	HandlerFunc context.Handler
}

type Routes []Route	