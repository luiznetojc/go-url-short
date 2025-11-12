package db

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/luizn/go-url-short/internal/config"
)
func ConnectCassandra(cfg *config.Config) *gocql.Session {
	cluster := gocql.NewCluster(cfg.CassandraHost)
	cluster.Keyspace = cfg.CassandraKeyspace
	cluster.Port = cfg.CassandraPort
	cluster.ConnectTimeout = 5 * time.Second
	cluster.ProtoVersion = 4
	if cfg.CassandraUsername != "" && cfg.CassandraPassword != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: cfg.CassandraUsername,
			Password: cfg.CassandraPassword,
		}
	}

	switch cfg.CassandraConsistency {
	case "ONE":
		cluster.Consistency = gocql.One
	case "LOCAL_QUORUM":
		cluster.Consistency = gocql.LocalQuorum
	default:
		cluster.Consistency = gocql.Quorum
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar no Cassandra: %v", err)
	}

	log.Printf("✅ Conectado ao Cassandra em %s:%d (keyspace: %s)", cfg.CassandraHost, cfg.CassandraPort, cfg.CassandraKeyspace)
	return session
}