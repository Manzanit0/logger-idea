# Random logger idea

With Go 1.21, the `slog` package was released, which got me thinking...
wouldn't it be great if people could use that and it just "work" in a company's
platform?

That triggered the questions: what tend to be some of the hard requirements for
log messages to play nicely in a platform? Things the SDK team generally takes
care of? One of the key requirements is most possibly the addition of certain
structured fields which allow the pipeline to correlate logs to services. This
means making log messages look somewhat like (notice the `service` field):

```json
{"time":"2023-09-07T10:20:09.115191+02:00","level":"INFO","msg":"test log","service":"app"}
```

That's where this POC comes into play: by leveraging the Go build linker... we
could sort this out! Just make sure that the build pipeline always runs
something along the lines of:

```sh
go build -v -ldflags="-X github.com/manzanit0/logger-idea/pkg/logger.Service=app" ./cmd/app
```

And then as long as the service imports the package like below, you're golden:

```go
import (
	_ "github.com/manzanit0/logger-idea/pkg/logger"
)
```

