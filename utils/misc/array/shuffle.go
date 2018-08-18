package array

import (
	"github.com/cuixin/mt"
)

func ShuffleString(mt *mt.MT19937_64, data []string) {
	size := len(data)
	for i := size - 1; i > 0; i-- {
		if j := int(mt.IntN(uint64(i + 1))); i != j {
			temp := data[i]
			data[i], data[j] = data[j], temp
		}
	}
}

func ShuffleInt(mt *mt.MT19937_64, data []int) {
	size := len(data)
	for i := size - 1; i > 0; i-- {
		if j := int(mt.IntN(uint64(i + 1))); i != j {
			temp := data[i]
			data[i], data[j] = data[j], temp
		}
	}
}

func ShuffleInt32(mt *mt.MT19937_64, data []int32) {
	size := len(data)
	for i := size - 1; i > 0; i-- {
		if j := int(mt.IntN(uint64(i + 1))); i != j {
			temp := data[i]
			data[i], data[j] = data[j], temp
		}
	}
}

func ShuffleUint32(mt *mt.MT19937_64, data []uint32) {
	size := len(data)
	for i := size - 1; i > 0; i-- {
		if j := int(mt.IntN(uint64(i + 1))); i != j {
			temp := data[i]
			data[i], data[j] = data[j], temp
		}
	}
}
