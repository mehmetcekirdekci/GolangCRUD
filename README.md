# GolangCRUD

Unit Test Packages:

*`go install github.com/golang/mock/mockgen@v1.6.0` </br>
*`go get go.mongodb.org/mongo-driver/x/mongo/driver/ocsp@v1.9.1`  </br>
*`go get golang.org/x/sync/errgroup`  </br>
*`go get github.com/golang/mock/mockgen/model` </br>
*`go install github.com/golang/mock/mockgen@v1.6.0` </br>
*`go get github.com/stretchr/testify/assert`

Mockgen generate commant:

`mockgen --destination=./mocks/product/services/mockBaseService.go github.com/mehmetcekirdekci/GolangCRUD/app/product/services BaseService`
