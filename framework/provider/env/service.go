package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaEnvService struct {
	container framework.Container
	//.env所在的目录
	envFileFolder string
	//所有环境变量
	envMap map[string]string
}

func NewZimaEnvService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	folder := params[1].(string)
	service := &ZimaEnvService{
		container:     c,
		envFileFolder: folder,
		envMap:        map[string]string{contract.AppEnv: contract.EnvDev},
	}

	//加载.env文件内容
	envFile := filepath.Join(folder, ".env")
	file, err := os.Open(envFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, "=", 2)
		if len(items) != 2 {
			continue
		}
		service.envMap[items[0]] = items[1]
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	//读取当前环境变量内容覆盖.env默认值
	envs := os.Environ()
	for _, v := range envs {
		items := strings.SplitN(v, "=", 2)
		if len(items) != 2 {
			continue
		}
		service.envMap[items[0]] = items[1]
	}

	return service, nil
}

// 当前APP_ENV是dev、test、prod
func (service *ZimaEnvService) AppEnv() string {
	return service.envMap[contract.AppEnv]
}

// 判断环境变量是否存在
func (service *ZimaEnvService) IsExist(key string) bool {
	_, ok := service.envMap[key]
	return ok
}

// 获取指定环境变量值
func (service *ZimaEnvService) Get(key string) string {
	if v, ok := service.envMap[key]; ok {
		return v
	}
	return ""
}

// 获取所有环境变量
func (service *ZimaEnvService) All() map[string]string {
	return service.envMap
}
