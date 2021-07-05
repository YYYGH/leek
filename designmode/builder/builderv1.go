package builder

import (
	"errors"
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 9
	defaultMinIdle  = 1
)

// ResourcePoolConfig resource pool config
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// ResourcePoolConfigBuilder resource pool config builder
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func (b *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	b.name = name
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <= 0, input:[%d]", maxTotal)
	}
	b.maxTotal = maxTotal
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max idle cannot < 0, input:[%d]", maxIdle)
	}
	b.maxIdle = maxIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("min idle cannot < 0, input:[%d]", minIdle)
	}
	b.minIdle = minIdle
	return nil
}

func (b *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if b.name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if b.minIdle == 0 {
		b.minIdle = defaultMinIdle
	}
	if b.maxIdle == 0 {
		b.maxIdle = defaultMaxIdle
	}
	if b.maxTotal == 0 {
		b.maxTotal = defaultMaxTotal
	}

	if b.maxTotal < b.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", b.maxTotal, b.maxIdle)
	}
	if b.minIdle > b.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", b.maxIdle, b.minIdle)
	}
	return &ResourcePoolConfig{
		name:     b.name,
		maxTotal: b.maxTotal,
		maxIdle:  b.maxIdle,
		minIdle:  b.minIdle,
	}, nil
}