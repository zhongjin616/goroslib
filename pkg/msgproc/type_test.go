package msgproc

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aler9/goroslib/pkg/msg"
)

type MsgExplicitPackage struct {
	msg.Package `ros:"my_package"`
	Value       uint16
}

type MsgImplicitPackage struct {
	Value uint16
}

func TestType(t *testing.T) {
	for _, ca := range []struct {
		name string
		msg  interface{}
		typ  string
	}{
		{
			"explicit package",
			MsgExplicitPackage{},
			"my_package/MsgExplicitPackage",
		},
		{
			"implicit package",
			MsgImplicitPackage{},
			"goroslib/MsgImplicitPackage",
		},
		{
			"anonymous",
			struct {
				A int
			}{},
			"goroslib/Msg",
		},
	} {
		t.Run(ca.name, func(t *testing.T) {
			typ, err := Type(ca.msg)
			require.NoError(t, err)
			require.Equal(t, ca.typ, typ)
		})
	}
}

func TestTypeErrors(t *testing.T) {
	for _, ca := range []struct {
		name string
		msg  interface{}
		err  string
	}{
		{
			"wrong message type",
			123,
			"message must be a struct",
		},
	} {
		t.Run(ca.name, func(t *testing.T) {
			_, err := Type(ca.msg)
			require.EqualError(t, err, ca.err)
		})
	}
}
