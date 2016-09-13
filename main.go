package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/soh335/sliceflag"
)

var (
	directories = sliceflag.String(flag.CommandLine, "dir", nil, "directory for search license")
	carthage    = flag.Bool("carthage", false, "carthage")
	config      = flag.String("config", ".licenser.json", "path to .licenser.json")
)

func main() {
	flag.Parse()
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}

func _main() error {
	m, err := _config()
	if err != nil {
		return err
	}

	licenses := []string{}
	if *carthage {
		_licenses, err := _carthage()
		if err != nil {
			return err
		}
		licenses = append(licenses, _licenses...)
	}
	for _, dir := range *directories {
		_licenses, err := filepath.Glob(filepath.Join(dir, "LICENSE*"))
		if err != nil {
			return err
		}
		licenses = append(licenses, _licenses...)
	}

	var b bytes.Buffer

	for _, license := range licenses {
		f, err := os.Open(license)
		if err != nil {
			return err
		}
		defer f.Close()

		name := filepath.Base(filepath.Dir(license))
		if _, ok := m[name]; ok {
			name = m[name]
		}
		_markdownSection(name, &b, f)
	}

	unsafe := blackfriday.MarkdownCommon(b.Bytes())
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	fmt.Print(string(html))

	return nil
}

func _carthage() ([]string, error) {
	return filepath.Glob(filepath.Join("Carthage", "Checkouts", "*", "LICENSE*"))
}

func _markdownSection(name string, w io.Writer, r io.Reader) {
	fmt.Fprintf(w, "## %s\n\n", name)
	io.Copy(w, r)
	fmt.Fprint(w, "\n\n")
}

func _config() (map[string]string, error) {
	f, err := os.Open(*config)
	if err != nil {
		return nil, nil
	}
	defer f.Close()
	m := map[string]string{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}
