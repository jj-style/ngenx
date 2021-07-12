package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"gopkg.in/yaml.v3"
	_ "embed"

	. "github.com/jj-style/ngenx/src"
)

var (
	inputFile  string
	outputFile string

	//go:embed nginx.conf.tmpl
	nginxTemplate string
)

const (
	defaultSpecFile = "spec.yaml"
)

func init() {
	flag.StringVar(&inputFile, "input", defaultSpecFile, "file to read nginx config specification from. Use - for stdin")
	flag.StringVar(&outputFile, "output", "", "file to output generated nginx config to. Leave empty for stdout")
	flag.Parse()
}

func readFromFileOrStdin(filename string) (string, error) {
	var config string
	if filename == "-" {
		lines := make([]string, 0, 100)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		config = strings.Join(lines, "\n")
	} else {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}
		config = string(bytes)
	}
	return config, nil
}

func main() {
	// load the config
	cfgString, err := readFromFileOrStdin(inputFile)
	if err != nil {
		log.Fatal(fmt.Errorf("reading config: %v", err))
	}
	var config Config
	err = yaml.Unmarshal([]byte(cfgString), &config)
	if err != nil {
		log.Fatal(fmt.Errorf("parsing config: %v", err))
	}
	config.Prepare()

	// prepare the templater
	//nginxTmpl, err := ioutil.ReadFile("nginx.conf.tmpl")
	//if err != nil {
	//	log.Fatal(fmt.Errorf("reading template file: %v", err))
	//}
	tmpl, err := template.New("nginx").Parse(nginxTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("parsing template: %v", err))
	}

	// prepare where the templated config will be written to
	var writer io.Writer
	if outputFile == "" {
		writer = os.Stdout
	} else {
		writer, err = os.OpenFile(outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			log.Fatal(fmt.Errorf("opening output file: %v", err))
		}
	}

	// template!
	err = tmpl.Execute(writer, config)
	if err != nil {
		log.Fatal(fmt.Errorf("executng templating: %v", err))
	}
}
