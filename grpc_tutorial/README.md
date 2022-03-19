I tried youtube tutorial to build simple gRPC API. I can highly recommend those videos for anyone who want to learn grpc.

## Youtube tutorial
- Building an Basic API with gRPC and protobuf
    - Server performs add or multiply task when Client sends two number.

https://www.youtube.com/watch?v=Y92WWaZJI24

# How to use
- download all the pacakges needed

    -> type this command under directory grpc_tutorial/
        ```
        go mod tidy
        ```
- prepare 2 terminals
- execute it right under directory grpc_tutorial/
1. on first terminal

```
go run server/main.go
```

2. on second terminal
```
go run client/main.go
```

3. Open your web browser and type these on URL

- you can put any random numbers in that num1, num2 position

- add
    ```
    localhost:8080/add/num1/num2
    ```
- multiply
    ```
    localhost:8080/mult/num1/num2
    ```