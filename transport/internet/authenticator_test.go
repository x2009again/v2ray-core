package internet_test

import (
	"testing"

	"github.com/v2ray/v2ray-core/testing/assert"
	. "github.com/v2ray/v2ray-core/transport/internet"
	_ "github.com/v2ray/v2ray-core/transport/internet/authenticators/noop"
	_ "github.com/v2ray/v2ray-core/transport/internet/authenticators/srtp"
	_ "github.com/v2ray/v2ray-core/transport/internet/authenticators/utp"
)

func TestAllAuthenticatorLoadable(t *testing.T) {
	assert := assert.On(t)

	noopAuth, err := CreateAuthenticator("none", nil)
	assert.Error(err).IsNil()
	assert.Int(noopAuth.Overhead()).Equals(0)

	srtp, err := CreateAuthenticator("srtp", nil)
	assert.Error(err).IsNil()
	assert.Int(srtp.Overhead()).Equals(4)

	utp, err := CreateAuthenticator("utp", nil)
	assert.Error(err).IsNil()
	assert.Int(utp.Overhead()).Equals(4)
}
