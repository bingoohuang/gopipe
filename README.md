# gopipe

[![Travis CI](https://travis-ci.com/bingoohuang/gopipe.svg?branch=master)](https://travis-ci.com/bingoohuang/gopipe)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/bingoohuang/gopipe/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/bingoohuang/gopipe)
[![Coverage Status](http://codecov.io/github/bingoohuang/gopipe/coverage.svg?branch=master)](http://codecov.io/github/bingoohuang/gopipe?branch=master)
[![goreport](https://www.goreportcard.com/badge/github.com/bingoohuang/gopipe)](https://www.goreportcard.com/report/github.com/bingoohuang/gopipe)

## usage

### define the pipeline in the yaml file.

```yaml
# https://www.jianshu.com/p/3bb437b4edb9
stages:
  - build
  - test
  - deploy

job 1:
  stage: build
  script:
    - mkdir .public
    - cp -r * .public
    - mv .public public

job 2:
  stage: test
  script: make test

job 4:
  stage: deploy
  script: make deploy
```

### load the yaml and run the pipeline

```go
import (
    "github.com/bingoohuang/golog"
    "github.com/bingoohuang/gopipe/pkg/gopipe"
)

func init() {
    golog.SetupLogrus(nil, "")
}

func main() {
    config, err := ioutil.ReadFile("YourPipeline.yaml")
    if err != nil {
        panic(err)
    }
    
    pipe := &gopipe.PipelineConfig{}
    if err := c.Parse(pipe); err != nil {
        panic(err)
    }
    
    pipe :=.Run()
}
```

example output:

```log
2020-07-19 16:46:07.225    INFO 5741 --- [   19] [-]            run.go:25 : --> stage build
2020-07-19 16:46:07.225    INFO 5741 --- [   19] [-]            run.go:38 : --> job "job 1"
2020-07-19 16:46:07.225    INFO 5741 --- [   19] [-]            run.go:48 : --> script "mkdir .public"
2020-07-19 16:46:07.229 WARNING 5741 --- [   19] [-]            run.go:59 : mkdir: .public: File exists
2020-07-19 16:46:07.229    INFO 5741 --- [   19] [-]            run.go:63 : <-- script "mkdir .public"
2020-07-19 16:46:07.229    INFO 5741 --- [   19] [-]            run.go:48 : --> script "cp -r * .public"
2020-07-19 16:46:07.236    INFO 5741 --- [   19] [-]            run.go:63 : <-- script "cp -r * .public"
2020-07-19 16:46:07.236    INFO 5741 --- [   19] [-]            run.go:48 : --> script "mv .public public"
2020-07-19 16:46:07.241    INFO 5741 --- [   19] [-]            run.go:63 : <-- script "mv .public public"
2020-07-19 16:46:07.243    INFO 5741 --- [   19] [-]            run.go:44 : <-- job "job 1"
2020-07-19 16:46:07.243    INFO 5741 --- [   19] [-]            run.go:33 : <-- stage build
2020-07-19 16:46:07.243    INFO 5741 --- [   19] [-]            run.go:25 : --> stage test
2020-07-19 16:46:07.243    INFO 5741 --- [   19] [-]            run.go:38 : --> job "job 2"
2020-07-19 16:46:07.243    INFO 5741 --- [   19] [-]            run.go:48 : --> script "make test"
2020-07-19 16:46:07.258 WARNING 5741 --- [   19] [-]            run.go:59 : make: *** No rule to make target `test'.  Stop.
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:63 : <-- script "make test"
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:44 : <-- job "job 2"
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:33 : <-- stage test
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:25 : --> stage deploy
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:38 : --> job "job 4"
2020-07-19 16:46:07.258    INFO 5741 --- [   19] [-]            run.go:48 : --> script "make deploy"
2020-07-19 16:46:07.275 WARNING 5741 --- [   19] [-]            run.go:59 : make: *** No rule to make target `deploy'.  Stop.
2020-07-19 16:46:07.275    INFO 5741 --- [   19] [-]            run.go:63 : <-- script "make deploy"
2020-07-19 16:46:07.275    INFO 5741 --- [   19] [-]            run.go:44 : <-- job "job 4"
2020-07-19 16:46:07.275    INFO 5741 --- [   19] [-]            run.go:33 : <-- stage deploy
```