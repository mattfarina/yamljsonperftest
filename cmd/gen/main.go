// The gen command builds an example yaml file that's fairly massive to use in
// benchmarking.
//
// Some quick design notes...
// - Instead of using yaml tooling using templates
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/ghodss/yaml"
	"helm.sh/helm/v3/pkg/repo"
)

var header = `apiVersion: v1
entries:
`

var footer = "generated: 2018-04-11T16:56:56.656249201Z"

var chart = `  dummy-chart-{{.Num}}:
`

var release = `  - created: 2017-07-06T01:33:50.952906435Z
    description: Example description
    digest: 249e27501dbfe1bd93d4039b04440f0ff19c707ba720540f391b5aefa3571455
    home: https://example.com
    icon: https://example.com/foo.png
    keywords:
    - A
    - B
    maintainers:
    - email: bar@example.com
      name: Bar
    name: dummy-chart-{{.Num}}
    sources:
    - https://example.com
    - https://example.com
    urls:
    - https://example.com
    version: 1.2.{{.Num2}}
`

type wrapper struct {
	Num  int
	Num2 int
}

func main() {

	os.Mkdir("testdata", 0644)

	// Generate a YAML file for testing
	genYaml()

	// Also generate a json version of the same file content for testing
	genJson()

	fmt.Println("Done generating testing files")
}

func genYaml() {
	fmt.Println("Generating index.yaml for testing")

	f, err := os.Create("./testdata/index.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	charttmpl, err := template.New("chart").Parse(chart)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reltmpl, err := template.New("release").Parse(release)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var w wrapper
	for i := 0; i < 100; i++ {
		w = wrapper{Num: i}
		err = charttmpl.Execute(f, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for j := 0; j < 5000; j++ {
			w.Num2 = j
			err = reltmpl.Execute(f, w)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	_, err = f.WriteString(footer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func genJson() {
	yml, err := os.ReadFile("./testdata/index.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generating json file")

	in := &repo.IndexFile{}
	err = yaml.Unmarshal(yml, &in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generating index.json.yaml for testing")
	out, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = os.WriteFile("./testdata/index.json.yaml", out, 0644)
}
