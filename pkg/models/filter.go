package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"math/big"
	"sync/atomic"
)

type (
	// additional represents topic filters used in Ethereum logs.
	// single stores a list of topics that match in any position.
	// multi stores multiple sets of topics allowing filtering logs at specific positions.
	additional struct {
		topics [][]common.Hash
	}

	// quantity represents a block range, either as a numeric value or a tag (e.g., "latest").
	// tagSwitch ensures atomic updates between numeric and tag representations.
	quantity struct {
		n         *big.Int
		tag       string
		tagSwitch atomic.Bool
	}

	// scope represents a list of Ethereum addresses for filtering logs.
	scope struct {
		addr []common.Address
	}

	// Filter is the main struct for constructing Ethereum log filters.
	// It includes block ranges, address filters, topic filters, and an internal error state.
	Filter struct {
		fromBlock quantity
		toBlock   quantity
		address   scope
		topics    additional
		_err      error
	}
)

// NF is an alias for NewFilter().
func NF() *Filter {
	return NewFilter()
}

// NewFilter creates and returns a new instance of the Filter struct with default values.
// The default values for BlockFrom and BlockTo are set to 'latest'.
func NewFilter() *Filter {
	return &Filter{
		fromBlock: quantity{
			tag:       methods.Latest,
			tagSwitch: atomic.Bool{},
		},
		toBlock: quantity{
			tag:       methods.Latest,
			tagSwitch: atomic.Bool{},
		},
	}
}

func (f *Filter) debugRange() (res string) {
	if f.fromBlock.tagSwitch.Load() {
		res += f.fromBlock.tag
	} else {
		res += f.fromBlock.n.String()
	}
	if f.toBlock.tagSwitch.Load() {
		res += " --> " + f.toBlock.tag
	} else {
		res += " --> " + f.toBlock.n.String()
	}
	return
}

// FromBlock sets the `fromBlock` in the Filter.
// Accepts a *big.Int or a string as valid input. Adds an error if the input is invalid.
func (f *Filter) FromBlock(val any) *Filter {
	switch v := val.(type) {
	case *big.Int:
		f.setRangeInt(v, true)
	case string:
		f.setRangeString(v, true)
	default:
		f.addError(fmt.Errorf("invalid block range: %v, type of %T", val, val))
	}
	return f
}

// ToBlock sets the `toBlock` in the Filter.
// Accepts a *big.Int or a string as valid input. Adds an error if the input is invalid.
func (f *Filter) ToBlock(val any) *Filter {
	switch v := val.(type) {
	case *big.Int:
		f.setRangeInt(v, false)
	case string:
		f.setRangeString(v, false)
	default:
		f.addError(fmt.Errorf("invalid block range: %v, type of %T", val, val))
	}
	return f
}

// AddTopic appends topics to the filter's topic criteria for Ethereum logs.
// Topics are processed in order and are position-dependent:
// - The first topic added matches the logs' first indexed topic.
// - The second topic added matches the logs' second indexed topic, and so on.
// Input types supported:
// - []common.Hash: Adds multiple topics for filtering logs matching any of these topics in a specific position.
// - *[]common.Hash: Dereferences the pointer and appends the topics to the multi-topic list.
// - common.Hash: Appends a single topic for filtering logs in a specific position.
// - *common.Hash: Dereferences the pointer and appends to the single-topic list.
// - []*common.Hash: Dereferences each pointer and appends the topics to the single-topic list.
func (f *Filter) AddTopic(val any) *Filter {
	switch v := val.(type) {
	case []common.Hash:
		f.topics.topics = append(f.topics.topics, v)
	case *[]common.Hash:
		f.topics.topics = append(f.topics.topics, *v)
	case common.Hash:
		f.topics.topics = append(f.topics.topics, []common.Hash{v})
	case *common.Hash:
		f.topics.topics = append(f.topics.topics, []common.Hash{*v})
	case []*common.Hash:
		for _, hash := range v {
			f.topics.topics = append(f.topics.topics, []common.Hash{*hash})
		}
	default:
		f.addError(fmt.Errorf("invalid topic: %v, type of %T", val, val))
	}
	return f
}

// AddAddress appends an Ethereum address to the filter's address list.
func (f *Filter) AddAddress(addr common.Address) *Filter {
	f.address.addr = append(f.address.addr, addr)
	return f
}

// AddAddresses appends multiple Ethereum addresses to the filter's address list and returns the updated filter.
func (f *Filter) AddAddresses(addr []common.Address) *Filter {
	f.address.addr = append(f.address.addr, addr...)
	return f
}

// Validate checks for errors in the Filter configuration and wraps them in a prefixed error string if any exist.
func (f *Filter) Validate() error {
	if f._err != nil {
		return fmt.Errorf("forefinger: filter: %w", f._err)
	}
	return nil
}

