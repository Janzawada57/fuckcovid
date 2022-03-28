package extensions

import (
	"io"

	"github.com/cli/cli/v2/internal/ghrepo"
)

type ExtTemplateType int

const (
	GitTemplateType      ExtTemplateType = 0
	GoBinTemplateType    ExtTemplateType = 1
	OtherBinTemplateType ExtTemplateType = 2
)

//go:generate moq -rm -out extension_mock.go . Extension
type Extension interface {
	Name() string // Extension Name without gh-
	Path() string // Path to executable
	URL() string
	IsLocal() bool
	CurrentVersion() string
	UpdateAvailable() bool
	IsBinary() bool
}

//go:generate moq -rm -out manager_mock.go . ExtensionManager
type ExtensionManager interface {
	List(includeMetadata bool) []Extension
	Install(ghrepo.Interface) error
	InstallLocal(dir string) error
	Upgrade(name string, force bool) error
	Remove(name string) error
	Dispatch(args []string, stdin io.Reader, stdout, stderr io.Writer) (bool, error)
	Create(name string, tmplType ExtTemplateType) error
}