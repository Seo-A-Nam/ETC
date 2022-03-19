# gRPC practice (Go)
create grpc server and client that serves chat service program

- install grpc packages

```
go get -u google.golang.org/grpc
```

- compile chat.proto into chat directory

```
mkdir chat
protoc --go_out=plugins=grpc:chat chat.proto
```

- execute grpc server and client
    - prepare 2 termial and type each command in each terminal.  
```
go run server.go
go run client.go
```


# reference
- youtbe tutorial

https://www.youtube.com/watch?v=BdzYdN_Zd9Q

____

go를 이용한 protobuf 기초 실습을 하고 난 뒤, gRPC도 간단히 실습해본다.