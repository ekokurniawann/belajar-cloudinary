package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "./../../.env"

func TestLoadConfig(t *testing.T) {
	err := LoadConfig(path)

	require.Nil(t, err)
}
