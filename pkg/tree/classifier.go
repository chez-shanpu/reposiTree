package tree

import (
	"errors"
	"github.com/h2non/filetype"
	svg "github.com/h2non/go-is-svg"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	MaxFiletype = 8

	TypeOther    = 0
	TypeSource   = 1
	TypeMakefile = 2
	TypeDocker   = 3
	TypeConfig   = 4
	TypeStatic   = 5
	TypeDocument = 6
	TypeImage    = 7
)

var (
	GoSourceRegex         = regexp.MustCompile(`.*\.go$`)
	JsSourceRegex         = regexp.MustCompile(`.*\.js$`)
	PythonSourceRegex     = regexp.MustCompile(`.*\.py$`)
	JavaSourceRegex       = regexp.MustCompile(`.*\.java$`)
	PhpSourceRegex        = regexp.MustCompile(`.*\.php$`)
	CsharpSourceRegex     = regexp.MustCompile(`.*\.cs$`)
	CppSourceRegex        = regexp.MustCompile(`.*\.cpp$`)
	TypescriptSourceRegex = regexp.MustCompile(`.*\.ts$`)
	ShellSourceRegex      = regexp.MustCompile(`.*\.sh$`)
	CSourceRegex          = regexp.MustCompile(`.*\.c$`)
	RubySourceRegex       = regexp.MustCompile(`.*\.rb$`)

	MakefileRegex   = regexp.MustCompile(`^Makefile$`)
	DockerfileRegex = regexp.MustCompile(`^Dockerfile$`)
	ConfigfileRegex = regexp.MustCompile(`.*\.(env|cfg)$`)
	StaticfileRegex = regexp.MustCompile(`.*\.(html|css|scss)$`)
	DocumentsRegex  = regexp.MustCompile(`.*\.(md|txt)$`)
)

func FileClassifier(filePath, language string) (int, error) {
	buf, _ := ioutil.ReadFile(filePath)
	if filetype.IsImage(buf) {
		return TypeImage, nil
	} else if filetype.IsDocument(buf) {
		return TypeDocument, nil
	} else if svg.Is(buf) {
		return TypeImage, nil
	}

	var sourceRegex *regexp.Regexp
	switch strings.ToLower(language) {
	case "go":
		sourceRegex = GoSourceRegex
	case "javascript":
		sourceRegex = JsSourceRegex
	case "js":
		sourceRegex = JsSourceRegex
	case "python":
		sourceRegex = PythonSourceRegex
	case "java":
		sourceRegex = JavaSourceRegex
	case "php":
		sourceRegex = PhpSourceRegex
	case "c#":
		sourceRegex = CsharpSourceRegex
	case "cs":
		sourceRegex = CsharpSourceRegex
	case "cpp":
		sourceRegex = CppSourceRegex
	case "typescript":
		sourceRegex = TypescriptSourceRegex
	case "ts":
		sourceRegex = TypescriptSourceRegex
	case "shell":
		sourceRegex = ShellSourceRegex
	case "c":
		sourceRegex = CSourceRegex
	case "ruby":
		sourceRegex = RubySourceRegex
	case "rb":
		sourceRegex = RubySourceRegex
	default:
		return 0, errors.New("argument --language is invalid (go, javascript, python, java, php, c#, cpp,typescript, shell, c, ruby, rb available)")
	}

	_, fileName := filepath.Split(filePath)
	switch {
	case sourceRegex.MatchString(fileName): // Matching Sourcefiles
		return TypeSource, nil
	case MakefileRegex.MatchString(fileName): // Matching Makefiles
		return TypeMakefile, nil
	case DockerfileRegex.MatchString(fileName): // Matching Dockerfiles
		return TypeDocker, nil
	case ConfigfileRegex.MatchString(fileName): // Matching Configfiles
		return TypeConfig, nil
	case StaticfileRegex.MatchString(fileName): // Matching Staticfiles
		return TypeStatic, nil
	case DocumentsRegex.MatchString(fileName): // Matching Documents
		return TypeDocument, nil
	default: // Others
		log.Println(fileName + " is classified as a others.")
		return TypeOther, nil
	}
}
