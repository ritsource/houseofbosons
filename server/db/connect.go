package db

import (
	"crypto/tls"
	"flag"
	"net"
	"os"

	"gopkg.in/mgo.v2"
)

// GetDBName .
func GetDBName() string {
	if flag.Lookup("test.v") == nil {
		return os.Getenv("DB_NAME")
	}
	return os.Getenv("TEST_DB_NAME")
}

// Credenticals
var (
	MongoURI             = os.Getenv("MONGO_URI")
	DBName               = GetDBName()
	MongoClusterURI      = os.Getenv("MONGO_CLUSTER_URI")
	MongoClusterDBName   = os.Getenv("MONGO_CLUSTER_DB_NAME")
	MongoClusterUser     = os.Getenv("MONGO_CLUSTER_USER")
	MongoClusterPassword = os.Getenv("MONGO_CLUSTER_PASSWORD")
)

// DBConnect .
func DBConnect() (*mgo.Session, error) {
	if MongoURI != "" {
		tlsConfig := &tls.Config{}

		dialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoClusterURI, MongoClusterURI, MongoClusterURI},
			Database: MongoClusterDBName,
			Username: MongoClusterUser,
			Password: MongoClusterPassword,
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		return mgo.DialWithInfo(dialInfo)
	}

	return DBConnect()
}
