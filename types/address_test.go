package types

import (
	"testing"

	"github.com/youthonline/bitshares/config"
	"github.com/stretchr/testify/assert"
)

var addresses = []string{
	"BTSFN9r6VYzBK8EKtMewfNbfiGCr56pHDBFi",
}

func TestAddress(t *testing.T) {
	config.SetCurrent(config.ChainIDBTS)

	for _, add := range addresses {
		address, err := NewAddressFromString(add)
		if err != nil {
			assert.FailNow(t, err.Error(), "NewAddressFromString")
		}

		assert.Equal(t, add, address.String())
	}
}
