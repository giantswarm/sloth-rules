package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
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

	outputDir := path.Join(currentDirectory, "helm", "sloth-rules", "files")

	// Cleanup previously generated files.
	files, err := filepath.Glob(path.Join(outputDir, "*.yaml"))
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
							name := strings.TrimSuffix(filepath.Base(slo), ".yaml")

							destinationFileName := fmt.Sprintf("%s-%s-%s-%s.yaml", area.Name(), team.Name(), service.Name(), name)
							destinationPath := path.Join(outputDir, destinationFileName)
							log.Printf("Generating prometheusservicelevel.sloth.slok.dev named %s in %s", destinationFileName, outputDir)

							rawData, err := os.ReadFile(slo)
							if err != nil {
								log.Fatal(err)
							}

							data := map[string]any{}
							err = yaml.Unmarshal(rawData, data)
							if err != nil {
								log.Fatal(err)
							}

							data["area"] = area.Name()
							data["team"] = team.Name()
							data["service"] = service.Name()
							data["slo"] = name

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
func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
