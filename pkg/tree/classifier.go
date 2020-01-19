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
	MaxFiletype = 7

	TypeOther    = 0
	TypeSource   = 1
	TypeBuild    = 2
	TypeConfig   = 3
	TypeStatic   = 4
	TypeDocument = 5
	TypeImage    = 6
)

// see https://github.com/github/linguist/blob/master/lib/linguist/languages.yml
var (
	GoSourceRegex         = regexp.MustCompile(`.*\.go$`)
	JsSourceRegex         = regexp.MustCompile(`.*\.(js|jsx|_js|bones|cjs|es|es6|frag|gs|jake|jsb|jscad|jsfl|jsm|jss|mjs|njs|pac|sjs|ssjs|xsjs|xsjslib|js\.erb|vue|)$`)
	PythonSourceRegex     = regexp.MustCompile(`.*\.(py|pyx|pxd|pxi|numpy|numpyw|numsc|bzl|cgi|fcgi|gyp|gypi|lmi|py3|pyde|pyi|pyp|pyt|pyw|rpy|smk|spec|tac|wsgi|xpy|pytb|sage|sagews)$`)
	JavaSourceRegex       = regexp.MustCompile(`.*\.(java|properties|jsp)$`)
	PhpSourceRegex        = regexp.MustCompile(`.*\.(php|hack|hh|hhi|aw|ctp|fcgi|inc|php3|php4|php5|phps|php_cs|php_cs\.dist|zig)$`)
	CsharpSourceRegex     = regexp.MustCompile(`.*\.(cs|cake|csx)$`)
	CppSourceRegex        = regexp.MustCompile(`.*\.(cpp|c\+\+|cc|cp|cxx|h|h\+\+|hh|hpp|hxx|inc|inl|ino|ipp|re|tcc|tpp)$`)
	TypescriptSourceRegex = regexp.MustCompile(`.*\.(ts|tsx)$`)
	ShellSourceRegex      = regexp.MustCompile(`(.*\.(sh|ebuild|eclass|ps1|psd1|psm1|bash|bats|cgi|command|fcgi|ksh|sh\.in|tmux|tool|zsh|tcsh|csh)$)|(^(\.(bash_aliases|bash_history|bash_logout|bash_profile|bashrc|cshrc|login|profile|zlogin|zlogout|zprofile|zshenv|zshrc)|9fs|PKGBUILD|bash_aliases|bash_logout|bash_profile|bashrc|cshrc|gradlew|login|man|profile|zlogin|zlogout|zprofile|zshenv|zshrc)$)`)
	CSourceRegex          = regexp.MustCompile(`.*\.(c|cats|h|idc)$`)
	RubySourceRegex       = regexp.MustCompile(`(.*\.(rb|builder|eye|fcgi|gemspec|god|jbuilder|mspec|pluginspec|podspec|rabl|rake|rbi|rbuild|rbw|rbx|ru|ruby|spec|thor|watchr)$)|(^(\.irbrc|\.pryrc|Appraisals|Berksfile|Brewfile|Buildfile|Capfile|Dangerfile|Deliverfile|Fastfile|Gemfile|Gemfile\.lock|Guardfile|Jarfile|Mavenfile|Podfile|Puppetfile|Rakefile|Snapfile|Thorfile|Vagrantfile|buildfile)$)`)

	BuildfileRegex  = regexp.MustCompile(`(.*\.(mak|d|make|mk|mkfile|dockerfile)$)|(^(Makefile|BSDmakefile|GNUmakefile|Kbuild|Makefile\.am|Makefile\.boot|Makefile\.frag|Makefile\.in|Makefile\.inc|Makefile\.wat|makefile|makefile\.sco|mkfile|Dockerfile)$)`)
	ConfigfileRegex = regexp.MustCompile(`.*\.(env|cfg)$`)
	StaticfileRegex = regexp.MustCompile(`.*\.(html|htm|html\.hl|inc|st|xht|xhtml|jinja|jinja2|mustache|njk|ecr|eex|erb|erb\.deface|phtml|cshtml|razor|haml|haml\.deface|handlebars|hbs|kit|latte|liquid|mtml|marko|jade|pug|rhtml|scaml|slim|svelte|twig|css|less|mss|pcss|postcss|sass|styl|sss|scss)$`)
	DocumentsRegex  = regexp.MustCompile(`.*\.(md|markdown|mdown|mdwn|mdx|mkd|mkdn|mkdown|ronn|wrokbook|rmd|txt)$`)
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
	case "csharp":
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
	case BuildfileRegex.MatchString(fileName): // Matching Buidldfiles
		return TypeBuild, nil
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
