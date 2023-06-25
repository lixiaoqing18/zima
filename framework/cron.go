package framework

import (
	"github.com/robfig/cron/v3"
)

var zimaCron *cron.Cron
var cronSpecs []CronSpec

type CronSpec struct {
	Type string
	Func func()
	Spec string
}

func StartCron() {
	if zimaCron != nil {
		zimaCron.Start()
	}
}

func AddCron(spec string, f func()) {
	if zimaCron == nil {
		zimaCron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		cronSpecs = []CronSpec{}
	}

	cronSpecs = append(cronSpecs, CronSpec{
		Type: "common-cron",
		Func: f,
		Spec: spec,
	})

	zimaCron.AddFunc(spec, f)

}
