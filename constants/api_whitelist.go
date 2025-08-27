package constant

const (
    ApiUserLogin = "/api/v1/user/login"
	ApiUserRegister = "/api/v1/user/register"
)

var ApiWhiteList = []string{
    ApiUserLogin,
	ApiUserRegister,
}
