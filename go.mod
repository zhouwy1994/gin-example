module github.com/zhouwy1994/gin-example

go 1.13

replace (
	github.com/zhouwy1994/gin-example/conf => ./conf
	github.com/zhouwy1994/gin-example/middleware => ./middleware
	github.com/zhouwy1994/gin-example/models => ./models
	github.com/zhouwy1994/gin-example/pkg/setting => ./pkg/setting
)

require (
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0
	github.com/astaxie/beego v1.12.2
	github.com/denisenkom/go-mssqldb v0.0.0-20200620013148-b91950f658ec // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.60.1
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/gorm v0.0.0-20180213101209-6e1387b44c64
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.4.2
	github.com/tebeka/strftime v0.1.5 // indirect
)
