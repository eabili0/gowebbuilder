package main

import (
	_ "github.com/go-errors/errors"
	_ "github.com/gorilla/mux"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	_ "github.com/spf13/viper"
	_ "github.com/labbsr0x/goh/gohclient"
	_ "github.com/labbsr0x/goh/gohserver"
	_ "github.com/labbsr0x/goh/gohtypes"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/kataras/iris"
	_ "github.com/kataras/golog"
	_ "github.com/kataras/pio"
	_ "github.com/valyala/fasthttp"
	_ "github.com/klauspost/compress"
	_ "github.com/klauspost/cpuid"
	_ "github.com/valyala/bytebufferpool"
	_ "github.com/graphql-go/graphql"
	_ "github.com/graphql-go/handler"
	// _ "github.com/iris-contrib/middleware/cors"
	_ "github.com/google/uuid"
	_ "github.com/gin-gonic/gin"
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/prometheus/client_golang/prometheus"
	
	_ "gorm.io/gorm"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"

	_ "golang.org/x/oauth2"
	_ "golang.org/x/oauth2/clientcredentials"
)

func main() {
}
