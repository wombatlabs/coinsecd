package consensushashing_test

import (
	"encoding/hex"
	"fmt"
	"github.com/wombatlabs/coinsecd/domain/consensus/utils/subnetworks"
	"testing"

	"github.com/kaspanet/go-secp256k1"

	"github.com/wombatlabs/coinsecd/domain/consensus/utils/consensushashing"
	"github.com/wombatlabs/coinsecd/domain/consensus/utils/txscript"
	"github.com/wombatlabs/coinsecd/domain/consensus/utils/utxo"
	"github.com/wombatlabs/coinsecd/domain/dagconfig"
	"github.com/wombatlabs/coinsecd/util"

	"github.com/wombatlabs/coinsecd/domain/consensus/model/externalapi"
)

// shortened versions of SigHash types to fit in single line of test case
const (
	all                = consensushashing.SigHashAll
	none               = consensushashing.SigHashNone
	single             = consensushashing.SigHashSingle
	allAnyoneCanPay    = consensushashing.SigHashAll | consensushashing.SigHashAnyOneCanPay
	noneAnyoneCanPay   = consensushashing.SigHashNone | consensushashing.SigHashAnyOneCanPay
	singleAnyoneCanPay = consensushashing.SigHashSingle | consensushashing.SigHashAnyOneCanPay
)

func modifyOutput(outputIndex int) func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	return func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
		clone := tx.Clone()
		clone.Outputs[outputIndex].Value = 100
		return clone
	}
}

func modifyInput(inputIndex int) func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	return func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
		clone := tx.Clone()
		clone.Inputs[inputIndex].PreviousOutpoint.Index = 2
		return clone
	}
}

func modifyAmountSpent(inputIndex int) func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	return func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
		clone := tx.Clone()
		utxoEntry := clone.Inputs[inputIndex].UTXOEntry
		clone.Inputs[inputIndex].UTXOEntry = utxo.NewUTXOEntry(666, utxoEntry.ScriptPublicKey(), false, 100)
		return clone
	}
}

func modifyScriptPublicKey(inputIndex int) func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	return func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
		clone := tx.Clone()
		utxoEntry := clone.Inputs[inputIndex].UTXOEntry
		scriptPublicKey := utxoEntry.ScriptPublicKey()
		scriptPublicKey.Script = append(scriptPublicKey.Script, 1, 2, 3)
		clone.Inputs[inputIndex].UTXOEntry = utxo.NewUTXOEntry(utxoEntry.Amount(), scriptPublicKey, false, 100)
		return clone
	}
}

func modifySequence(inputIndex int) func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	return func(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
		clone := tx.Clone()
		clone.Inputs[inputIndex].Sequence = 12345
		return clone
	}
}

func modifyPayload(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	clone := tx.Clone()
	clone.Payload = []byte{6, 6, 6, 4, 2, 0, 1, 3, 3, 7}
	return clone
}

func modifyGas(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	clone := tx.Clone()
	clone.Gas = 1234
	return clone
}

func modifySubnetworkID(tx *externalapi.DomainTransaction) *externalapi.DomainTransaction {
	clone := tx.Clone()
	clone.SubnetworkID = externalapi.DomainSubnetworkID{6, 6, 6, 4, 2, 0, 1, 3, 3, 7}
	return clone
}

