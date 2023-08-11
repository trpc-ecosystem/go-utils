package copyutils_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"trpc.group/trpc-go/trpc-utils/copyutils"
)

type S struct{}

func TestNew(t *testing.T) {
	i := copyutils.New(3)
	require.Equal(t, reflect.Ptr, reflect.TypeOf(i).Kind())
	require.Equal(t, reflect.Int, reflect.TypeOf(i).Elem().Kind())

	s := copyutils.New(S{})
	require.Equal(t, reflect.Ptr, reflect.TypeOf(s).Kind())
	require.Equal(t, reflect.TypeOf(S{}), reflect.TypeOf(s).Elem())

	s = copyutils.New(&S{})
	require.Equal(t, reflect.Ptr, reflect.TypeOf(s).Kind())
	require.Equal(t, reflect.TypeOf(S{}), reflect.TypeOf(s).Elem())
}
