/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/basic_auth"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var userKey = "userInfo"

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	// your-realm:   安全域字符串，本例中会以 Www-Authenticate: Basic realm="your-realm" 的形式保存在响应头中
	// your-userKey: 认证通过后会以 userKey 为键 username 为值的形式设置在上下文中
	h.Use(basic_auth.BasicAuthForRealm(map[string]string{
		"username3": "password3",
		"username4": "password4",
	}, "Authorization Required", userKey))

	h.GET("/basicAuth", func(ctx context.Context, c *app.RequestContext) {
		user, exists := c.Get(userKey)
		if exists != true {
			c.String(consts.StatusCreated, "no user")
		}
		println("user", user)
		c.String(consts.StatusOK, "user:"+user.(string))
	})

	h.Spin()
}