func TestCalculateSignatureHashSchnorr(t *testing.T) {
	nativeTx, subnetworkTx, err := generateTxs()
	if err != nil {
		t.Fatalf("Error from generateTxs: %+v", err)
	}

	// Note: Expected values were generated by the same code that they test,
	// As long as those were not verified using 3rd-party code they only check for regression, not correctness
	tests := []struct {
		name                  string
		tx                    *externalapi.DomainTransaction
		hashType              consensushashing.SigHashType
		inputIndex            int
		modificationFunction  func(*externalapi.DomainTransaction) *externalapi.DomainTransaction
		expectedSignatureHash string
	}{
		// native transactions

		// sigHashAll
		{name: "native-all-0", tx: nativeTx, hashType: all, inputIndex: 0,
			expectedSignatureHash: "03b7ac6927b2b67100734c3cc313ff8c2e8b3ce3e746d46dd660b706a916b1f5"},
		{name: "native-all-0-modify-input-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyInput(1), // should change the hash
			expectedSignatureHash: "a9f563d86c0ef19ec2e4f483901d202e90150580b6123c3d492e26e7965f488c"},
		{name: "native-all-0-modify-output-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // should change the hash
			expectedSignatureHash: "aad2b61bd2405dfcf7294fc2be85f325694f02dda22d0af30381cb50d8295e0a"},
		{name: "native-all-0-modify-sequence-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifySequence(1), // should change the hash
			expectedSignatureHash: "0818bd0a3703638d4f01014c92cf866a8903cab36df2fa2506dc0d06b94295e8"},
		{name: "native-all-anyonecanpay-0", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "24821e466e53ff8e5fa93257cb17bb06131a48be4ef282e87f59d2bdc9afebc2"},
		{name: "native-all-anyonecanpay-0-modify-input-0", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyInput(0), // should change the hash
			expectedSignatureHash: "d09cb639f335ee69ac71f2ad43fd9e59052d38a7d0638de4cf989346588a7c38"},
		{name: "native-all-anyonecanpay-0-modify-input-1", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyInput(1), // shouldn't change the hash
			expectedSignatureHash: "24821e466e53ff8e5fa93257cb17bb06131a48be4ef282e87f59d2bdc9afebc2"},
		{name: "native-all-anyonecanpay-0-modify-sequence", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "24821e466e53ff8e5fa93257cb17bb06131a48be4ef282e87f59d2bdc9afebc2"},

		// sigHashNone
		{name: "native-none-0", tx: nativeTx, hashType: none, inputIndex: 0,
			expectedSignatureHash: "38ce4bc93cf9116d2e377b33ff8449c665b7b5e2f2e65303c543b9afdaa4bbba"},
		{name: "native-none-0-modify-output-1", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "38ce4bc93cf9116d2e377b33ff8449c665b7b5e2f2e65303c543b9afdaa4bbba"},
		{name: "native-none-0-modify-sequence-0", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifySequence(0), // should change the hash
			expectedSignatureHash: "d9efdd5edaa0d3fd0133ee3ab731d8c20e0a1b9f3c0581601ae2075db1109268"},
		{name: "native-none-0-modify-sequence-1", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "38ce4bc93cf9116d2e377b33ff8449c665b7b5e2f2e65303c543b9afdaa4bbba"},
		{name: "native-none-anyonecanpay-0", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "06aa9f4239491e07bb2b6bda6b0657b921aeae51e193d2c5bf9e81439cfeafa0"},
		{name: "native-none-anyonecanpay-0-modify-amount-spent", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyAmountSpent(0), // should change the hash
			expectedSignatureHash: "f07f45f3634d3ea8c0f2cb676f56e20993edf9be07a83bf0dfdb3debcf1441bf"},
		{name: "native-none-anyonecanpay-0-modify-script-public-key", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyScriptPublicKey(0), // should change the hash
			expectedSignatureHash: "20a525c54dc33b2a61201f05233c086dbe8e06e9515775181ed96550b4f2d714"},

		// sigHashSingle
		{name: "native-single-0", tx: nativeTx, hashType: single, inputIndex: 0,
			expectedSignatureHash: "44a0b407ff7b239d447743dd503f7ad23db5b2ee4d25279bd3dffaf6b474e005"},
		{name: "native-single-0-modify-output-0", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifyOutput(0), // should change the hash
			expectedSignatureHash: "0fcaca1211f7ea44997717eb84c04c9c807db8b59797bc6800c2ccb135a5271c"},
		{name: "native-single-0-modify-output-1", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "44a0b407ff7b239d447743dd503f7ad23db5b2ee4d25279bd3dffaf6b474e005"},
		{name: "native-single-0-modify-sequence-0", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifySequence(0), // should change the hash
			expectedSignatureHash: "83796d22879718eee1165d4aace667bb6778075dab579c32c57be945f466a451"},
		{name: "native-single-0-modify-sequence-1", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "44a0b407ff7b239d447743dd503f7ad23db5b2ee4d25279bd3dffaf6b474e005"},
		{name: "native-single-2-no-corresponding-output", tx: nativeTx, hashType: single, inputIndex: 2,
			expectedSignatureHash: "022ad967192f39d8d5895d243e025ec14cc7a79708c5e364894d4eff3cecb1b0"},
		{name: "native-single-2-no-corresponding-output-modify-output-1", tx: nativeTx, hashType: single, inputIndex: 2,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "022ad967192f39d8d5895d243e025ec14cc7a79708c5e364894d4eff3cecb1b0"},
		{name: "native-single-anyonecanpay-0", tx: nativeTx, hashType: singleAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "43b20aba775050cf9ba8d5e48fc7ed2dc6c071d23f30382aea58b7c59cfb8ed7"},
		{name: "native-single-anyonecanpay-2-no-corresponding-output", tx: nativeTx, hashType: singleAnyoneCanPay, inputIndex: 2,
			expectedSignatureHash: "846689131fb08b77f83af1d3901076732ef09d3f8fdff945be89aa4300562e5f"},

		// subnetwork transaction
		{name: "subnetwork-all-0", tx: subnetworkTx, hashType: all, inputIndex: 0,
			expectedSignatureHash: "b2f421c933eb7e1a91f1d9e1efa3f120fe419326c0dbac487752189522550e0c"},
		{name: "subnetwork-all-modify-payload", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyPayload, // should change the hash
			expectedSignatureHash: "12ab63b9aea3d58db339245a9b6e9cb6075b2253615ce0fb18104d28de4435a1"},
		{name: "subnetwork-all-modify-gas", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyGas, // should change the hash
			expectedSignatureHash: "2501edfc0068d591160c4bd98646c6e6892cdc051182a8be3ccd6d67f104fd17"},
		{name: "subnetwork-all-subnetwork-id", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifySubnetworkID, // should change the hash
			expectedSignatureHash: "a5d1230ede0dfcfd522e04123a7bcd721462fed1d3a87352031a4f6e3c4389b6"},
	}

	for _, test := range tests {
		tx := test.tx
		if test.modificationFunction != nil {
			tx = test.modificationFunction(tx)
		}

		actualSignatureHash, err := consensushashing.CalculateSignatureHashSchnorr(
			tx, test.inputIndex, test.hashType, &consensushashing.SighashReusedValues{})
		if err != nil {
			t.Errorf("%s: Error from CalculateSignatureHashSchnorr: %+v", test.name, err)
			continue
		}

		if actualSignatureHash.String() != test.expectedSignatureHash {
			t.Errorf("%s: expected signature hash: '%s'; but got: '%s'",
				test.name, test.expectedSignatureHash, actualSignatureHash)
		}
	}
}

