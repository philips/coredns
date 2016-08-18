//+build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
)

// A shell script in Go form:
// cd ../../mholt/caddy/caddy
// ./build.bash
// cp ../../mholt/caddy/caddy/caddy coredns

// cd ../../mholt/caddy/caddy
// ./build.bash
func BuildCaddy(dir string) ([]byte, error) {
	// path.Join does not work when adding './'
	bld := exec.Command("./" + build)
	bld.Dir = caddydir
	return bld.CombinedOutput()
}

// cp ../../mholt/caddy/caddy/caddy coredns
func CopyCaddy(dir string) error {
	data, err := ioutil.ReadFile(path.Join(dir, caddy))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(coredns, data, 0755)
}

const (
	// If everything is OK and we are sitting in CoreDNS' dir, this is where caddy should be.
	caddydir = "../../mholt/caddy/caddy/"
	caddy    = "caddy"
	coredns  = "coredns"
	build    = "build.bash"
)

func main() {
	if out, err := BuildCaddy(caddydir); err != nil {
		fmt.Printf("\n%s\n", out)
		log.Fatalf("failed to build caddy: %s", err)

	}
	if err := CopyCaddy(caddydir); err != nil {
		log.Fatalf("failed to copy caddy to coredns: %s", err)
	}
}