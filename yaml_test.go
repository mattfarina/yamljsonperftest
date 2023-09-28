package yaml

import (
	"fmt"
	"os"
	"testing"

	"helm.sh/helm/v3/pkg/repo"
	"sigs.k8s.io/yaml"
)

// Loading the index once and caching it is an optimization so we don't reload
// the data for every test.
var index []byte
var jsonindex []byte

func getIndex() []byte {
	if len(index) == 0 {
		var err error
		index, err = os.ReadFile("./testdata/index.yaml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return index
}

func getJSONIndex() []byte {
	if len(jsonindex) == 0 {
		var err error
		jsonindex, err = os.ReadFile("./testdata/index.json.yaml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return jsonindex
}

func BenchmarkYaml(b *testing.B) {
	data := getIndex()
	for n := 0; n < b.N; n++ {
		i := &repo.IndexFile{}
		err := jsonOrYamlUnmarshal(data, i)
		if err != nil {
			b.Errorf("yaml err: %s", err)
		}
	}
}

func BenchmarkJsonThroughYaml(b *testing.B) {
	data := getJSONIndex()
	for n := 0; n < b.N; n++ {
		i := &repo.IndexFile{}
		err := yaml.UnmarshalStrict(data, i)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}

}

func BenchmarkJson(b *testing.B) {
	data := getJSONIndex()
	for n := 0; n < b.N; n++ {
		i := &repo.IndexFile{}
		err := jsonOrYamlUnmarshal(data, i)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}

}
