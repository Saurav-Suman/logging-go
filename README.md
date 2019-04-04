# logging-go


# What is `logging-go`?

`logging-go` supports log levels And you can publish log to slack and rabbitMQ for now.

# _Another_ logging library, Why?

> Logging libraries are like opinions, everyone seems to have one depends upon the need.


## And how is `logging-go` different?

- Slack Integration
- RabbitMQ Integration



# Usage/examples:



Log to Slack or Publish to RabbitMQ using a logging-go instantiated like so:

```golang

loggerConf := logger.EnableLogging(logger.Conf{
		"SlackURL":         "Slack Webhook URL",
		"LoggerTimeFormat": "time.RFC3339",
		"RabbitmqURL":      "amqp://RabbitMQURL/",
		"RabbitmqQueue":    "test",
	})
  
logger.Critical(loggerConf, "slack", "Divide by zero")
logger.Critical(loggerConf, "publish", "Divide by zero")
```



# Contributing

1. Create an issue, describe the bugfix/feature you wish to implement.
2. Fork the repository
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request

