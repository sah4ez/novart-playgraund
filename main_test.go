package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {
	testCases := []struct {
		name   string
		price1 float64
		price2 float64
		exp    bool
	}{
		{
			"10%%",
			100.0,
			90.0,
			true,
		},
		{
			"0.1%",
			100.0,
			99.9,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := compare(tc.price1, tc.price2)
			require.Equal(t, tc.exp, act)
		})
	}
}
