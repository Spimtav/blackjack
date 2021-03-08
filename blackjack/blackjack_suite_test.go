package blackjack_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClarpjack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clarpjack Suite")
}
