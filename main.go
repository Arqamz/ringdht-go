package main

import (
	"fmt"
	"math/big"
	"github.com/google/btree"
)

type Node struct {
	ID   *big.Int
	Data *btree.BTree
}
