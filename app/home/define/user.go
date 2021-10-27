package define

// UserInput 用户登录请求参数
type UserInput struct {
	Username string `p:"username" v:"required|length:5,10"`
	Password string `p:"password" v:"required|length:6,16"`
}
