# echo-swagger
Create folders for development we will call the app.

mkdir app
mkdir docs
mkdir docs/app

Initialize go mod for go vendor system.

go mod init

Pull the Echo library.

go get -v github.com/labstack/echo/v4

Pull swagger libraries.

go get -v github.com/swaggo/swag/cmd/swag
go get -v github.com/swaggo/echo-swagger

Generate the Swagger Specification.

swag init -g app/main.go --output docs/app

If the operation is successful, you should see 3 new files inside folder docs/echosimple. These files are:

docs.go => Requires to generate SwaggerUI.
swagger.json => The Swagger Specification in json file format.
swagger.yaml => The Swagger Specification in yaml file format.