func TestCalculateSignatureHashECDSA(t *testing.T) {
	nativeTx, subnetworkTx, err := generateTxs()
	if err != nil {
		t.Fatalf("Error from generateTxs: %+v", err)
	}

	// Note: Expected values were generated by the same code that they test,
	// As long as those were not verified using 3rd-party code they only check for regression, not correctness
	tests := []struct {
		name                  string
		tx                    *externalapi.DomainTransaction
		hashType              consensushashing.SigHashType
		inputIndex            int
		modificationFunction  func(*externalapi.DomainTransaction) *externalapi.DomainTransaction
		expectedSignatureHash string
	}{
		// native transactions

		// sigHashAll
		{name: "native-all-0", tx: nativeTx, hashType: all, inputIndex: 0,
			expectedSignatureHash: "1d679268414c20ffe952e3c255befd892e60e86ae1657fce8a20225e5dc87d64"},
		{name: "native-all-0-modify-input-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyInput(1), // should change the hash
			expectedSignatureHash: "c573469b9ec6551507371d26eaa75417905420577f56d0277c4234a99f2305d7"},
		{name: "native-all-0-modify-output-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // should change the hash
			expectedSignatureHash: "a92450b0662c120db3993b6bb258d06d2bcb983aa591d85f97adf8b7207a5db5"},
		{name: "native-all-0-modify-sequence-1", tx: nativeTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifySequence(1), // should change the hash
			expectedSignatureHash: "c7a7524096499e4401a1592f892bada1afe7f7d276c4f0e691c993e17c03cf7d"},
		{name: "native-all-anyonecanpay-0", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "13270aeb5b56d844d064d5d2cf06af7dbd0341fd55069b9af56d5e3c99f2eddd"},
		{name: "native-all-anyonecanpay-0-modify-input-0", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyInput(0), // should change the hash
			expectedSignatureHash: "981959e8c427ba4a06c3d53abc93514ba179d8cc7e94daeb4f516a0c2c685f86"},
		{name: "native-all-anyonecanpay-0-modify-input-1", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyInput(1), // shouldn't change the hash
			expectedSignatureHash: "13270aeb5b56d844d064d5d2cf06af7dbd0341fd55069b9af56d5e3c99f2eddd"},
		{name: "native-all-anyonecanpay-0-modify-sequence", tx: nativeTx, hashType: allAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "13270aeb5b56d844d064d5d2cf06af7dbd0341fd55069b9af56d5e3c99f2eddd"},

		// sigHashNone
		{name: "native-none-0", tx: nativeTx, hashType: none, inputIndex: 0,
			expectedSignatureHash: "a45955ca970039160bb91b1ea42e070b4ff21598654aad91c562e8b9af922c5f"},
		{name: "native-none-0-modify-output-1", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "a45955ca970039160bb91b1ea42e070b4ff21598654aad91c562e8b9af922c5f"},
		{name: "native-none-0-modify-sequence-0", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifySequence(0), // should change the hash
			expectedSignatureHash: "e1a430a24d77bc259ae572e1505dd67d3444ba0ffbc7918e06ae7e907e575adb"},
		{name: "native-none-0-modify-sequence-1", tx: nativeTx, hashType: none, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "a45955ca970039160bb91b1ea42e070b4ff21598654aad91c562e8b9af922c5f"},
		{name: "native-none-anyonecanpay-0", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "6bae2a0f1f7b9fd59804f4720a1a918be31b7ec12e184585fa16bda8c7f8c35c"},
		{name: "native-none-anyonecanpay-0-modify-amount-spent", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyAmountSpent(0), // should change the hash
			expectedSignatureHash: "6653d3d882d2ffc1c3ff5b7ccf07f7970c5973b824abb5b117803809c5a884c7"},
		{name: "native-none-anyonecanpay-0-modify-script-public-key", tx: nativeTx, hashType: noneAnyoneCanPay, inputIndex: 0,
			modificationFunction:  modifyScriptPublicKey(0), // should change the hash
			expectedSignatureHash: "da3cb9094d905de69b3881cf8d4e2d5268bcf029dec5b62a972fcab90e6abde9"},

		// sigHashSingle
		{name: "native-single-0", tx: nativeTx, hashType: single, inputIndex: 0,
			expectedSignatureHash: "964d03d8477dd468f3d9933676b5b4a976a68fee1760eae037be4247c36cc207"},
		{name: "native-single-0-modify-output-0", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifyOutput(0), // should change the hash
			expectedSignatureHash: "7c51b4a7c6a6e786b1c420c859c2853131d7041b8ba8de72cbcd026b2e0d511b"},
		{name: "native-single-0-modify-output-1", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "964d03d8477dd468f3d9933676b5b4a976a68fee1760eae037be4247c36cc207"},
		{name: "native-single-0-modify-sequence-0", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifySequence(0), // should change the hash
			expectedSignatureHash: "d31e5a9e71560d2f66483e7e4e7d41418b66cd22814450848a2d8fa78045d6a0"},
		{name: "native-single-0-modify-sequence-1", tx: nativeTx, hashType: single, inputIndex: 0,
			modificationFunction:  modifySequence(1), // shouldn't change the hash
			expectedSignatureHash: "964d03d8477dd468f3d9933676b5b4a976a68fee1760eae037be4247c36cc207"},
		{name: "native-single-2-no-corresponding-output", tx: nativeTx, hashType: single, inputIndex: 2,
			expectedSignatureHash: "667b6b65682a6e1e14aec699a177d22ce1392661828e54dcd97cd83b05233d41"},
		{name: "native-single-2-no-corresponding-output-modify-output-1", tx: nativeTx, hashType: single, inputIndex: 2,
			modificationFunction:  modifyOutput(1), // shouldn't change the hash
			expectedSignatureHash: "667b6b65682a6e1e14aec699a177d22ce1392661828e54dcd97cd83b05233d41"},
		{name: "native-single-anyonecanpay-0", tx: nativeTx, hashType: singleAnyoneCanPay, inputIndex: 0,
			expectedSignatureHash: "a11c2fbcd4f09bffce9e5fca62323388a2cf9037fd3be66211c7869c067123a2"},
		{name: "native-single-anyonecanpay-2-no-corresponding-output", tx: nativeTx, hashType: singleAnyoneCanPay, inputIndex: 2,
			expectedSignatureHash: "00f429dfb9150d7a96aa3f179bcc6f8fbf9bce481f00c6bb97af67e108e5d0ff"},

		// subnetwork transaction
		{name: "subnetwork-all-0", tx: subnetworkTx, hashType: all, inputIndex: 0,
			expectedSignatureHash: "807d351414ff592ba097daa5c7937311d6382107f23a6ae415954e248a0527e0"},
		{name: "subnetwork-all-modify-payload", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyPayload, // should change the hash
			expectedSignatureHash: "0bb2a9a37cc27a60c91c1c9b5ff29bc09f1b39faa3ec55edb15dcbc6c9ce03d7"},
		{name: "subnetwork-all-modify-gas", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifyGas, // should change the hash
			expectedSignatureHash: "78dcfa1ea6a6f01c31805bda3cc71d7356f32b87a8bf3b80b4a4d0d5f95e8741"},
		{name: "subnetwork-all-subnetwork-id", tx: subnetworkTx, hashType: all, inputIndex: 0,
			modificationFunction:  modifySubnetworkID, // should change the hash
			expectedSignatureHash: "6412917f0d5d856c37897d9a98c3817dc1f1668deff73efeefbe2529e00e3511"},
	}

	for _, test := range tests {
		tx := test.tx
		if test.modificationFunction != nil {
			tx = test.modificationFunction(tx)
		}

		actualSignatureHash, err := consensushashing.CalculateSignatureHashECDSA(
			tx, test.inputIndex, test.hashType, &consensushashing.SighashReusedValues{})
		if err != nil {
			t.Errorf("%s: Error from CalculateSignatureHashECDSA: %+v", test.name, err)
			continue
		}

		if actualSignatureHash.String() != test.expectedSignatureHash {
			t.Errorf("%s: expected signature hash: '%s'; but got: '%s'",
				test.name, test.expectedSignatureHash, actualSignatureHash)
		}
	}
}

