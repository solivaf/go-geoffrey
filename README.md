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
to initialize new instance of geoffrey. By default geoffrey runs on port 9090 but you can configure through
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