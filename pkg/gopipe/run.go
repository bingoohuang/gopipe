package gopipe

import (
	"time"

	"github.com/gobars/cmd"
	"github.com/sirupsen/logrus"
)

func (c PipelineConfig) Run() {
	if len(c.Stages) > 0 {
		c.runStages()
	} else {
		c.runJobs()
	}
}

func (c PipelineConfig) runJobs() {
	for _, job := range c.Jobs {
		job.Run()
	}
}

func (c PipelineConfig) runStages() {
	for _, stage := range c.Stages {
		c.runStage(stage)
	}
}

func (c PipelineConfig) runStage(stage string) {
	started := false
	for _, job := range c.Jobs {
		if job.Stage == stage {
			started = true
			logrus.Infof("--> stage %s", stage)
			job.Run()
		}
	}

	if !started {
		logrus.Warnf("No jobs found for stage %s", stage)
	} else {
		logrus.Infof("<-- stage %s", stage)
	}
}

func (c Job) Run() {
	logrus.Infof("--> job %q", c.Name)

	for _, script := range c.Scripts {
		c.runScript(script)
	}

	logrus.Infof("<-- job %q", c.Name)
}

func (c Job) runScript(script string) {
	logrus.Infof("--> script %q", script)

	_, status := cmd.Bash(script, cmd.Timeout(30*time.Second))
	if err := status.Error; err != nil {
		logrus.Warnf("error %v", err)
	} else {
		for _, line := range status.Stdout {
			logrus.Infof("%s", line)
		}

		for _, line := range status.Stderr {
			logrus.Warnf("%s", line)
		}
	}

	logrus.Infof("<-- script %q", script)
}
