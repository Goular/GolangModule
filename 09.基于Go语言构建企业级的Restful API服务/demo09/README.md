使用 go-playground/validator 作为表单校验库

在 gin 中可以设置 3 种类型的 middleware：

全局中间件
单个路由中间件
群组中间件

全局中间件：注册中间件的过程之前设置的路由，将不会受注册的中间件所影响。只有注册了中间件之后代码的路由函数规则，才会被中间件装饰。
单个路由中间件：需要在注册路由时注册中间件
r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
群组中间件：只要在群组路由上注册中间件函数即可。

JWT由三段内容构成
即header，payload，signature三部分构成

如何进行 API 身份验证
API 身份认证包括两步：

签发 token
API 添加认证 middleware

可以看到携带 token 后验证通过，成功创建用户。通过 HTTP Header Authorization: Bearer $token 来携带 token。携带 token 后不需要再次查询数据库核对密码，即可完成授权。
