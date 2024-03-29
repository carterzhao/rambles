package namespaces

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestNewNamespaceWithSuffix(t *testing.T) {
	cases := []struct {
		ns string
	}{
		{
			ns: "test",
		},
	}

	api := &KubernetesAPI{
		Suffix: "unit-test",
		Client: testclient.NewSimpleClientset(),
	}

	for _, c := range cases {
		// create the postfixed namespace
		err := api.NewNamespaceWithSuffix(c.ns)
		if err != nil {
			t.Fatal(err.Error())
		}

		_, err = api.Client.CoreV1().Namespaces().Get("test-unit-test1", v1.GetOptions{})

		if err != nil {
			t.Fatal(err.Error())
		}
	}
}
