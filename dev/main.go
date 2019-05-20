package main

import (
	"os"

	g "github.com/ysmood/gokit"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const covPath = "fixtures/profile.cov"

var (
	app       = kingpin.New("dev", "dev tool for gokit")
	cmdTest   = app.Command("test", "run test").Default()
	cmdLab    = app.Command("lab", "run lab")
	cmdBuild  = app.Command("build", "cross build project")
	testMatch = cmdTest.Arg("match", "match test name").String()
	deployTag = cmdBuild.Flag("deploy", "release to github with tag (install hub.github.com first)").Short('d').Bool()
	viewCov   = app.Command("cov", "view html coverage report")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cmdLab.FullCommand():
		lab()

	case cmdTest.FullCommand():
		test()

	case cmdBuild.FullCommand():
		g.E(g.Exec("go", "test", "./..."))
		build(deployTag)

	case viewCov.FullCommand():
		g.E(g.Exec("go", "tool", "cover", "-html="+covPath))
	}
}

func lab() {
	g.Guard([]string{"go", "run", "./dev/lab"})
}

func test() {
	g.Exec("go", "get", "github.com/kyoh86/richgo")

	g.Guard([]string{
		"richgo", "test",
		"-coverprofile=" + covPath,
		"-covermode=count",
		"-run", *testMatch,
	})
}
