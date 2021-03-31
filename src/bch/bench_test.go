package bch_test

import (
	"context"
	"encoding/json"
	"github.com/benka-me/redis-proto-vs-json/src/pb"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"strconv"
	"testing"
)

var (
	ctx   = context.Background()
	tests = []Test{}
	rdb   = &redis.Client{}
)

type Test struct {
	key   string
	value *pb.Message
}

func init() {
	for i := 0; i < 1; i++ {
		tests = append(tests, Test{
			key: strconv.Itoa(i),
			value: &pb.Message{
				String_:               gofakeit.AppName(),
				StringString2:         gofakeit.Username(),
				StringStringString3:   gofakeit.Adverb(),
				String4String4String4: gofakeit.Animal(),
				Int:                   int32(gofakeit.Number(100, 5000)),
			},
		})
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sOmE_sEcUrE_pAsS",
		DB:       0,                  // use default DB
	})
}

func BenchmarkProto(b *testing.B) {
	total := []byte{}
	for i := 0; i < b.N; i++ {

		value, err := proto.Marshal(tests[0].value)
		if err != nil {
			panic(err)
		}
		total = append(total, value...)
		save(tests[0].key+".proto", value)
	}
	for i := 0; i < b.N; i++ {
		get(tests[0].key + ".proto")
	}
	b.Log("proto len: ", len(total))
}

func BenchmarkJson(b *testing.B) {
	total := []byte{}
	for i := 0; i < b.N; i++ {
		value, err := json.Marshal(tests[0].value)
		if err != nil {
			panic(err)
		}
		total = append(total, value...)
		save(tests[0].key+".json", value)
	}
	for i := 0; i < b.N; i++ {
		get(tests[0].key + ".json")
	}
	b.Log("json len: ", len(total))
}

func save(key string, value interface{}) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func get(key string) {
	err := rdb.Get(ctx, key).Err()
	if err != nil {
		panic(err)
	}
}
