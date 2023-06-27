package framework

import (
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

var zimaCron *cron.Cron
var cronSpecs []CronSpec

type CronSpec struct {
	Type string
	Func func()
	Cmd  *cobra.Command
	Spec string
}

func StartCron() {
	if zimaCron != nil {
		//异步运行
		//zimaCron.Start()

		//同步运行
		zimaCron.Run()
	}
}

func AddCron(spec string, f func(), cmd *cobra.Command) {
	if zimaCron == nil {
		zimaCron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		cronSpecs = []CronSpec{}
	}

	cronSpecs = append(cronSpecs, CronSpec{
		Type: "common-cron",
		Func: f,
		Cmd:  cmd,
		Spec: spec,
	})

	zimaCron.AddFunc(spec, f)

}

func ListCronSpec() [][]string {
	result := [][]string{}
	for _, v := range cronSpecs {
		line := []string{v.Type, v.Spec, v.Cmd.Use, v.Cmd.Short}
		result = append(result, line)
	}
	return result
}
