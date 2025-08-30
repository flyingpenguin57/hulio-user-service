package constant

const (
	ApiUserLogin    = "/api/v1/user/login"
	ApiUserRegister = "/api/v1/user/register"
	MockPanic       = "/api/v1/mock/panic"
	Health          = "/health"
)

var ApiWhiteList = []string{
	ApiUserLogin,
	ApiUserRegister,
	MockPanic,
	Health,
}
