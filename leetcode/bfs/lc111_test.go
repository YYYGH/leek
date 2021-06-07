package bfs

import (
	al "leek/algorithm"
	"testing"
)

func TestMinDepth(t *testing.T) {
	trees := map[int][]interface{}{
		2: {3, 9, 20, nil, nil, 15, 7},
	}

	for d, tree := range trees {
		root := al.CreateBinTree(tree)
		// info := al.LevelOrder(root)
		depth := MinDepth(root)
		if depth != d {
			t.Errorf("mindepth(%v)= %d; expected %d", root, depth, d)
		}
	}
}
