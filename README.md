# go-condbuild-demo
To demonstrate how to utilize conditional compilation in GO

### ex1 -- simple
* reference: please see how GO runtime package abstacts the netpoll interface across Linux & xBSD
  - [netpoll.go](https://golang.org/src/runtime/netpoll.go)
  - [netpoll_epoll.go](https://golang.org/src/runtime/netpoll_epoll.go)
  - [netpoll_kqueue.go](https://golang.org/src/runtime/netpoll_kqueue.go)

### ex2 -- standard
* reference:  please see how GO io
  - [Reader Interface](http://golang.org/pkg/io/#Reader)
  - [Reader Implementation by bytes](https://github.com/golang/go/blob/c264c87335ff4b3111d43f830dbe37eac1509f2e/src/bytes/reader.go#L17)
  - [Reader Implementation by bufio](https://github.com/golang/go/blob/c264c87335ff4b3111d43f830dbe37eac1509f2e/src/bufio/bufio.go#L31)
