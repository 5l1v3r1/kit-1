package example

import (
	g "github.com/ysmood/gokit"
)

func Example() {
	g.Guard(
		[]string{"go", "run", "./server"},
		g.GuardOptions{
			ExecOpts: g.ExecOptions{
				Prefix: "server | @yellow",
			},
		},
	)
}
