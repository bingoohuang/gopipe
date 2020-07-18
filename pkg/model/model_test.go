package model_test

import (
	"io/ioutil"
	"testing"

	"github.com/bingoohuang/gopipe/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestParsePipelineConfig(t *testing.T) {
	config, err := ioutil.ReadFile("testdata/a.yaml")
	assert.Nil(t, err)

	c := &model.PipelineConfig{}
	assert.Nil(t, c.Parse(config))
	assert.Equal(t, &model.PipelineConfig{
		Stages: []string{"build", "test", "deploy"},
		Jobs: []model.Job{
			{
				Name:  "job 1",
				Stage: "build",
				Script: []string{
					"mkdir .public",
					"cp -r * .public",
					"mv .public public",
				},
			},
			{
				Name:  "job 2",
				Stage: "test",
				Script: []string{
					"make test",
				},
			},
			{
				Name:  "job 4",
				Stage: "deploy",
				Script: []string{
					"make deploy",
				},
			},
		},
	}, c)
}
