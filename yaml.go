package yaml

import (
	"encoding/json"

	"sigs.k8s.io/yaml"
)

func jsonOrYamlUnmarshal(b []byte, i interface{}) error {
	if json.Valid(b) {
		return json.Unmarshal(b, i)
	}
	return yaml.UnmarshalStrict(b, i)
}
