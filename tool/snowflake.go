package tool

import "github.com/bwmarrin/snowflake"

func SnowflakeId(node int64) (int64, error) {
	newNode, err := snowflake.NewNode(node)
	if err != nil {
		return -1, err
	}

	id := newNode.Generate()
	return id.Int64(), nil
}
