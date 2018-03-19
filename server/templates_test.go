package server

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTemplatesCheckSymlinkDirs(t *testing.T) {
	templateDir, err := createTemplates()
	if err != nil {
		t.Fatal(err)
	}

	c := webConfig{}
	_, err = loadTemplates(c, templateDir)

	if err != nil {
		t.Errorf("expected err to not have ocurred, got %q", err)
	}
}

// createTemplates creates a directory structure containing empty templates, a directory and
// a directory symlink
// ├── ..data          -> ..actual
// ├── <template>.html
// └── ..actual
func createTemplates() (string, error) {
	templateDir, err := ioutil.TempDir("", "")
	if err != nil {
		return "", err
	}
	actualDir := join(templateDir, "..actual")
	symlinkDir := join(templateDir, "..data")
	if err = os.Mkdir(actualDir, 0755); err != nil {
		return "", err
	}
	if err = os.Symlink(actualDir, symlinkDir); err != nil {
		return "", err
	}
	for _, template := range requiredTmpls {
		templatePath := join(templateDir, template)
		if err = ioutil.WriteFile(templatePath, []byte(template), 0444); err != nil {
			return "", err
		}
	}
	return templateDir, nil
}
