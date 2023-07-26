package distributed

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
)

type ZimaDistributedFileLockService struct {
	container framework.Container
}

func NewZimaDistributedFileLockService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	return &ZimaDistributedFileLockService{
		container: c,
	}, nil
}

func (service *ZimaDistributedFileLockService) Select(serviceName string, appID string, lockTime time.Duration) (string, error) {
	settingService := framework.MustMake(contract.SettingKey).(contract.Setting)
	lockFilePath := filepath.Join(settingService.RuntimeFolder(), serviceName+".lock")

	// windows 文件不存在才创建，如果杀进程，文件不会清除，下次运行会导致永远无法获取文件锁
	file, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		//文件被锁定
		fmt.Println(err)
		currentAppIDByte, err := os.ReadFile(lockFilePath)
		if err != nil {
			return "", err
		}
		return string(currentAppIDByte), nil
	}

	//延迟释放文件锁
	go func() {
		defer func() {
			file.Close()
			os.Remove(lockFilePath)
		}()
		timer := time.NewTimer(lockTime)
		<-timer.C
	}()

	//写入appID到文件
	if _, err := file.WriteString(appID); err != nil {
		return "", err
	}
	return appID, nil
}
