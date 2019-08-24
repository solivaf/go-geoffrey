# go-geoffrey
Is a lightweight config server written in Go. It is easy to use and getting started.

### Getting Started
You can clone this repo, build the latest version or run the following command:

```bash
$ go install github.com/solivaf/go-geoffrey
```
and run 
```bash
$ go-geoffrey
```
to initialize new instance of Geoffrey. By default Geoffrey runs on port 9090 but you can configure through
a config file named **application.yml** in any location which you can specify by environment variable **GEOFFREY_CONFIG**.

```bash
$ GEOFFREY_CONFIG=./config/new-application.yml go-geoffrey
```
and the content should be similar as
```yaml
server:
  port: 9090

git:
  url: https://github.com/solivaf/go-maria #default repository
  credential:
    username: solivaf
    password: somedumbpassword
  repositories:
    - name: config-properties #some specific repository
      url: https://github.com/solivaf/config-properties-foo
      credential:
        username: solivaf
        password: somedumbpassword
```

### Running with Docker

You can run Geoffrey with docker and must provide a base configuration file through the environment variable **GEOFFREY_CONFIG**
as you can see in the example below:

```yaml
version: '3'

services:
  go-geoffrey:
    image: fernandosolivas/go-geoffrey:latest
    volumes:
      - ./config/:/app/config
    environment:
      - GEOFFREY_CONFIG=/app/config/
    entrypoint: /go/bin/go-geoffrey
```
and you will have your config server running.

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/solivaf/20f1873a92f0a0c7376f0a92537658a6) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **[Fernando Soliva](https://github.com/solivaf)** - *Initial work*

See also the list of [contributors](https://github.com/solivaf/go-geoffrey/contributors) who participated in this project.