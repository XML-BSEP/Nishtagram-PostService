module post-service

go 1.16

replace github.com/jelena-vlajkov/logger/logger => ../../Nishtagram-Logger/

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/casbin/casbin/v2 v2.31.2
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-resty/resty/v2 v2.6.0
	github.com/gocql/gocql v0.0.0-20210515062232-b7ef815b4556
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.2.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/jelena-vlajkov/logger/logger v1.0.0
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/microcosm-cc/bluemonday v1.0.10
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.29.0 // indirect
	github.com/spf13/viper v1.7.1
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/grpc v1.38.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
