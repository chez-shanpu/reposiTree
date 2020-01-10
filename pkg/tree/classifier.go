package tree

import (
	"errors"
	"log"
	"regexp"
)

var (
	// TODO Available in other languages
	GoSourceRegex = regexp.MustCompile(`.*\.go$`)
	JsSourceRegex = regexp.MustCompile(`.*\.js$`)
	PySourceRegex = regexp.MustCompile(`.*\.py$`)

	ShellRegex      = regexp.MustCompile(`.*\.sh$`)
	MakefileRegex   = regexp.MustCompile(`^Makefile$`)
	DockerfileRegex = regexp.MustCompile(`^Dockerfile$`)
	ConfigfileRegex = regexp.MustCompile(`.*\.(env|cfg)$`)
	StaticfileRegex = regexp.MustCompile(`.*\.(html|css|scss)$`)
	DocumentsRegex  = regexp.MustCompile(`.*\.(md|txt)$`)
	ImageRegex      = regexp.MustCompile(`.*\.(jpeg|jpg|png|svc)$`)
)

func FileClassifier(fileName, language string) (int, error) {
	var sourceRegex *regexp.Regexp

	switch language {
	case "go":
		sourceRegex = GoSourceRegex
	case "javascript":
		sourceRegex = JsSourceRegex
	case "js":
		sourceRegex = JsSourceRegex
	case "python":
		sourceRegex = PySourceRegex
	default:
		return 0, errors.New("argument --language is invalid (go, javascript, python available)")
	}

	switch {
	case sourceRegex.MatchString(fileName): // Matching Sourcefiles
		return 0, nil
	case ShellRegex.MatchString(fileName): // Matching Shellscripts
		return 1, nil
	case MakefileRegex.MatchString(fileName): // Matching Makefiles
		return 2, nil
	case DockerfileRegex.MatchString(fileName): // Matching Dockerfiles
		return 3, nil
	case ConfigfileRegex.MatchString(fileName): // Matching Configfiles
		return 4, nil
	case StaticfileRegex.MatchString(fileName): // Matching Staticfiles
		return 5, nil
	case DocumentsRegex.MatchString(fileName): // Matching Documents
		return 6, nil
	case ImageRegex.MatchString(fileName): // Matching Images
		return 7, nil
	default: // Others
		log.Println(fileName + " is classified as a others.")
		return 8, nil
	}
}
