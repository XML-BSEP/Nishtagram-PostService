package cassandra_config

import (
	"github.com/gocql/gocql"
	logger "github.com/jelena-vlajkov/logger/logger"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func init_viper() {
	viper.SetConfigFile(`src/config/cassandra.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
}

const (
	CreateKeyspace ="CREATE KEYSPACE if not exists post_keyspace WITH replication = { 'class': 'SimpleStrategy', 'replication_factor': '1' };"
)

func NewCassandraSession(logger *logger.Logger) (*gocql.Session, error) {
	init_viper()
	var domain string
	if viper.GetBool(`docker`){
		domain = viper.GetString(`server.domain_docker`) + ":" + viper.GetString(`server.port_docker`)
	}else{
		domain = viper.GetString(`server.domain_localhost`) + ":" + viper.GetString(`server.port_localhost`)
	}
	cluster := gocql.NewCluster(domain)
	cluster.ProtoVersion, _ = strconv.Atoi(viper.GetString(`proto_version`))
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = time.Second * 1000
	//cluster.Keyspace = "post_keyspace"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	cluster.DisableInitialHostLookup = true

	session, err := cluster.CreateSession()
	if err != nil {
		logger.Logger.Fatalf("failed to connect to Cassandra Post DB, error: %v\n", err)
		return nil, err
	}

	err = session.Query(CreateKeyspace).Exec()

	if err != nil {
		logger.Logger.Fatalf("failed to create keyspace in Cassandra Post DB, error: %v\n", err)
		return nil, err
	}
	return session, err
}

