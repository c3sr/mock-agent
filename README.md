# Mock Agent

The Mock Agent connects to the message queue and whenever 
any request are sent to it, the mock agent will return with canned responses.


```
docker build . -t mock-agent-dev

docker run -it --rm --name mock-agent -e MQ_USER=guest -e MQ_PASSWORD=guest \
-e MQ_HOST=host.docker.internal -e  MQ_PORT=5672 mock-agent-dev worker
```

## Environment Variables

`MQ_HOST` the hostname of the message queue server

`MQ_PORT` the port number the message queue server listens on (normally 5672)

`MQ_USER` a username that can login to the message queue server

`MQ_PASSWORD` the password that corresponds to the username

`MOCK_FILE` this is what file it will return on the MQ for all queue requests (defaults to classification.json)

`MQ_RESPONSE_CHANNEL` overrides where the agent sends responses to jobs
