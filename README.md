# Geoffrey
A lightweight config server written in Go based on Spring Cloud Config Server. 

## Getting Started
You can clone this repo, build the latest version or run the following command:

```bash
$ go install github.com/solivaf/go-geoffrey #requires go 1.12+

$ GEOFFREY_CONFIG=./config/application.yml go-geoffrey
```

Geoffrey depends on a yaml configuration file with repository configuration which will be used to download
your configuration files. 
```bash
$ GEOFFREY_CONFIG=./config/new-application.yml go-geoffrey
```

By default, Geoffrey needs one repository to run and provide configurations through an http request. You can use a default
repository by using the following configuration file.

```yaml
#new-application.yml with single configuration repository 
server:
  port: 9090

git:
  url: https://github.com/solivaf/go-maria
  credential:
    username: solivaf
    password: somepassword
```

Or multiple configuration repositories. 

```yaml
#new-application.yml with single configuration repository
server:
  port: 9090

git:
  url: https://github.com/solivaf/go-maria
  credential:
    username: solivaf
    password: somepassword
  repositories:
    - name: config-properties #some specific repository
      url: https://github.com/solivaf/config-properties-foo
      credential:
        username: solivaf
        password: somepassword
    - name: second-config-properties #another specific repository
          url: https://github.com/solivaf/config-properties-bar
          credential:
            username: solivaf
            password: somepassword
```

### Running with Docker

You must provide a base configuration file through the environment variable **GEOFFREY_CONFIG** as you can see in the 
example below:

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

## Usage
You can run Geoffrey on a docker container with the following command:

```bash
$ docker run -it --name geoffrey -v <absolute-config-path>:/app/config -e GEOFFREY_CONFIG=/app/config/ fernandosolivas/go-geoffrey:latest /go/bin/go-geoffrey
```

Now you will have a server running and you can get your configurations through an http request as:

```bash
$ curl http://localhost:8080/message/dev
#response
bar:
  foo: testPropertiesYml
```

In this case, geoffrey will search a file named ***message-dev.yml*** inside the repositories specified inside the application.yml
in GEOFFREY_CONFIG path.

## Comparison

Comparing Geoffrey with Spring Cloud Config Server we see a big difference in resource allocation. Just running docker
containers without any particular configuration we see something like the image below

![vs spring cloud](images/stats.png)

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/solivaf/20f1873a92f0a0c7376f0a92537658a6) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **[Fernando Soliva](https://github.com/solivaf)** - *Initial work*

See also the list of [contributors](https://github.com/solivaf/go-geoffrey/contributors) who participated in this project.