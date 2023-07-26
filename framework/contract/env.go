package contract

const (
	EnvKey = "zima:env"

	AppEnv  = "APP_ENV"
	EnvDev  = "dev"
	EnvTest = "test"
	EnvProd = "prod"
)

type Env interface {
	//当前APP_ENV是dev、test、prod
	AppEnv() string
	//判断环境变量是否存在
	IsExist(string) bool
	//获取指定环境变量值
	Get(string) string
	//获取所有环境变量
	All() map[string]string
}