func generateTxs() (nativeTx, subnetworkTx *externalapi.DomainTransaction, err error) {
	genesisCoinbase := dagconfig.SimnetParams.GenesisBlock.Transactions[0]
	genesisCoinbaseTransactionID := consensushashing.TransactionID(genesisCoinbase)

	address1Str := "coinsecsim:qzpj2cfa9m40w9m2cmr8pvfuqpp32mzzwsuw6ukhfduqpp32mzzws59e8fapc"
	address1, err := util.DecodeAddress(address1Str, util.Bech32PrefixCoinsecSim)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding address1: %+v", err)
	}
	address1ToScript, err := txscript.PayToAddrScript(address1)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating script: %+v", err)
	}

	address2Str := "coinsecsim:qr7w7nqsdnc3zddm6u8s9fex4ysk95hm3v30q353ymuqpp32mzzws59e8fapc"
	address2, err := util.DecodeAddress(address2Str, util.Bech32PrefixCoinsecSim)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding address2: %+v", err)
	}
	address2ToScript, err := txscript.PayToAddrScript(address2)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating script: %+v", err)
	}

	txIns := []*externalapi.DomainTransactionInput{
		{
			PreviousOutpoint: *externalapi.NewDomainOutpoint(genesisCoinbaseTransactionID, 0),
			Sequence:         0,
			UTXOEntry:        utxo.NewUTXOEntry(100, address1ToScript, false, 0),
		},
		{
			PreviousOutpoint: *externalapi.NewDomainOutpoint(genesisCoinbaseTransactionID, 1),
			Sequence:         1,
			UTXOEntry:        utxo.NewUTXOEntry(200, address2ToScript, false, 0),
		},
		{
			PreviousOutpoint: *externalapi.NewDomainOutpoint(genesisCoinbaseTransactionID, 2),
			Sequence:         2,
			UTXOEntry:        utxo.NewUTXOEntry(300, address2ToScript, false, 0),
		},
	}

	txOuts := []*externalapi.DomainTransactionOutput{
		{
			Value:           300,
			ScriptPublicKey: address2ToScript,
		},
		{
			Value:           300,
			ScriptPublicKey: address1ToScript,
		},
	}

	nativeTx = &externalapi.DomainTransaction{
		Version:      0,
		Inputs:       txIns,
		Outputs:      txOuts,
		LockTime:     1615462089000,
		SubnetworkID: subnetworks.SubnetworkIDNative,
	}
	subnetworkTx = &externalapi.DomainTransaction{
		Version:      0,
		Inputs:       txIns,
		Outputs:      txOuts,
		LockTime:     1615462089000,
		SubnetworkID: externalapi.DomainSubnetworkID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Gas:          250,
		Payload:      []byte{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}

	return nativeTx, subnetworkTx, nil
}

