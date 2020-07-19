package gopipe_test

import (
	"io/ioutil"
	"testing"

	"github.com/bingoohuang/golog"
	"github.com/bingoohuang/gopipe/pkg/gopipe"

	"github.com/stretchr/testify/assert"
)

func TestParsePipelineConfig(t *testing.T) {
	config, err := ioutil.ReadFile("testdata/a.yaml")
	assert.Nil(t, err)

	c := &gopipe.PipelineConfig{}
	assert.Nil(t, c.Parse(config))
	assert.Equal(t, &gopipe.PipelineConfig{
		Stages: []string{"build", "test", "deploy"},
		Jobs: []gopipe.Job{
			{
				Name:  "job 1",
				Stage: "build",
				Scripts: []string{
					"mkdir .public",
					"cp -r * .public",
					"mv .public public",
				},
			},
			{
				Name:  "job 2",
				Stage: "test",
				Scripts: []string{
					"make test",
				},
			},
			{
				Name:  "job 4",
				Stage: "deploy",
				Scripts: []string{
					"make deploy",
				},
			},
		},
	}, c)

	golog.SetupLogrus(nil, "")
	c.Run()
}