// setRangeString updates the block range as a tag (e.g., "earliest", "latest").
func (f *Filter) setRangeString(tag string, isFromBlock bool) {
	switch isFromBlock {
	case true:
		switch tag {
		case methods.Earliest, methods.Latest, methods.Pending, methods.Safe, methods.Finalized:
			f.fromBlock.tag = tag
			f.fromBlock.tagSwitch.CompareAndSwap(false, true)
		default:
			if i, ok := big.NewInt(-1).SetString(v, 0); ok {
				f.setRangeInt(i, isFromBlock)
			} else {
				f.addError(fmt.Errorf("invalid fromBlock tag: %s", tag))
			}
		}
	case false:
		switch tag {
		case methods.Earliest, methods.Latest, methods.Pending, methods.Safe, methods.Finalized:
			f.toBlock.tag = tag
			f.toBlock.tagSwitch.CompareAndSwap(false, true)
		default:
			if i, ok := big.NewInt(-1).SetString(v, 0); ok {
				f.setRangeInt(i, isFromBlock)
			} else {
				f.addError(fmt.Errorf("invalid toBlock tag: %s", tag))
			}
		}
	}
}

// setRangeInt updates the block range as a numeric value.
func (f *Filter) setRangeInt(i *big.Int, isFromBlock bool) {
	switch isFromBlock {
	case true:
		f.fromBlock.n = big.NewInt(0).Set(i)
		f.fromBlock.tagSwitch.CompareAndSwap(true, false)
	case false:
		f.toBlock.n = big.NewInt(0).Set(i)
		f.toBlock.tagSwitch.CompareAndSwap(true, false)
	}
}

// addError appends an error to the filter's internal error field.
func (f *Filter) addError(err error) {
	f._err = fmt.Errorf("latest filter error: %w", err)
}

func (f *Filter) MarshalJSON() ([]byte, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}
	return easyjson.Marshal(f)
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	if err := f.Validate(); err != nil {
		return err
	}

	return easyjson.Unmarshal(data, f)
}

func (f *Filter) MarshalEasyJSON(w *jwriter.Writer) {
	w.RawByte('{')
	w.RawString(`"fromBlock":`)
	if f.fromBlock.tagSwitch.Load() {
		w.String(f.fromBlock.tag)
	} else {
		w.String(f.fromBlock.n.String())
	}
	w.RawByte(',')
	w.RawString(`"toBlock":`)
	if f.toBlock.tagSwitch.Load() {
		w.String(f.toBlock.tag)
	} else {
		w.String(f.toBlock.n.String())
	}
	if len(f.address.addr) != 0 {
		w.RawByte(',')
		w.RawString(`"address":`)
		if len(f.address.addr) > 1 {
			w.RawByte('[')
			for i, addr := range f.address.addr {
				if i > 0 {
					w.RawByte(',')
				}
				w.String(addr.Hex())
			}
			w.RawByte(']')
		} else if len(f.address.addr) == 1 {
			w.String(f.address.addr[0].Hex())
		}
	}
	if len(f.topics.topics) != 0 {
		w.RawByte(',')
		w.RawString(`"topics":`)
		w.RawByte('[')
		for _, seq := range f.topics.topics {
			if len(seq) > 1 {
				w.RawByte('[')
				for _, topic := range seq {
					w.String(topic.Hex())
					w.RawByte(',')
				}
				w.RawByte(']')
			} else if len(seq) == 1 {
				w.String(seq[0].Hex())
				w.RawByte(',')
			}
		}
		w.RawByte(']')
	}
	w.RawByte('}')
}

func (f *Filter) UnmarshalEasyJSON(w *jlexer.Lexer) {
	w.Delim('{')
	for !w.IsDelim('}') {
		key := w.String()
		w.WantColon()
		switch key {
		case "fromBlock":
			f.setRangeString(w.String(), true)
		case "toBlock":
			f.setRangeString(w.String(), false)
		case "address":
			if w.IsDelim('[') {
				w.Delim('[')
				for !w.IsDelim(']') {
					addr := common.HexToAddress(w.String())
					f.address.addr = append(f.address.addr, addr)
					w.WantComma()
				}
				w.Delim(']')
			} else {
				addr := common.HexToAddress(w.String())
				f.address.addr = append(f.address.addr, addr)
			}
		case "topics":
			w.Delim('[')
			for !w.IsDelim(']') {
				if w.IsDelim('[') {
					var topicGroup []common.Hash
					w.Delim('[')
					for !w.IsDelim(']') {
						topic := common.HexToHash(w.String())
						topicGroup = append(topicGroup, topic)
						w.WantComma()
					}
					w.Delim(']')
					f.topics.topics = append(f.topics.topics, topicGroup)
				} else {
					topic := common.HexToHash(w.String())
					f.topics.topics = append(f.topics.topics, []common.Hash{topic})
				}
			}
			w.Delim(']')
		default:
			w.SkipRecursive()
		}
		w.WantComma()
	}
	w.Delim('}')
}
