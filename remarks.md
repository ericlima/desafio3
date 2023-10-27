para atualizar o grpc

 ir para a pasta 
 ~/estudos/desafio3/internal/infra/grpc

execute
 protoc --go_out=pb/ --go-grpc_out=pb/ ./protofiles/order.proto


 para testar com o evans:
 cd tests
 evans --proto ./orders.proto repl


 go run github.com/99designs/gqlgen generate