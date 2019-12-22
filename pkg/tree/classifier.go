package tree

import (
	"log"
	"regexp"
)

var (
	// TODO Available in other languages
	SourceRegex = regexp.MustCompile(`.*\.go$`)

	ShellRegex      = regexp.MustCompile(`.*\.sh$`)
	MakefileRegex   = regexp.MustCompile(`^Makefile$`)
	DockerfileRegex = regexp.MustCompile(`^Dockerfile$`)
	ConfigfileRegex = regexp.MustCompile(`.*\.(env|cfg)$`)
	StaticfileRegex = regexp.MustCompile(`.*\.(html|css|scss)$`)
	DocumentsRegex  = regexp.MustCompile(`.*\.(md|txt)$`)
	ImageRegex      = regexp.MustCompile(`.*\.(jpeg|jpg|png|svc)$`)
)

func FileClassifier(fileName string) int {
	switch {
	case SourceRegex.MatchString(fileName): // Matching Sourcefiles
		return 0
	case ShellRegex.MatchString(fileName): // Matching Shellscripts
		return 1
	case MakefileRegex.MatchString(fileName): // Matching Makefiles
		return 2
	case DockerfileRegex.MatchString(fileName): // Matching Dockerfiles
		return 3
	case ConfigfileRegex.MatchString(fileName): // Matching Configfiles
		return 4
	case StaticfileRegex.MatchString(fileName): // Matching Staticfiles
		return 5
	case DocumentsRegex.MatchString(fileName): // Matching Documents
		return 6
	case ImageRegex.MatchString(fileName): // Matching Images
		return 7
	default: // Others
		log.Println(fileName + " is classified as a others.")
		return 8
	}
}
