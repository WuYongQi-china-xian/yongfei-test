package test_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	"os"
)

var _ = Describe("Test", func() {
	It("test", func() {
		fmt.Println("-----------")
		fmt.Println(os.Getenv("KUBECONFIG"))
		fmt.Println("----------")
	})
})
