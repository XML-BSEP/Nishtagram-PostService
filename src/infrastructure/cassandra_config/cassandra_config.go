package cassandra_config

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func init_viper() {
	viper.SetConfigFile(`config/cassandra.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
}

const (
	CreateKeyspace ="CREATE KEYSPACE if not exists post_keyspace WITH replication = { 'class': 'SimpleStrategy', 'replication_factor': '1' };"
)

func NewCassandraSession() (*gocql.Session, error) {
	init_viper()
	domain := viper.GetString(`server.domain`)
	fmt.Println(domain)
	cluster := gocql.NewCluster(viper.GetString(`server.domain`))
	cluster.Port, _ = strconv.Atoi(viper.GetString(`server.port`))
	cluster.ProtoVersion, _ = strconv.Atoi(viper.GetString(`proto_version`))

	cluster.DisableInitialHostLookup = true

	session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = session.Query(CreateKeyspace).Exec()

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return session, err
}

