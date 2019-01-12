package main

import (
	"flag"
	"os"

	"github.com/op/go-logging"
	"github.com/xuguruogu/rdb"
)

var (
	log = logging.MustGetLogger("[rdb report]")
)

func init() {
	logging.SetBackend(logging.NewLogBackend(os.Stderr, "", 0))
	logging.SetFormatter(logging.MustStringFormatter(
		`%{color}%{time:2006-01-02 15:04:05.000} %{shortfile} %{shortfunc} [%{level:.4s}] %{color:reset} %{module} %{message}`,
	))
}

func main() {
	var file string
	flag.StringVar(&file, "f", "dump.rdb", "rdb")
	flag.Parse()

	// step 2: decode and push to ckv_plus
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		log.Errorf("open rdb %s failed, err: %+v", file, err)
		return
	}
	defer f.Close()

	report := &RdbReport{}

	if err := rdb.Decode(f, report); err != nil {
		log.Errorf("decode rdb file %s failed, err: %+v", file, err)
		return
	}

	report.Report()

	return
}

// RdbReport ...
type RdbReport struct {
	cnt      uint64
	keyLen   uint64
	valueLen uint64
	mString  memHistogram
	mHash    memHistogram
	cHash    countHistogram
	mList    memHistogram
	cList    countHistogram
	mSet     memHistogram
	cSet     countHistogram
	mZset    memHistogram
	cZset    countHistogram
	vl       uint64
	ll       uint64
}

// StartRDB is called when parsing of a valid RDB file starts.
func (r *RdbReport) StartRDB() error {
	log.Info("StartRDB")
	return nil
}

// StartDatabase is called when database n starts.
// Once a database starts, another database will not start until EndDatabase is called.
func (r *RdbReport) StartDatabase(n int) error {
	return nil
}

// Aux field
func (r *RdbReport) Aux(key, value []byte) error {
	log.Info("Aux", string(key), string(value))
	return nil
}

// ResizeDatabase hint
func (r *RdbReport) ResizeDatabase(dbSize, expiresSize uint32) error {
	return nil
}

// Set is called once for each string key.
func (r *RdbReport) Set(key, value []byte, expiry int64) error {
	r.cnt++
	r.keyLen += uint64(len(key) + 45)
	vl := uint64(len(value))
	r.valueLen += vl
	r.mString.add(vl)
	return nil
}

// StartHash is called at the beginning of a hash.
// Hset will be called exactly length times before EndHash.
func (r *RdbReport) StartHash(key []byte, length, expiry int64) error {
	r.cnt++
	r.keyLen += uint64(len(key) + 45)
	r.vl = 0
	r.ll = 0
	return nil
}

// Hset is called once for each field=value pair in a hash.
func (r *RdbReport) Hset(key, field, value []byte) error {
	r.vl += uint64(len(field) + len(value) + 38)
	r.ll++
	return nil
}

// EndHash is called when there are no more fields in a hash.
func (r *RdbReport) EndHash(key []byte) error {
	r.valueLen += r.vl
	r.mHash.add(r.vl)
	r.cHash.add(r.ll)
	return nil
}

// StartSet is called at the beginning of a set.
// Sadd will be called exactly cardinality times before EndSet.
func (r *RdbReport) StartSet(key []byte, cardinality, expiry int64) error {
	r.cnt++
	r.keyLen += uint64(len(key) + 45)
	r.vl = 0
	r.ll = 0
	return nil
}

// Sadd is called once for each member of a set.
func (r *RdbReport) Sadd(key, member []byte) error {
	r.vl += uint64(len(member) + 38)
	r.ll++
	return nil
}

// EndSet is called when there are no more fields in a set.
func (r *RdbReport) EndSet(key []byte) error {
	r.valueLen += r.vl
	r.mSet.add(r.vl)
	r.cSet.add(r.ll)
	return nil
}

// StartList is called at the beginning of a list.
// Rpush will be called exactly length times before EndList.
// If length of the list is not known, then length is -1
func (r *RdbReport) StartList(key []byte, length, expiry int64) error {
	r.cnt++
	r.keyLen += uint64(len(key) + 45)
	r.vl = 0
	r.ll = 0
	return nil
}

// Rpush is called once for each value in a list.
func (r *RdbReport) Rpush(key, value []byte) error {
	r.vl += uint64(len(value) + 19)
	r.ll++
	return nil
}

// EndList is called when there are no more values in a list.
func (r *RdbReport) EndList(key []byte) error {
	r.valueLen += r.vl
	r.mList.add(r.vl)
	r.cList.add(r.ll)
	return nil
}

// StartZSet is called at the beginning of a sorted set.
// Zadd will be called exactly cardinality times before EndZSet.
func (r *RdbReport) StartZSet(key []byte, cardinality, expiry int64) error {
	r.cnt++
	r.keyLen += uint64(len(key) + 45)
	r.vl = 0
	r.ll = 0
	return nil
}

// Zadd is called once for each member of a sorted set.
func (r *RdbReport) Zadd(key []byte, score float64, member []byte) error {
	r.vl += uint64(len(member) + 69 + 8)
	r.ll++
	return nil
}

// EndZSet is called when there are no more members in a sorted set.
func (r *RdbReport) EndZSet(key []byte) error {
	r.valueLen += r.vl
	r.mZset.add(r.vl)
	r.cZset.add(r.ll)
	return nil
}

// EndDatabase is called at the end of a database.
func (r *RdbReport) EndDatabase(n int) error {
	return nil
}

// EndRDB is called when parsing of the RDB file is complete.
func (r *RdbReport) EndRDB() error {
	log.Info("EndRDB")
	return nil
}
