package snowflake

import (
	"fmt"
	"sync"

	"github.com/bwmarrin/snowflake"
)

const (
	MaxNodeID = 1024
)

var (
	nodeID int64 = 1
	node   *snowflake.Node

	once sync.Once
)

func Init(nodeID int64) {
	if nodeID >= MaxNodeID {
		panic("node id too big")
	}

	once.Do(func() {
		nodeID = nodeID
		var err error
		node, err = snowflake.NewNode(nodeID)
		if err != nil {
			panic(fmt.Sprintf("err  %v", err))
		}
	})
}

func GetID() int64 {
	if node == nil {
		panic("node nil")
	}
	return node.Generate().Int64()
}
