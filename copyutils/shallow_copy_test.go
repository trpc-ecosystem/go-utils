package copyutils_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"trpc.group/trpc-go/trpc-utils/copyutils"
)

type shallowCopyTestData struct {
	I             int
	unexportedInt int
	IntPtr        *int
}

func TestShallowCopy(t *testing.T) {
	require.NotNil(t, copyutils.ShallowCopy(1, 2), "dst must be a pointer")

	var i int
	require.NotNil(t, copyutils.ShallowCopy(&i, 1.0), "type miss match")
	require.Nil(t, copyutils.ShallowCopy(&i, 2))
	require.Equal(t, 2, i)

	s := shallowCopyTestData{I: 1, unexportedInt: 2, IntPtr: &i}
	var ss shallowCopyTestData
	require.Nil(t, copyutils.ShallowCopy(&ss, &s), "src may or may not be a pointer")
	require.Equal(t, s, ss)
	require.Equal(t, s.unexportedInt, ss.unexportedInt)
	require.Equal(t,
		reflect.ValueOf(s.IntPtr).Pointer(),
		reflect.ValueOf(ss.IntPtr).Pointer())
}
