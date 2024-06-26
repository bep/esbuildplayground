package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"src/react1.jsx", "src/react2.jsx"},
		Bundle:      true,
		Splitting:   true,
		Outdir:      "out",
		Format:      api.FormatESModule,
		Write:       true,
	})

	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			os.Stderr.WriteString(err.Text + "\n")
		}
		os.Exit(1)
	}

	os.Exit(0)
}