func BenchmarkCalculateSignatureHashSchnorr(b *testing.B) {
	sigHashTypes := []consensushashing.SigHashType{
		consensushashing.SigHashAll,
		consensushashing.SigHashNone,
		consensushashing.SigHashSingle,
		consensushashing.SigHashAll | consensushashing.SigHashAnyOneCanPay,
		consensushashing.SigHashNone | consensushashing.SigHashAnyOneCanPay,
		consensushashing.SigHashSingle | consensushashing.SigHashAnyOneCanPay}

	for _, size := range []int{10, 100, 1000} {
		tx := generateTransaction(b, sigHashTypes, size)

		b.Run(fmt.Sprintf("%d-inputs-and-outputs", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reusedValues := &consensushashing.SighashReusedValues{}
				for inputIndex := range tx.Inputs {
					sigHashType := sigHashTypes[inputIndex%len(sigHashTypes)]
					_, err := consensushashing.CalculateSignatureHashSchnorr(tx, inputIndex, sigHashType, reusedValues)
					if err != nil {
						b.Fatalf("Error from CalculateSignatureHashSchnorr: %+v", err)
					}
				}
			}
		})
	}
}

func generateTransaction(b *testing.B, sigHashTypes []consensushashing.SigHashType, inputAndOutputSizes int) *externalapi.DomainTransaction {
	sourceScript := getSourceScript(b)
	tx := &externalapi.DomainTransaction{
		Version:      0,
		Inputs:       generateInputs(inputAndOutputSizes, sourceScript),
		Outputs:      generateOutputs(inputAndOutputSizes, sourceScript),
		LockTime:     123456789,
		SubnetworkID: externalapi.DomainSubnetworkID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Gas:          125,
		Payload:      []byte{9, 8, 7, 6, 5, 4, 3, 2, 1},
		Fee:          0,
		Mass:         0,
		ID:           nil,
	}
	signTx(b, tx, sigHashTypes)
	return tx
}

