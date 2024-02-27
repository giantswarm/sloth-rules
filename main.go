package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("prometheusservicelevel.sloth.slok.dev.yaml").ParseFiles(path.Join(currentDirectory, "templates", "prometheusservicelevel.sloth.slok.dev.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	outputDir := path.Join(currentDirectory, "helm", "sloth-rules", "templates")

	// Cleanup previously generated files.
	files, err := filepath.Glob(path.Join(outputDir, "./*.yaml"))
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Fatal(err)
		}
	}

	areasDir := path.Join(currentDirectory, "areas")

	areas, err := os.ReadDir(areasDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range areas {
		if area.IsDir() {
			log.Printf("Area: %s", area.Name())

			teamsDir := path.Join(areasDir, area.Name(), "teams")
			teams, err := os.ReadDir(teamsDir)
			if err != nil {
				log.Fatal(err)
			}
			for _, team := range teams {
				if team.IsDir() {
					log.Printf("Team: %s", team.Name())

					servicesDir := path.Join(teamsDir, team.Name())
					services, err := os.ReadDir(servicesDir)
					if err != nil {
						log.Fatal(err)
					}
					for _, service := range services {
						log.Printf("Alert: %s", service.Name())
						slos, err := filepath.Glob(path.Join(servicesDir, service.Name(), "slos", "*.yaml"))
						if err != nil {
							log.Fatal(err)
						}

						for _, slo := range slos {
							rawData, err := os.ReadFile(slo)

							if err != nil {
								log.Fatal(err)
							}

							data := map[string]any{}
							err = yaml.Unmarshal(rawData, data)
							if err != nil {
								log.Fatal(err)
							}

							name := strings.TrimSuffix(filepath.Base(slo), ".yaml")

							var provider = "all"
							providerStr, found := data["provider"]
							if found {
								provider = providerStr.(string)
							}

							destinationFileName := fmt.Sprintf("%s-%s-%s-%s.yaml", area.Name(), team.Name(), service.Name(), name)
							destinationPath := path.Join(outputDir, destinationFileName)
							log.Printf("Generating prometheusservicelevel.sloth.slok.dev named %s in %s", destinationFileName, outputDir)

							dir := path.Dir(destinationPath)
							if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
								err := os.Mkdir(dir, os.ModePerm)
								if err != nil {
									log.Fatal(err)
								}
							}

							data["area"] = area.Name()
							data["team"] = team.Name()
							data["service"] = service.Name()
							data["slo"] = name
							if provider != "all" {
								data["provider"] = provider
							}
							f, err := os.Create(destinationPath)
							if err != nil {
								log.Fatal(err)
							}
							defer f.Close()

							w := bufio.NewWriter(f)
							err = tmpl.Execute(w, data)
							if err != nil {
								log.Fatal(err)
							}
							w.Flush()
						}
					}
				}
			}
		}
	}
}
