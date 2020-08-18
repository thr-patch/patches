package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/thr-patch/patches/tools/pkg/model"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	filesDir := flag.String("patches-path", "./files", "Path to directory containing patch files")
	manifestPath := flag.String("manifest-path", "./manifest.json", "Location of manifest file")
	flag.Parse()

	manifest, err := createManifest(*filesDir)
	if err != nil {
		log.Fatalf("Failed to create manifest: %s", err.Error())
	}
	if err := writeManifest(*manifestPath, manifest); err != nil {
		log.Fatalf("Failed to write manifest: %s", err.Error())
	}
	log.Println("Updated Manifest!")
}

func createManifest(filesDir string) (model.Manifest, error) {
	// merge in any new graphics
	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir (%s): %w", filesDir, err)
	}

	manifest := make(model.Manifest, 0)
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".thrl6p") {
			continue
		}

		p := model.Patch{
			Filename: f.Name(),
		}

		thrl6p, err := getThrl6pData(path.Join(filesDir, f.Name()))
		if err != nil {
			log.Printf("%s has invalid patch data: %s\n", f.Name(), err.Error())
			continue
		}
		p.Tone = thrl6p.Data.Tone
		p.Name = thrl6p.Data.Meta.Name

		p.Metadata, err = getMetadata(path.Join(filesDir, f.Name()))
		if err != nil {
			log.Printf("%s has invalid metadata: %s\n", f.Name(), err.Error())
		}

		manifest = append(manifest, p)
	}

	return manifest, nil
}

func getThrl6pData(path string) (model.THRL6P, error) {
	fileData, err := os.Open(path)
	if err != nil {
		return model.THRL6P{}, fmt.Errorf("failed to open file (%s): %w", path, err)
	}
	defer fileData.Close()

	thrl6p := model.THRL6P{}
	if err := json.NewDecoder(fileData).Decode(&thrl6p); err != nil {
		return thrl6p, err
	}

	return thrl6p, nil
}

func getMetadata(thrl6pFilePath string) (model.Metadata, error) {

	meta := model.Metadata{
		Author:      "Anonymous",
		Description: "NA",
	}

	fileData, err := os.Open(fmt.Sprintf("%s.meta.json", strings.TrimSuffix(thrl6pFilePath, ".thrl6p")))
	if err != nil {
		return meta, fmt.Errorf("failed to open file: %w", err)
	}
	defer fileData.Close()

	if err := json.NewDecoder(fileData).Decode(&meta); err != nil {
		return meta, fmt.Errorf("failed to decide metadata: %w", err)
	}

	return meta, nil
}

func writeManifest(path string, data model.Manifest) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
