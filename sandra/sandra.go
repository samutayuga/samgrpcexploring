package sandra

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

const (
	chost = "192.168.0.171"
	cport = "9042"
	cks   = "system"
	ccons = "LOCAL_QUORUM"

	dropks       = "DROP KEYSPACE %s ;"
	descKs       = "describe keyspaces ;"
	listAllFiles = ";"
)

var (
	//Csess ...
	Csess *gocql.Session

	ccluster *gocql.ClusterConfig
)

//Csessinit ...
func Csessinit(ks string) {

	port := func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9043
		}
		return i

	}

	ccluster = gocql.NewCluster(chost)
	ccluster.Port = port(cport)
	ccluster.Keyspace = ks
	ccluster.Timeout = 10 * time.Second
	ccluster.Consistency = gocql.All
	s, err := ccluster.CreateSession()
	if err != nil {
		log.Printf("ERROR: fail create cassandra session, %s", err.Error())
		os.Exit(1)
	}
	log.Printf("Cassandra session is created connected to keyspace %s", ks)

	Csess = s
	//ShowTables()
}

//Csessclose ...
func Csessclose() {
	Csess.Close()
	log.Println("Disconnected from cassandra")
}

//CdelKeyspace ...
func CdelKeyspace(ksName string) error {
	stDel := fmt.Sprintf(dropks, ksName)
	if err := Csess.Query(stDel).Exec(); err != nil {
		log.Printf("ERROR: fail to drop keyspace %s, %v", ksName, err.Error())
		return err
	}

	return nil
}
