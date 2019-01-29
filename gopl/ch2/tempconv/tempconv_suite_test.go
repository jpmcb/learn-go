package tempconv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTempconv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tempconv Suite")
}
