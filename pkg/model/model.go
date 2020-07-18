package model

import (
	"github.com/goccy/go-yaml"
)

// PipelineConfig defines the configuration for a pipeline.
type PipelineConfig struct {
	Stages []string

	Jobs []Job
}

type Job struct {
	Name   string
	Stage  string
	Script []string
}

type RawYAML struct {
	Raw []byte
}

func (r *RawYAML) UnmarshalYAML(raw []byte) error {
	r.Raw = raw
	return nil
}

// PipelineConfig parses the configuration from a YAML string.
func (c *PipelineConfig) Parse(config []byte) error {
	// 3 reasons to use github.com/goccy/go-yaml to handle YAML in Go
	m := map[string]RawYAML{}
	if err := yaml.Unmarshal(config, &m); err != nil {
		return err
	}

	v, ok := m["stages"]
	if ok {
		if err := yaml.Unmarshal(v.Raw, &c.Stages); err != nil {
			return err
		}
	}

	delete(m, "stages")

	c.Jobs = make([]Job, 0, len(m))

	for k, v := range m {
		job := Job{Name: k}

		jobMap := map[string]RawYAML{}

		if err := yaml.Unmarshal(v.Raw, &jobMap); err != nil {
			return err
		}

		job.Stage = string(jobMap["stage"].Raw)

		scriptRaw := jobMap["script"]

		slice, err := TrySlice(scriptRaw)
		if err != nil {
			return err
		}
		job.Script = slice

		c.Jobs = append(c.Jobs, job)
	}

	return nil
}

func TrySlice(raw RawYAML) (v []string, err error) {
	if err := yaml.Unmarshal(raw.Raw, &v); err == nil {
		return v, nil
	}

	return append(v, string(raw.Raw)), nil
}
