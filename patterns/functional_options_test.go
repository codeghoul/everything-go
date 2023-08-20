package patterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewServer_WithCustomOpts(t *testing.T) {

	srv := NewServer(WithMaxConn(1), WithTLS, WithID("custom"))

	assert.Equal(t, &Server{opts: Opts{
		maxConn: 1,
		tls:     true,
		id:      "custom",
	}}, srv)
}

func Test_NewServer_WithDefault(t *testing.T) {

	srv := NewServer()

	assert.Equal(t, &Server{opts: Opts{
		maxConn: 10,
		tls:     false,
		id:      "default",
	}}, srv)
}
