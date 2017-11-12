/*
Copyright 2017 IBM Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// MockStub GetCreator is not implemented.

// TODO some verification of the state of products might be good to do..

package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	uuid "github.com/satori/go.uuid"
)

func Test_Init(t *testing.T) {
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)

	// instantiate without args (good)
	t.Run("without arg-good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		var args [][]byte
		res := stub.MockInit("1", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	// instantiate with args (bad)
	t.Run("with arg-bad", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("init"), []byte("")}
		res := stub.MockInit("", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})
}

// createProduct and getIndex tests
func Test_Invoke(t *testing.T) {

	uuidStrings := []string{
		"75878381-fed5-4c86-8ef5-e92544e32b04",
		"fb1abb0e-391d-4786-b307-e6bdb0f46bb0",
		"f1f51fbe-e927-42b4-a456-90e2e3d31224",
		"bf89b66d-1f6a-4d7c-992b-834c8f7fea58",
		"d3b8a8d2-d248-4f48-9196-3f5cb008aad2",
		"fa33c333-236f-44ce-bd83-062f8eaa51b8",
		"c0f27d99-09f2-479e-aacd-13bf8d9666bf",
		"aebc1142-a1bd-4c8a-a082-6a22356563aa",
		"d67ee172-36b2-4681-b844-264fbde72084",
		"61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296",
	}

	// initialize ledger
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("1", [][]byte{})

	t.Run("createProduct good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[0])}
		res = stub.MockInvoke("1", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("createProduct no arg", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("createProduct")}
		res = stub.MockInvoke("2", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("createProduct bad arg", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("createProduct"), []byte("garbage")}
		res = stub.MockInvoke("3", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("createProduct duplicate serial number", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[0])}
		res = stub.MockInvoke("4", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("createProduct 9 more good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		for i := 1; i < 10; i++ {
			args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[i])}
			res = stub.MockInvoke("5", args)
			fmt.Println("Response Message:", res.GetMessage())
			if res.Status == shim.OK {
				fmt.Println(" - OK")
			} else {
				t.Fail()
			}
		}
	})

	t.Run("getIndex with 10 created", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("getIndex")}
		res = stub.MockInvoke("6", args)
		var index []uuid.UUID
		json.Unmarshal(res.GetPayload(), &index)
		fmt.Println("Response Message:", res.GetMessage())
		if len(index) == 10 {
			fmt.Println(" - OK")
		} else {
			fmt.Println("unable to retrieve all indexes..")
			t.Fail()
		}
	})
}

func Test_Invoke_Invalid_Function(t *testing.T) {
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("garbage", [][]byte{})
	args := [][]byte{[]byte("garbage")}
	res = stub.MockInvoke("garbage", args)
	fmt.Println("Response Message:", res.GetMessage())
	if res.Status == shim.ERROR {
		fmt.Println(" - OK")
	} else {
		t.Fail()
	}
}

func Test_Invoke_getProduct(t *testing.T) {

	uuidStrings := []string{
		"75878381-fed5-4c86-8ef5-e92544e32b04",
		"fb1abb0e-391d-4786-b307-e6bdb0f46bb0",
		"f1f51fbe-e927-42b4-a456-90e2e3d31224",
		"bf89b66d-1f6a-4d7c-992b-834c8f7fea58",
		"d3b8a8d2-d248-4f48-9196-3f5cb008aad2",
	}
	/*
			"fa33c333-236f-44ce-bd83-062f8eaa51b8",
			"c0f27d99-09f2-479e-aacd-13bf8d9666bf",
			"aebc1142-a1bd-4c8a-a082-6a22356563aa",
			"d67ee172-36b2-4681-b844-264fbde72084",
			"61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296",
		}
	*/

	// initialize ledger
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("7", [][]byte{})
	for i := 0; i < 5; i++ {
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[i])}
		res = stub.MockInvoke("7", args)
	}

	t.Run("getProduct good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("getProduct"), []byte(uuidStrings[0])}
		res = stub.MockInvoke("8", args)
		fmt.Println("Response Payload:", string(res.GetPayload()))
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("getProduct no args", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("getProduct")}
		res = stub.MockInvoke("8", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("getProduct bad arg", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("getProduct"), []byte("garbage")}
		res = stub.MockInvoke("8", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("getProduct doesn't exist", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("getProduct"), []byte("fa33c333-236f-44ce-bd83-062f8eaa51b8")}
		res = stub.MockInvoke("8", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})
}

func Test_Invoke_offerProduct(t *testing.T) {
	uuidStrings := []string{
		"75878381-fed5-4c86-8ef5-e92544e32b04",
		"fb1abb0e-391d-4786-b307-e6bdb0f46bb0",
		"f1f51fbe-e927-42b4-a456-90e2e3d31224",
		"bf89b66d-1f6a-4d7c-992b-834c8f7fea58",
		"d3b8a8d2-d248-4f48-9196-3f5cb008aad2",
	}
	// initialize ledger
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("9", [][]byte{})
	for i := 0; i < 5; i++ {
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[i])}
		res = stub.MockInvoke("9", args)
	}

	t.Run("offerProduct good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("offerProduct"), []byte(uuidStrings[0])}
		res = stub.MockInvoke("10", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("offerProduct already offered", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("offerProduct"), []byte(uuidStrings[0])}
		res = stub.MockInvoke("11", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("offerProduct incorrect num args", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("offerProduct")}
		res = stub.MockInvoke("12", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("offerProduct bad uuid", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("offerProduct"), []byte("garbage")}
		res = stub.MockInvoke("13", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})
}

func Test_Invoke_tradeProduct(t *testing.T) {
	uuidStrings := []string{
		"75878381-fed5-4c86-8ef5-e92544e32b04",
		"fb1abb0e-391d-4786-b307-e6bdb0f46bb0",
		"f1f51fbe-e927-42b4-a456-90e2e3d31224",
		"bf89b66d-1f6a-4d7c-992b-834c8f7fea58",
		"d3b8a8d2-d248-4f48-9196-3f5cb008aad2",
	}

	acquiredProductBytes := []byte(`{"serial_number":"61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296","producer_id":"bogus","produced_date":1510464632,"offerer_id":"bogus2","offered_date":1510864632,"trader_id":"","traded_date":0,"traded_for_serial_number":"","consumer_id":"","consumed_date":0,"product_type":"orgX","state":"offered"}`)

	// initialize ledger
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("14", [][]byte{})
	for i := 0; i < 5; i++ {
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[i])}
		res = stub.MockInvoke("14", args)
	}

	t.Run("tradeProduct good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("tradeProduct"), []byte("61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296"), acquiredProductBytes}
		res = stub.MockInvoke("15", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("tradeProduct incorrect num args", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("tradeProduct")}
		res = stub.MockInvoke("16", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("tradeProduct bad uuid", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("tradeProduct"), []byte("garbage"), acquiredProductBytes}
		res = stub.MockInvoke("17", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("tradeProduct bad uuid", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("tradeProduct"), []byte("garbage"), acquiredProductBytes}
		res = stub.MockInvoke("17", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})
}

func Test_Invoke_consumeProduct(t *testing.T) {
	uuidStrings := []string{
		"75878381-fed5-4c86-8ef5-e92544e32b04",
		"fb1abb0e-391d-4786-b307-e6bdb0f46bb0",
		"f1f51fbe-e927-42b4-a456-90e2e3d31224",
		"bf89b66d-1f6a-4d7c-992b-834c8f7fea58",
		"d3b8a8d2-d248-4f48-9196-3f5cb008aad2",
	}

	acquiredProductBytes := []byte(`{"serial_number":"61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296","producer_id":"bogus","produced_date":1510464632,"offerer_id":"bogus2","offered_date":1510864632,"trader_id":"","traded_date":0,"traded_for_serial_number":"","consumer_id":"","consumed_date":0,"product_type":"orgX","state":"offered"}`)

	// initialize ledger
	scc := new(ProductChaincode)
	stub := shim.NewMockStub("procon", scc)
	res := stub.MockInit("18", [][]byte{})
	for i := 0; i < 5; i++ {
		args := [][]byte{[]byte("createProduct"), []byte(uuidStrings[i])}
		res = stub.MockInvoke("18", args)
	}
	args := [][]byte{[]byte("tradeProduct"), []byte("61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296"), acquiredProductBytes}
	res = stub.MockInvoke("18", args)

	t.Run("consumeProduct good", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct"), []byte("61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296")}
		res = stub.MockInvoke("19", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.OK {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("consumeProduct wrong num args", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct")}
		res = stub.MockInvoke("20", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("consumeProduct bad uuid", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct"), []byte("garbage")}
		res = stub.MockInvoke("21", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("consumeProduct unknown uuid", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct"), []byte("aebc1142-a1bd-4c8a-a082-6a22356563ab")}
		res = stub.MockInvoke("22", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("consumeProduct already consumed", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct"), []byte("61f4d0ce-f4ea-4bbb-8893-1ea7a7d4b296")}
		res = stub.MockInvoke("23", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})

	t.Run("consumeProduct not traded", func(t *testing.T) {
		fmt.Println(t.Name() + ":")
		args := [][]byte{[]byte("consumeProduct"), []byte("fb1abb0e-391d-4786-b307-e6bdb0f46bb0")}
		res = stub.MockInvoke("23", args)
		fmt.Println("Response Message:", res.GetMessage())
		if res.Status == shim.ERROR {
			fmt.Println(" - OK")
		} else {
			t.Fail()
		}
	})
}

func Test_Main(t *testing.T) {
	main()
}
