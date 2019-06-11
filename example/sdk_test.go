package smoke

import (
	"testing"

	"github.com/dappledger/AnnChain-go-sdk"
)

const (
	accPriv = "48deaa73f328f38d5fcb29d076b2b639c8491f97d245fc22e95a86366687903a"
	accAddr = "28112ca022224ae7757bcd559666be5340ff109a"
)

var client *sdk.GoSDK

func init() {
	client = sdk.New(accPriv, "127.0.0.1:46657", sdk.ZaCryptoType)
}

func TestPutGet(t *testing.T) {

	hash, err := client.Put([]byte("myname"), sdk.TypeSyn)

	t.Log(hash, err)

	value, err := client.Get(hash)

	t.Log(value, err)
}
