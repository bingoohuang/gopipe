package gopipe

// PipelineConfig defines the configuration for a pipeline.
type PipelineConfig struct {
	Stages []string
	Jobs   []Job
}

type Job struct {
	Name    string
	Stage   string
	Scripts []string
}
