package controllers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	GinkgoWriter.Println("Setting up Controllers test suite...")
})

var _ = AfterSuite(func() {
	GinkgoWriter.Println("Tearing down Controllers test suite...")
})