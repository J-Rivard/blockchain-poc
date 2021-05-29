package models

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain(originatorPublicKey, originatorPrivateKey string) *BlockChain {
	return &BlockChain{
		blocks: make([]*Block, 0),
	}
}

func (b *BlockChain) Length() int {
	return len(b.blocks)
}

func (b *BlockChain) AddToBlockChain(block *Block) {
	b.blocks = append(b.blocks, block)
}

func (b *BlockChain) IsBlockChainValid() bool {
	for index, _ := range b.blocks {
		if index > 0 && b.blocks[index-1].Hash != b.blocks[index].PreviousHash {
			return false
		}
	}

	return true
}
