package dbmate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDriver_Postgres(t *testing.T) {
	drv, err := GetDriver("postgres")
	require.NoError(t, err)
	_, ok := drv.(PostgresDriver)
	require.Equal(t, true, ok)
}

func TestGetDriver_MySQL(t *testing.T) {
	drv, err := GetDriver("mysql")
	require.NoError(t, err)
	_, ok := drv.(MySQLDriver)
	require.Equal(t, true, ok)
}

func TestGetDriver_Error(t *testing.T) {
	drv, err := GetDriver("foo")
	require.EqualError(t, err, "unsupported driver: foo")
	require.Nil(t, drv)
}
