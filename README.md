# gologin
create by gin and gorm 

1. inits - contains all the initializer files 
2. controllers - contains controller files that will handle route functions 
3. middlewares – contains middleware files 
4. migrations - contains migration files 
5. models - contains model files

### Install libraries
In this section, you will install all the frameworks and libraries necessary for your project.

Run the following command to install the CompileDaemon package for automatic builds:

go get github.com/githubnemo/CompileDaemon
You’ll then need to install it using the go install command like so:

go install github.com/githubnemo/CompileDaemon
Once that's done, run the following command to install the godotenv package for securing your application's secrets:

go get github.com/joho/godotenv
Next, install the Gin framework with the following command:

go get -u github.com/gin-gonic/gin
Once Gin is installed, run the following command to install Gorm:

go get -u gorm.io/gorm
You'll need a database driver to work with Gorm. Run the following command to install a database driver that will be used to connect with your database:

go get -u gorm.io/driver/mysql
The command above will install a database driver for mysql, the database system that XAMPP provides. Visit the XAMPP Downloads Page to install XAMPP for your machine.

Note: You can use and install drivers for other database systems such as PostgreSQL and MySQL. Visit Gorm's Connecting to the Database page for specific instructions.

Once Gorm and the database driver are installed, run the following command to install the bcrypt package, which will be used to hash passwords:

go get -u golang.org/x/crypto/bcrypt
Next, run the following command to install the jwt-go package, which will be used to generate tokens and authenticate users:

go get -u github.com/golang-jwt/jwt/v5
Once that's done, create a main.go file in the root of your project and add the following code inside it:

package main

import "fmt"

func main() {
 fmt.Println("Hello, World!")
}
Next, you'll need to run the ComplileDaemon command so that the project builds automatically every time you save a file:

```
CompileDaemon -command="./gologin"
```

### build

区分环境

```shell
ENV=prod go build main.go
```

### doc

```shell
go install github.com/swaggo/swag/cmd/swag@latest

# 每次修改文档都需要执行
swag init -g router/index.go
# or 解决依赖问题
swag init -g router/index.go --parseDependency --parseInternal

# 访问 http://localhost:5000/swagger/index.html
```

### 部署

```shell
docker compose up

# rebuild
docker-compose up --build
```

### TODO

* token redis
* 参数校验
* 文档
* 日志 log
* 封装返回