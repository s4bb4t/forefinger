package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"math/big"
)

type Int struct {
	n *big.Int
}

func (i *Int) UnmarshalJSON(data []byte) error {
	return easyjson.Unmarshal(data, i)
}

func (i *Int) UnmarshalEasyJSON(w *jlexer.Lexer) {
	i.n = big.NewInt(0)
	i.n.SetString(w.String(), 0)
}

func (c *Client) Call(ctx context.Context, res any, method methods.Method, args ...any) error {
	cl, release := c.client()
	defer release()
	return cl.CallContext(ctx, res, method.Method(), args...)
}

// BatchCall executes len(result) request separated to batches whose size is determined by batchLim
func (c *Client) BatchCall(ctx context.Context, batchLim int, method methods.Method, result *[]any, args [][]any) (error, []error) {
	if batchLim <= 0 {
		return fmt.Errorf("batchLim must be positive"), nil
	}

	batch := make([]rpc.BatchElem, batchLim)
	errs := make([]error, len(*result))

	for i := 0; i < len(*result); i++ {
		batch[i%batchLim] = rpc.BatchElem{
			Method: method.Method(),
			Args:   args[i],
			Result: &(*result)[i],
			Error:  errs[i],
		}

		if i%batchLim == batchLim-1 {
			cl, release := c.client()
			if err := cl.BatchCallContext(ctx, batch); err != nil {
				return err, errs
			}
			release()
		}
	}

	return nil, errs
}

// SequenceBatchCall executes len(result) request separated to batches whose size is determined by batchLim
func (c *Client) SequenceBatchCall(ctx context.Context, batchLim int, sequence *methods.Sequence) (error, []error) {
	if batchLim <= 0 {
		return fmt.Errorf("batchLim must be positive"), nil
	}

	batch := make([]rpc.BatchElem, batchLim)
	errs := make([]error, len(*sequence))

	for i := 0; i < len(*sequence); i++ {
		batch[i%batchLim] = rpc.BatchElem{
			Method: (*sequence)[i].Method.Method(),
			Args:   (*sequence)[i].Args,
			Result: (*sequence)[i].Result,
			Error:  (*sequence)[i].Err,
		}

		if i%batchLim == batchLim-1 {
			cl, release := c.client()
			if err := cl.BatchCallContext(ctx, batch); err != nil {
				return fmt.Errorf("failed to execute batch: %w", err), nil
			}
			release()
		}
	}

	return nil, errs
}
