package stackqueue

func isProperlyNested(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var expecteds []string
	for _, r := range s {
		var right string
		if right, ok := parens[string(r)]; ok {
			expecteds = append(expecteds, right)
			continue
		}

		if len(expecteds) == 0 {
			return false
		}
		right = string(r)
		expected := expecteds[len(expecteds)-1]
		expecteds = expecteds[:len(expecteds)-1]
		if right != expected {
			return false
		}
	}

	return len(expecteds) == 0
}

var parens = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}

func countFishesAlive(fishes []*fish) int {
	var ups, downs fishStack
	for _, fish := range fishes {
		if fish.direction == downstream {
			downs.push(fish)
			continue
		}

		for 0 < len(downs) {
			down := downs.peek()
			if down.canEat(fish) {
				break
			}

			downs.pop()
		}
		if len(downs) == 0 {
			ups.push(fish)
		}
	}

	return len(append(ups, downs...))
}

type fishStack []*fish

func (fs *fishStack) push(fish *fish) {
	*fs = append(*fs, fish)
}

func (fs *fishStack) pop() *fish {
	peeked := fs.peek()
	*fs = (*fs)[:len(*fs)-1]

	return peeked
}

func (fs *fishStack) peek() *fish {
	return (*fs)[len(*fs)-1]
}

type fish struct {
	position  int
	size      int
	direction direction
}

const (
	_ direction = iota
	upstream
	downstream
)

type direction int

func (f *fish) canEat(target *fish) bool {
	return target.size < f.size
}

func countBlocksForStoneWall(blocks []*block) int {
	var useds blockStack
	var cnt int
	for _, block := range blocks {
		for 0 < len(useds) && block.height < useds.peek().height {
			useds.pop()
		}

		if len(useds) == 0 || useds.peek().height < block.height {
			useds.push(block)
			cnt++
		}
	}

	return cnt
}

type blockStack []*block

func (bs *blockStack) push(block *block) {
	*bs = append(*bs, block)
}

func (bs *blockStack) pop() *block {
	peeked := bs.peek()
	*bs = (*bs)[:len(*bs)-1]

	return peeked
}

func (bs *blockStack) peek() *block {
	return (*bs)[len(*bs)-1]
}

type block struct {
	height int
}