func signTx(b *testing.B, tx *externalapi.DomainTransaction, sigHashTypes []consensushashing.SigHashType) {
	sourceAddressPKStr := "a4d85b7532123e3dd34e58d7ce20895f7ca32349e29b01700bb5a3e72d2570eb"
	privateKeyBytes, err := hex.DecodeString(sourceAddressPKStr)
	if err != nil {
		b.Fatalf("Error parsing private key hex: %+v", err)
	}
	keyPair, err := secp256k1.DeserializeSchnorrPrivateKeyFromSlice(privateKeyBytes)
	if err != nil {
		b.Fatalf("Error deserializing private key: %+v", err)
	}
	for i, txIn := range tx.Inputs {
		signatureScript, err := txscript.SignatureScript(
			tx, i, sigHashTypes[i%len(sigHashTypes)], keyPair, &consensushashing.SighashReusedValues{})
		if err != nil {
			b.Fatalf("Error from SignatureScript: %+v", err)
		}
		txIn.SignatureScript = signatureScript
	}

}

func generateInputs(size int, sourceScript *externalapi.ScriptPublicKey) []*externalapi.DomainTransactionInput {
	inputs := make([]*externalapi.DomainTransactionInput, size)

	for i := 0; i < size; i++ {
		inputs[i] = &externalapi.DomainTransactionInput{
			PreviousOutpoint: *externalapi.NewDomainOutpoint(
				externalapi.NewDomainTransactionIDFromByteArray(&[32]byte{12, 3, 4, 5}), 1),
			SignatureScript: nil,
			Sequence:        uint64(i),
			UTXOEntry:       utxo.NewUTXOEntry(uint64(i), sourceScript, false, 12),
		}
	}

	return inputs
}

func getSourceScript(b *testing.B) *externalapi.ScriptPublicKey {
	sourceAddressStr := "coinsecsim:qz6f9z6l3x4v3lf9mgf0t934th4nx5kgzu663x9yjh"

	sourceAddress, err := util.DecodeAddress(sourceAddressStr, util.Bech32PrefixCoinsecSim)
	if err != nil {
		b.Fatalf("Error from DecodeAddress: %+v", err)
	}

	sourceScript, err := txscript.PayToAddrScript(sourceAddress)
	if err != nil {
		b.Fatalf("Error from PayToAddrScript: %+v", err)
	}
	return sourceScript
}

func generateOutputs(size int, script *externalapi.ScriptPublicKey) []*externalapi.DomainTransactionOutput {
	outputs := make([]*externalapi.DomainTransactionOutput, size)

	for i := 0; i < size; i++ {
		outputs[i] = &externalapi.DomainTransactionOutput{
			Value:           uint64(i),
			ScriptPublicKey: script,
		}
	}

	return outputs
}
