package graph

import (
	"encoding/base64"
	"fmt"
	"strings"

	redis "github.com/go-redis/redis/v7"
)

// RedisGraph struct.
// Implementation of the graph.Graph interface using
// a redis backbone.
type RedisGraph struct {
	prefix string
	id     string
	client *redis.Client
}

// NewRedisGraph func
func NewRedisGraph(id string, encodeID bool) *RedisGraph {
	client := redis.NewClient(&redis.Options{
		Addr:     "docker_redis_1:6379",
		Password: "",
		DB:       0,
	})
	if encodeID {
		id = base64.StdEncoding.EncodeToString([]byte(id))
	}
	return &RedisGraph{
		prefix: "raiding-raccoon",
		id:     id,
		client: client,
	}
}

// Exists func
func (graph *RedisGraph) Exists(id string, encodeID bool) bool {
	if encodeID {
		id = base64.StdEncoding.EncodeToString([]byte(id))
	}
	keys, _, _ := graph.client.Scan(0, fmt.Sprintf("%v:%v:*", graph.prefix, id), 1).Result()
	return len(keys) != 0
}

func (graph *RedisGraph) makeID(isNode bool, id ...string) string {
	cat := "edge"
	if isNode {
		cat = "node"
	}
	base := fmt.Sprintf("%v:%v:%v", graph.prefix, graph.id, cat)
	for _, idE := range id {
		b64id := base64.StdEncoding.EncodeToString([]byte(idE))
		base = fmt.Sprintf("%v:%v", base, b64id)
	}
	return base
}

// AddNode func
func (graph *RedisGraph) AddNode(id string) bool {
	return graph.client.SetNX(graph.makeID(true, id), 0, 0).Val()
}

// AddEdge func
func (graph *RedisGraph) AddEdge(source string, dest string) bool {
	return graph.client.SetNX(graph.makeID(false, source, dest), 0, 0).Val()
}

// Nodes func
func (graph *RedisGraph) Nodes(out chan<- string) {
	keys := []string{}
	cursor := uint64(0)
	for {
		keys, cursor, _ = graph.client.Scan(cursor, fmt.Sprintf("%v:%v:node:*", graph.prefix, graph.id), 10).Result()
		for _, k := range keys {
			val := strings.Split(k, ":")[3]
			valDec, _ := base64.StdEncoding.DecodeString(val)
			out <- string(valDec)
		}
		if cursor == 0 {
			break
		}
	}
	defer close(out)
}

// Edges func
func (graph *RedisGraph) Edges(out chan<- Edge) {
	keys := []string{}
	cursor := uint64(0)
	for {
		keys, cursor, _ = graph.client.Scan(cursor, fmt.Sprintf("%v:%v:edge:*", graph.prefix, graph.id), 10).Result()
		for _, k := range keys {
			val1 := strings.Split(k, ":")[3]
			valDec1, _ := base64.StdEncoding.DecodeString(val1)
			val2 := strings.Split(k, ":")[4]
			valDec2, _ := base64.StdEncoding.DecodeString(val2)
			out <- Edge{
				Source: string(valDec1),
				Dest:   string(valDec2),
			}
		}
		if cursor == 0 {
			break
		}
	}
	defer close(out)
}

// IterateCb func
func (graph *RedisGraph) IterateCb(cb func(string, string)) {
	edges := make(chan Edge)
	go graph.Edges(edges)
	for e := range edges {
		cb(e.Source, e.Dest)
	}
}
