// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	user "github.com/cloudwego/hertz-examples/bizdemo/hertz_session/biz/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)
}
