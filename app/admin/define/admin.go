package define

// AdminInput 注册/登录请求参数
type AdminInput struct {
	Username string `p:"username" v:"required|length:5,10"`
	Password string `p:"password" v:"required|length:6,16"`
}
