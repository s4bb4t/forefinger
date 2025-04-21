package client

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/s4bb4t/forefinger/pkg/methods"
	"math/big"
	"sync"
)

type (
	Int struct {
		n *big.Int
	}
)

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
func (c *Client) BatchCall(ctx context.Context, batchLim int, method methods.Method, result []any, args [][]any) (error, []error) {
	if batchLim <= 0 {
		return fmt.Errorf("batchLim must be positive"), nil
	}
	var wg sync.WaitGroup
	errs := make([]error, len(result))
	var err error

	for i := 0; i < len(result); i += batchLim {
		wg.Add(1)
		go func() {
			defer wg.Done()
			batchSize := batchLim
			if i+batchSize > len(result) {
				batchSize = len(result) - i
			}
			batch := make([]rpc.BatchElem, batchSize)

			for j := 0; j < batchSize; j++ {
				batch[j] = rpc.BatchElem{
					Method: method.Method(),
					Args:   args[i+j],
					Result: &(result)[i+j],
				}
			}
			cl, release := c.client()
			err = cl.BatchCallContext(ctx, batch)
			for j := 0; j < batchSize; j++ {
				errs[i+j] = batch[j].Error
			}

			release()
		}()
	}

	wg.Wait()
	if err != nil {
		return err, errs
	}

	return nil, errs
}

// TODO: implement batch call for BlockByNumber and so on
//func (c *Client) BatchBlockByNumber(ctx context.Context, batchLim int, method methods.Method, numbers []int64, args [][]any) (*[]models.Block,error, []error) {
//	if batchLim <= 0 {
//		return nil, fmt.Errorf("batchLim must be positive"), nil
//	}
//
//	errs := make([]error, len(numbers))
//
//	for i := 0; i < len(result); i += batchLim {
//		batchSize := batchLim
//		if i+batchSize > len(result) {
//			batchSize = len(result) - i
//		}
//
//		batch := make([]rpc.BatchElem, batchSize)
//
//		for j := 0; j < batchSize; j++ {
//			batch[j] = rpc.BatchElem{
//				Method: method.Method(),
//				Args:   args[i+j],
//				Result: &(result)[i+j],
//			}
//		}
//
//		cl, release := c.client()
//		if err := cl.BatchCallContext(ctx, batch); err != nil {
//			return err, errs
//		}
//
//		for j := 0; j < batchSize; j++ {
//			errs[i+j] = batch[j].Error
//		}
//
//		release()
//	}
//
//	return nil, errs
//}

/*
// BatchCallTyped executes batch requests with typed results
func BatchCallTyped[T any](ctx context.Context,c *Client, batchLim int, method methods.Method, results *[]T, args [][]any) (error, []error) {
	if batchLim <= 0 {
		return fmt.Errorf("batchLim must be positive"), nil
	}

	errs := make([]error, len(*results))

	for i := 0; i < len(*results); i += batchLim {
		batchSize := batchLim
		if i+batchSize > len(*results) {
			batchSize = len(*results) - i
		}

		batch := make([]rpc.BatchElem, batchSize)

		for j := 0; j < batchSize; j++ {
			batch[j] = rpc.BatchElem{
				Method: method.Method(),
				Args:   args[i+j],
				Result: &(*results)[i+j],
			}
		}

		cl, release := c.client()
		if err := cl.BatchCallContext(ctx, batch); err != nil {
			return err, errs
		}

		for j := 0; j < batchSize; j++ {
			errs[i+j] = batch[j].Error
		}

		release()
	}

	return nil, errs
}

// CreateBatchCall создает массив результатов указанного типа и выполняет запросы пакетами
func  CreateBatchCall[T any](ctx context.Context, c *Client, batchLim int, method methods.Method, count int, args [][]any) (*[]T, error, []error) {
	if batchLim <= 0 {
		return nil, fmt.Errorf("batchLim must be positive"), nil
	}

	results := make([]T, count)
	errs := make([]error, count)

	for i := 0; i < count; i += batchLim {
		batchSize := batchLim
		if i+batchSize > count {
			batchSize = count - i
		}

		batch := make([]rpc.BatchElem, batchSize)

		for j := 0; j < batchSize; j++ {
			batch[j] = rpc.BatchElem{
				Method: method.Method(),
				Args:   args[i+j],
				Result: &results[i+j],
			}
		}

		cl, release := c.client()
		if err := cl.BatchCallContext(ctx, batch); err != nil {
			return &results, err, errs
		}

		for j := 0; j < batchSize; j++ {
			errs[i+j] = batch[j].Error
		}

		release()
	}

	return &results, nil, errs
}
*/

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
