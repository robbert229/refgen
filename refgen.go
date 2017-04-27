package refgen

//go:generate go-bindata -pkg refgen -o template.go ./refgen.tmpl

import (
	"path"
	"os/exec"
	"github.com/pkg/errors"
	"os"
	"text/template"
	"io"
	"strings"
	"encoding/json"
)

var (
	t *template.Template
)

func init() {
	t = template.Must(template.New("template").Parse(string(MustAsset("refgen.tmpl"))))
}

// Parameter represents a structs parameter.
type Parameter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func buildPayload(pack, name string, writer io.Writer) error {
	folders := strings.Split(pack, "/")
	lastPath := folders[len(folders) - 1]

	payload := struct {
		Package string
		ShortPackage string
		Model string
	}{
		pack,
		lastPath,
		name,
	}
	return errors.Wrap(t.Execute(writer, payload), "unable to execute template")
}

// Keys returns the list of Parameters on a given struct. It uses the go compiler
func Keys(pack, name string) ([]Parameter, error) {
	folder := path.Join(os.TempDir(), "refgen")
	if err := os.Mkdir(folder, 0755); err != nil {
		if !os.IsExist(err) {
			return nil, errors.Wrap(err, "unable to get temporary directory")
		}
	}

	mainFilePath := path.Join(folder, "main.go")
	file, err := os.OpenFile(mainFilePath, os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0755)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file")
	}

	if err := buildPayload(pack, name, file); err != nil {
		return nil, errors.Wrap(err, "unable to write payload")
	}

	cmd := exec.Command("go","get","github.com/fatih/structs")
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "unable to install structs dependency")
	}

	cmd = exec.Command("go", "run", mainFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, "unable to run refgen")
	}

	var fields []Parameter
	if err := json.Unmarshal(output, &fields); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal")
	}

	if err := os.RemoveAll(folder); err != nil {
		return nil, errors.Wrap(err, "unable to delete folder")
	}

	return fields, nil
}