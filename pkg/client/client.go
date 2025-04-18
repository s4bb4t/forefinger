package client

import (
	"github.com/ethereum/go-ethereum/rpc"
	"sync"
)

type (
	Client struct {
		max uint8
		idx uint8
		sync.Mutex
		pool []*smart
	}

	smart struct {
		cl *rpc.Client
		sync.Mutex
	}
)

func NewClient(upstream string, velocity uint8) (*Client, error) {
	pool := make([]*smart, velocity)
	for i := range velocity {
		client, err := rpc.Dial(upstream)
		if err != nil {
			return nil, err
		}
		pool[i] = &smart{cl: client}
	}
	return &Client{pool: pool, max: velocity - 1}, nil
}

func (c *Client) Client() (*rpc.Client, func()) {
	return c.client()
}

func (c *Client) client() (*rpc.Client, func()) {
	c.Lock()
	defer c.Unlock()

	for !c.pool[c.idx].TryLock() {
		c.idx++

		if c.idx == c.max {
			c.idx = 0
		}
	}

	releaseFunc := c.pool[c.idx].Unlock

	c.idx++
	return c.pool[c.idx].cl, releaseFunc
}

func (c *Client) Close() {
	c.Lock()
	defer c.Unlock()

	for _, smart := range c.pool {
		smart.Lock()
		smart.cl.Close()
		smart.Unlock()
	}
}
