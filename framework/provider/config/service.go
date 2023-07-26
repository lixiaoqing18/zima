package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
)

const Config_Seperator = "."

type ZimaConfigService struct {
	container    framework.Container
	configFolder string
	envMap       map[string]string
	configMap    map[string]any
	configRaw    map[string][]byte
	lock         *sync.RWMutex
}

func NewZimaConfigService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	folder := params[1].(string)
	env := params[2].(map[string]string)
	service := &ZimaConfigService{
		container:    c,
		configFolder: folder,
		envMap:       env,
		configMap:    map[string]any{},
		configRaw:    map[string][]byte{},
		lock:         &sync.RWMutex{},
	}

	//读取folder下每一个配置文件
	filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			err := service.loadConfigFile(folder, info.Name())
			if err != nil {
				return err
			}
		}
		return nil
	})

	//监听配置文件变动
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(folder)
	if err != nil {
		return nil, err
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)

				if event.Has(fsnotify.Write) {
					fmt.Println("modified file:", event.Name)
					service.loadConfigFile(folder, filepath.Base(event.Name))
				} else if event.Has(fsnotify.Create) {
					fmt.Println("created file:", event.Name)
					service.loadConfigFile(folder, filepath.Base(event.Name))
				} else if event.Has(fsnotify.Remove) {
					fmt.Println("deleted file:", event.Name)
					service.removeConfigFile(folder, filepath.Base(event.Name))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	return service, nil
}

func (service *ZimaConfigService) loadConfigFile(path, filename string) error {
	service.lock.Lock()
	raw, err := os.ReadFile(filepath.Join(path, filename))
	if err != nil {
		return err
	}

	raw = RelaceWithEnv(raw, service.envMap)

	fileNames := strings.Split(filename, ".")
	fileNameExt := fileNames[1]
	if fileNameExt != "yaml" && fileNameExt != "yml" {
		return errors.New("zima config only support yaml file")
	}
	fileNameBase := fileNames[0]
	service.configRaw[fileNameBase] = raw

	mapValue := make(map[string]interface{})
	yaml.Unmarshal(raw, mapValue)
	service.configMap[fileNameBase] = mapValue

	service.lock.Unlock()

	//加载setting配置覆盖默认配置
	configMap := service.GetStringMapString("setting.path")
	if configMap != nil || len(configMap) > 0 {
		settingService := service.container.MustMake(contract.SettingKey).(contract.Setting)
		settingService.LoadConfigMap(configMap)
	}

	return nil
}

func (service *ZimaConfigService) removeConfigFile(path, filename string) error {
	service.lock.Lock()
	defer service.lock.Unlock()
	fileNames := strings.Split(filename, ".")
	fileNameExt := fileNames[1]
	if fileNameExt != "yaml" && fileNameExt != "yml" {
		return errors.New("zima config only support yaml file")
	}
	fileNameBase := fileNames[0]

	delete(service.configMap, fileNameBase)
	delete(service.configRaw, fileNameBase)

	return nil
}

func RelaceWithEnv(raw []byte, env map[string]string) []byte {
	if len(env) == 0 {
		return raw
	}
	for k, v := range env {
		key := "env(" + k + ")"
		raw = bytes.ReplaceAll(raw, []byte(key), []byte(v))
	}
	return raw
}

// IsExist 检查一个属性是否存在
func (service *ZimaConfigService) IsExist(key string) bool {

	return service.Get(key) != nil
}

// Get 获取一个属性值
func (service *ZimaConfigService) Get(key string) any {
	service.lock.RLock()
	defer service.lock.RUnlock()
	path := strings.Split(key, Config_Seperator)
	if len(path) < 2 {
		return nil
	}
	return search(path, service.configMap)
}

func search(path []string, configMap map[string]any) any {
	v, ok := configMap[path[0]]
	if ok {
		if len(path) == 1 {
			return v
		}

		switch v.(type) {
		case map[any]any:
			return search(path[1:], cast.ToStringMap(v))
		case map[string]any:
			return search(path[1:], v.(map[string]any))
		case map[string]string:
			return search(path[1:], v.(map[string]any))
		default:
			return nil
		}
	}

	return nil
}

// GetBool 获取一个 bool 属性
func (service *ZimaConfigService) GetBool(key string) bool {
	return cast.ToBool(service.Get(key))
}

// GetInt 获取一个 int 属性
func (service *ZimaConfigService) GetInt(key string) int {
	return cast.ToInt(service.Get(key))
}

// GetFloat64 获取一个 float64 属性
func (service *ZimaConfigService) GetFloat64(key string) float64 {
	return cast.ToFloat64(service.Get(key))
}

// GetTime 获取一个 time 属性
func (service *ZimaConfigService) GetTime(key string) time.Time {
	return cast.ToTime(service.Get(key))
}

// GetString 获取一个 string 属性
func (service *ZimaConfigService) GetString(key string) string {
	return cast.ToString(service.Get(key))
}

// GetIntSlice 获取一个 int 数组属性
func (service *ZimaConfigService) GetIntSlice(key string) []int {
	return cast.ToIntSlice(service.Get(key))
}

// GetStringSlice 获取一个 string 数组
func (service *ZimaConfigService) GetStringSlice(key string) []string {
	return cast.ToStringSlice(service.Get(key))
}

// GetStringMap 获取一个 string 为 key，interface 为 val 的 map
func (service *ZimaConfigService) GetStringMap(key string) map[string]any {
	return cast.ToStringMap(service.Get(key))
}

// GetStringMapString 获取一个 string 为 key，string 为 val 的 map
func (service *ZimaConfigService) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(service.Get(key))
}

// GetStringMapStringSlice 获取一个 string 为 key，数组 string 为 val 的 map
func (service *ZimaConfigService) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(service.Get(key))
}

// Load 加载配置到某个对象
func (service *ZimaConfigService) Load(key string, val any) error {
	objMap := service.Get(key)
	decoder, err := mapstructure.NewDecoder(
		&mapstructure.DecoderConfig{
			TagName: "yaml",
			Result:  val,
		})
	if err != nil {
		return err
	}
	return decoder.Decode(objMap)
}
