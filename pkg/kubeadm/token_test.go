package kubeadm

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGenerateKubeadmToken(t *testing.T) {
	token := GenerateKubeadmToken()
	assert.Regexp(t, regexp.MustCompile("[a-z0-9]{6}.[a-z0-9]{16}"), token)
}