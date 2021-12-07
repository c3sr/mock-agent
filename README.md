# Mock Agent

The Mock Agent connects to the message queue and whenever 
any request are sent to it, the mock agent will return with canned responses.


```
docker build . -t mock-agent-dev

docker run -it --rm --name mock-agent -e MQ_USER=guest -e MQ_PASSWORD=guest \
-e MQ_HOST=host.docker.internal -e  MQ_PORT=5672 mock-agent-dev worker
```

## Environment Variables

`MQ_USER` 

`MQ_PASSWORD`

`MQ_HOST`

`MQ_PORT=5672` (default rabbitmq port)

`MOCK_FILE` this is what file it will return on the MQ for all queue requests (defaults to classification.json)
