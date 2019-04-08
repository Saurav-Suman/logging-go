# What is `logging-go`?

`logging-go` supports log levels And you can publish log to rabbitMQ for now.

# _Another_ logging library, Why?

> Logging libraries are like opinions, everyone needs one depends upon the need.


## And how is `logging-go` different?

- RabbitMQ Integration



# Usage/examples:



Publish to RabbitMQ using a logging-go instantiated like so:

```golang

import (
	logger "github.com/Saurav-Suman/logging-go"
)

loggerConf := logger.EnableLogging(logger.Conf{
		"LoggerTimeFormat": "time.RFC3339",
		"RabbitmqURL":      "amqp://RabbitMQURL/",
		"RabbitmqQueue":    "test",
	})
  
logger.Critical(loggerConf, "publish", "Divide by zero")
```



# Contributing

1. Create an issue, describe the bugfix/feature you wish to implement.
2. Fork the repository
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request

