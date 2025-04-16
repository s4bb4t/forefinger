package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/s4bb4t/forefinger/pkg/methods"
)

func TestFilter_FromBlock(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantError bool
	}{
		{"ValidBigInt", big.NewInt(12345), false},
		{"ValidStringTag", methods.Latest, false},
		{"InvalidStringTag", "random", true},
		{"InvalidType", 12345, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.FromBlock(tt.input)
			if (f.Validate() != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", f.Validate())
			}
			t.Log(f.debugRange())
		})
	}
}

func TestFilter_ToBlock(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantError bool
	}{
		{"ValidBigInt", big.NewInt(54321), false},
		{"ValidStringTag", methods.Latest, false},
		{"InvalidStringTag", "invalidTag", true},
		{"InvalidType", 54321.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.ToBlock(tt.input)
			if (f.Validate() != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", f.Validate())
			}
			t.Log(f.debugRange())
		})
	}
}

func TestFilter_AddTopic(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantError bool
	}{
		{"SingleHash", common.HexToHash("0x1234"), false},
		{"PointerToSingleHash", common.HexToHash("0x1234"), false},
		{"SliceOfHashes", []common.Hash{common.HexToHash("0x1234")}, false},
		{"PointerToSliceOfHashes", &[]common.Hash{common.HexToHash("0x1234")}, false},
		{"InvalidType", "0x1234", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.AddTopic(tt.input)
			if (f.Validate() != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", f.Validate())
			}
			t.Log(f.topics)
		})
	}
}

func TestFilter_AddAddress(t *testing.T) {
	tests := []struct {
		name  string
		input common.Address
	}{
		{"SingleAddress", common.HexToAddress("0x1234")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.AddAddress(tt.input)
			if len(f.address.addr) == 0 {
				t.Errorf("expected address to be added to the filter")
			}
		})
	}
}

func TestFilter_Validate(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(f *Filter)
		wantError bool
	}{
		{"NoError", func(f *Filter) {}, false},
		{"WithError", func(f *Filter) {
			f.addError(errors.New("test error"))
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			tt.setup(f)
			if (f.Validate() != nil) != tt.wantError {
				t.Errorf("unexpected error: %v", f.Validate())
			}
		})
	}
}

func TestFilter_MarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Filter
		expectErr bool
	}{
		{
			"ValidFilterWithAddress",
			func() *Filter {
				return NewFilter().FromBlock(methods.Latest).ToBlock(methods.Latest).AddAddress(common.HexToAddress("0x123"))
			},
			false,
		},
		{
			"ValidFilterWithMultipleAddresses",
			func() *Filter {
				return NewFilter().
					FromBlock(methods.Latest).
					ToBlock(methods.Earliest).
					AddAddresses([]common.Address{
						common.HexToAddress("0x1230000000000000000000000000000000000123"),
						common.HexToAddress("0x456"),
					})
			},
			false,
		},
		{
			"InvalidFilterWithError",
			func() *Filter {
				f := NewFilter()
				f.addError(fmt.Errorf("forced error"))
				return f
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup()
			data, err := f.MarshalJSON()
			if (err != nil) != tt.expectErr {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tt.expectErr {
				fmt.Println(string(data))

				var unmarshalled Filter
				if err := json.Unmarshal(data, &unmarshalled); err != nil {
					t.Fatalf("unexpected error during unmarshaling: %v", err)
				}
			}
		})
	}
}

func TestFilter_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		jsonData  string
		expectErr bool
	}{
		{
			"ValidDataWithSingleAddress",
			`{"fromBlock":"latest","toBlock":"latest","address":"0x123"}`,
			false,
		},
		{
			"ValidDataWithMultipleAddressesAndRange",
			`{"fromBlock":"earliest","toBlock":"pending","address":["0x123","0x456"]}`,
			false,
		},
		{
			"InvalidJsonStructure",
			`{"fromBlock":latest,toBlock:}`,
			true,
		},
		{
			"MissingAddressField",
			`{"fromBlock":"latest", "toBlock":"earliest"}`,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Filter
			err := f.UnmarshalJSON([]byte(tt.jsonData))
			if (err != nil) != tt.expectErr {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestFilter_setRangeString(t *testing.T) {
	tests := []struct {
		name         string
		tag          string
		isFromBlock  bool
		expectTag    string
		expectSwitch bool
		wantError    bool
	}{
		{"ValidFromTag", methods.Latest, true, methods.Latest, true, false},
		{"ValidToTag", methods.Earliest, false, methods.Earliest, true, false},
		{"InvalidFromTag", "invalid", true, "", false, true},
		{"InvalidToTag", "invalid", false, "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.setRangeString(tt.tag, tt.isFromBlock)
			if tt.wantError {
				if f.Validate() == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if tt.isFromBlock {
				if f.fromBlock.tag != tt.expectTag || f.fromBlock.tagSwitch.Load() != tt.expectSwitch {
					t.Errorf("unexpected result: got (%s, %v), expected (%s, %v)", f.fromBlock.tag, f.fromBlock.tagSwitch.Load(), tt.expectTag, tt.expectSwitch)
				}
			} else {
				if f.toBlock.tag != tt.expectTag || f.toBlock.tagSwitch.Load() != tt.expectSwitch {
					t.Errorf("unexpected result: got (%s, %v), expected (%s, %v)", f.toBlock.tag, f.toBlock.tagSwitch.Load(), tt.expectTag, tt.expectSwitch)
				}
			}
		})
	}
}

func TestFilter_setRangeInt(t *testing.T) {
	tests := []struct {
		name         string
		value        *big.Int
		isFromBlock  bool
		expected     *big.Int
		expectedBool bool
	}{
		{"ValidFromInt", big.NewInt(12345), true, big.NewInt(12345), false},
		{"ValidToInt", big.NewInt(54321), false, big.NewInt(54321), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFilter()
			f.setRangeInt(tt.value, tt.isFromBlock)
			if tt.isFromBlock {
				if f.fromBlock.n.Cmp(tt.expected) != 0 || f.fromBlock.tagSwitch.Load() != tt.expectedBool {
					t.Errorf("unexpected result: got (%v, %v), want (%v, %v)", f.fromBlock.n, f.fromBlock.tagSwitch.Load(), tt.expected, tt.expectedBool)
				}
			} else {
				if f.toBlock.n.Cmp(tt.expected) != 0 || f.toBlock.tagSwitch.Load() != tt.expectedBool {
					t.Errorf("unexpected result: got (%v, %v), want (%v, %v)", f.toBlock.n, f.toBlock.tagSwitch.Load(), tt.expected, tt.expectedBool)
				}
			}
		})
	}
}
