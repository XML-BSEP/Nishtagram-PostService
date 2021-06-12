module post-service

go 1.16

replace github.com/jelena-vlajkov/logger/logger => ../../Nishtagram-Logger/

require (
	github.com/casbin/casbin/v2 v2.31.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-resty/resty/v2 v2.6.0
	github.com/gocql/gocql v0.0.0-20210515062232-b7ef815b4556
	github.com/google/uuid v1.2.0
	github.com/jelena-vlajkov/logger/logger v1.0.0
	github.com/microcosm-cc/bluemonday v1.0.10
	github.com/spf13/viper v1.7.1
)
