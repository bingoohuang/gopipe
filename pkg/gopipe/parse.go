package gopipe

import (
	"context"
	"sort"

	"github.com/bingoohuang/go-yaml"
)

type RawMessage struct {
	Raw []byte
	Seq int
}

type contextKey string

const counterKey contextKey = "counter"

func (r *RawMessage) UnmarshalYAML(ctx context.Context, raw []byte) error {
	if ctx != nil {
		counter := ctx.Value(counterKey).(*int)
		*counter++
		r.Seq = *counter
	}

	r.Raw = raw
	return nil
}

// PipelineConfig parses the configuration from a YAML string.
func (c *PipelineConfig) Parse(config []byte) error {
	root, err := c.parseRoot(config)
	if err != nil {
		return err
	}

	if err := c.parseStages(root); err != nil {
		return err
	}

	jobNames := SortKeysByValueSequence(root)
	c.Jobs = make([]Job, 0, len(jobNames))

	for _, jobName := range jobNames {
		job, err := c.parseJob(jobName, root)
		if err != nil {
			return err
		}

		c.Jobs = append(c.Jobs, job)
	}

	return nil
}

func (c *PipelineConfig) parseRoot(config []byte) (map[string]RawMessage, error) {
	seq := 0
	m := map[string]RawMessage{}
	ctx := context.WithValue(context.Background(), counterKey, &seq)
	err := yaml.UnmarshalWithOptions(config, &m, yaml.WithContext(ctx))
	return m, err
}

func (c *PipelineConfig) parseJob(jobName string, m map[string]RawMessage) (Job, error) {
	job := Job{
		Name: jobName,
	}

	v := m[jobName]
	jobMap := map[string]RawMessage{}
	if err := yaml.Unmarshal(v.Raw, &jobMap); err != nil {
		return job, err
	}

	if scriptRaw, ok := jobMap["script"]; ok {
		slice, err := TrySlice(scriptRaw)
		if err != nil {
			return job, err
		}

		job.Scripts = slice
	}

	if stage, ok := jobMap["stage"]; ok {
		job.Stage = string(stage.Raw)
	}

	return job, nil
}

func (c *PipelineConfig) parseStages(m map[string]RawMessage) error {
	if v, ok := m["stages"]; ok {
		if err := yaml.Unmarshal(v.Raw, &c.Stages); err != nil {
			return err
		}
	}

	delete(m, "stages")
	return nil
}

func SortKeysByValueSequence(m map[string]RawMessage) []string {
	seqMap := map[int]string{}
	seqs := make([]int, 0, len(m))
	for k, v := range m {
		seq := v.Seq
		seqMap[seq] = k
		seqs = append(seqs, seq)
	}

	sort.Ints(seqs)

	keys := make([]string, len(seqs))
	for i, seq := range seqs {
		keys[i] = seqMap[seq]
	}

	return keys
}

func TrySlice(raw RawMessage) (v []string, err error) {
	if err := yaml.Unmarshal(raw.Raw, &v); err == nil {
		return v, nil
	}

	return append(v, string(raw.Raw)), nil
}
