# protocol buffer practice (Golang)

- you should install protoc pacakge before using proto files
    ```
    go get github.com/golang/protobuf
    go get github.com/golang/protobuf/proto
    ```
- install gen-go dependency
    ```
    go install google.golang.org/protobuf/cmd/protoc-gen-go
    ```
- how to compile .proto file
    ```
    protoc --go_out. *.proto
    ```
    -> this command generates new file named 'proto.pb.go' from 'person.proto'
- how to run go file that use protobuf
    ```
    go run main.go person.pb.go
    ```
    person.pb.go is the file made from person.proto, by protoc command.

# reference

    getting started with protocol buffers in go - tutorial
    https://www.youtube.com/watch?v=NoDRq6Twkts


____ 

2022-03-18

외국 유튜브 강의 튜토리얼을 통해 고 언어와 protobuf를 함께 사용하는 법을 간단히 익혔다. 

혼자 독학하려니 너무 헷갈렸는데 위의 영상을 통해 10분만에 명쾌하게 익힐 수 있어 정말 좋았다.