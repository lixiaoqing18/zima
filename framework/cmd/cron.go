package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/erikdubbelboer/gspt"
	"github.com/lixiaoqing18/zima/framework"
	"github.com/lixiaoqing18/zima/framework/contract"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var cronDeamon bool

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "cron任务控制命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个cron任务",
	RunE: func(c *cobra.Command, args []string) error {
		// start命令有一个deamon参数，简写为d
		settingService := framework.MustMake(contract.SettingKey).(contract.Setting)
		cronPidFile := filepath.Join(settingService.RuntimeFolder(), "cron.pid")
		cronLogFile := filepath.Join(settingService.LogFolder(), "cron.log")
		currentFolder := settingService.BaseFolder()
		if cronDeamon {
			// 创建一个Context
			cntxt := &daemon.Context{
				// 设置pid文件
				PidFileName: cronPidFile,
				PidFilePerm: 0664,
				// 设置日志文件
				LogFileName: cronLogFile,
				LogFilePerm: 0640,
				// 设置工作路径
				WorkDir: currentFolder,
				// 设置所有设置文件的mask，默认为750
				Umask: 027,
				// 子进程的参数，按照这个参数设置，子进程的命令为 ./zima cron start --daemon=true
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			// 启动子进程，d不为空表示当前是父进程，d为空表示当前是子进程
			d, err := cntxt.Reborn()
			if err != nil {
				return err
			}
			if d != nil {
				// 父进程直接打印启动成功信息，不做任何操作
				fmt.Println("cron service started, pid:", d.Pid)
				fmt.Println("log file:", cronLogFile)
				return nil
			}

			// 子进程执行Cron.Run
			defer cntxt.Release()
			fmt.Println("daemon started")
			gspt.SetProcTitle("zima cron")
			framework.StartCron()
			return nil
		} else {
			pid := cast.ToString(os.Getpid())
			err := ioutil.WriteFile(cronPidFile, []byte(pid), 0664)
			if err != nil {
				return err
			}
			gspt.SetProcTitle("zima cron")
			framework.StartCron()
			fmt.Println("cron服务已启动,pid=", pid)
			return nil
		}
	},
}

var cronListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有cron任务",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

var cronRestartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重启cron任务进程",
	RunE: func(c *cobra.Command, args []string) error {

		return nil
	},
}

var cronStopCommand = &cobra.Command{
	Use:   "stop",
	Short: " 停止cron任务进程",
	RunE: func(c *cobra.Command, args []string) error {

		return nil
	},
}

var cronStateCommand = &cobra.Command{
	Use:   "state",
	Short: "cron任务进程状态",
	RunE: func(c *cobra.Command, args []string) error {
		return nil
	},
}

func initCronCommand() *cobra.Command {
	cronStartCommand.Flags().BoolVarP(&cronDeamon, "daemon", "d", false, "start serve daemon")
	cronCommand.AddCommand(cronStartCommand)
	return cronCommand
}
