# log

`log` is an out-of-the-box logger, it wraps [zap](https://github.com/uber-go/zap) and makes it simpler to use.

## Quick Start

```go
package main

import (
	"errors"

	"github.com/juan-chan/log"
	"github.com/juan-chan/logfields"
)

func main() {
	log.Info("Hello, World!")
	log.Infof("Hello, %s!", "World")
	log.Named("greet").With(logfields.String("date", "2021-09-08")).Info("Hi, my name's Reed, what's your name?")
	log.Warn("Hello, World!", logfields.String("name", "Joan"))
	log.Error("an error occurred", logfields.Error(errors.New("nil pointer exception")))
}
```

Output:

```log
{"level":"info","time":"2021-09-08 10:13:36.461687044","caller":"log/exported.go:39","msg":"Hello, World!"}
{"level":"info","time":"2021-09-08 10:13:36.461789877","caller":"log/exported.go:148","msg":"Hello, World!"}
{"level":"info","time":"2021-09-08 10:37:26.372523411","name":"greet","caller":"log/logger_test.go:12","msg":"Hi, my name's Reed, what's your name?","date":"2021-09-08"}
{"level":"warn","time":"2021-09-08 10:13:36.461801031","caller":"log/exported.go:46","msg":"Hello, World!","name":"Joan"}
{"level":"error","time":"2021-09-08 10:13:36.461807823","caller":"log/exported.go:53","msg":"Hello error","error":"nil pointer exception","stacktrace":"github.com/juan-chan/log.Error\n\t/home/juan/Workspaces/github/juan-chan/log/exported.go:53\ngithub.com/juan-chan/log.TestLogger\n\t/home/juan/Workspaces/github/juan-chan/log/logger_test.go:14\ntesting.tRunner\n\t/usr/local/go/src/testing/testing.go:1193"}
```

## Sentry

`log` supports [sentry](https://github.com/TheZeroSlave/zapsentry) by default, you just need to set
an environment variable named: `SENTRY_DSN` and `log` will add a hook for it.
