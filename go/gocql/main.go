package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
	log "github.com/sirupsen/logrus"
)

func main() {
	nbIDs := flag.Int("n", 1000, "Number of ids to write & read")
	flag.Parse()

	log.Infof("Will write then read %d values", *nbIDs)

	cluster := gocql.NewCluster("localhost:9042")
	cluster.Timeout = 5 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("could not create the CQL session: %v", err)
	}
	defer session.Close()

	ctx := context.Background()

	ids := generateIDs(*nbIDs)

	if err := write(ctx, session, ids); err != nil {
		log.Fatal(err)
	}

	values, err := read(ctx, session, ids)
	if err != nil {
		log.Fatal(err)
	}
	if n := len(values); n != *nbIDs {
		log.Fatalf("only got %d rows instead of %d", n, *nbIDs)
	}
	log.Info("All good!")
}

func generateIDs(n int) []uint32 {
	res := []uint32{}
	for i := 0; i < n; i++ {
		res = append(res, uint32(i))
	}
	return res
}

func write(ctx context.Context, session *gocql.Session, ids []uint32) error {
	statement, _ := qb.Insert("toto.main").Columns("id", "value").ToCql()
	for _, id := range ids {
		query := session.Query(statement, id, rand.Float64())
		if err := query.Exec(); err != nil {
			return fmt.Errorf("failed to write a value for id %d: %w", id, err)
		}
	}
	return nil
}

func read(ctx context.Context, session *gocql.Session, ids []uint32) (map[uint32]float64, error) {
	statement, _ := qb.Select("toto.main").Columns("id", "value").Where(qb.In("id")).ToCql()
	query := session.Query(statement, ids)
	defer query.Release()

	if err := query.Exec(); err != nil {
		return nil, fmt.Errorf("failed to exec the query: %w", err)
	}

	iter := query.Iter()
	log.Infof("numRows: %d", iter.NumRows())

	values, err := scan(iter, ids)
	if err != nil {
		return nil, err
	}

	for _, warning := range iter.Warnings() {
		log.Warnf("received warning from gocql client: %s", warning)
	}
	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("failed to close the CQL row iterator: %w", err)
	}
	return values, nil
}

func scan(iter *gocql.Iter, ids []uint32) (map[uint32]float64, error) {
	values := make(map[uint32]float64, len(ids))
	var (
		id uint32
		v  float64
	)
	for iter.Scan(&id, &v) {
		values[id] = v
	}
	return values, nil
}

func badScan(iter *gocql.Iter, ids []uint32) (map[uint32]float64, error) {
	values := make(map[uint32]float64, len(ids))
	for _, id := range ids {
		var v float64
		if !iter.Scan(&v) {
			return nil, fmt.Errorf("failed to read value for id %d", id)
		}
		values[id] = v
	}
	return values, nil
}
