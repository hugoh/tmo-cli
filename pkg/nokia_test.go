package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoginSuccess(t *testing.T) {
	success := &loginResp{
		Success:   0,
		Reason:    0,
		Sid:       "foo",
		CsrfToken: "bar",
	}
	assert.True(t, success.success())
}

func Test_LoginFailure(t *testing.T) {
	fail := &loginResp{
		Success: 0,
		Reason:  600,
	}
	assert.False(t, fail.success())
}
