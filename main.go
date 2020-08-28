package main

import (
	"context"

	"github.com/aclivo/fast"
	"github.com/aclivo/olap"
	"github.com/aclivo/std"
)

func main() {
	ctx := context.Background()

	fastStorage := fast.NewStorage(0)
	server := std.NewServer(fastStorage, 0)

	server.AddDimension(ctx, olap.Dimension{Name: "Product"})
}
