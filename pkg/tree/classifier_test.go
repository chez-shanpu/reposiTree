package tree_test

import (
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"testing"
)

func TestSourceFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.go", ".go"}
	expectRes := tree.TypeSource
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".gohoge", "go", ".go.piyo"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestMakeFileClassifier(t *testing.T) {
	fileNames := []string{"Makefile"}
	expectRes := tree.TypeMakefile
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{"hogeMakefile", "Makefilehoge", "makefile"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestDockerFileClassifier(t *testing.T) {
	fileNames := []string{"Dockerfile"}
	expectRes := tree.TypeDocker
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{"hogeDockerfile", "Dockerfilehoge", "dockerfile"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestConfigFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.env", "hoge.cfg", ".env", ".cfg"}
	expectRes := tree.TypeConfig
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".envhoge", ".cfghoge", "env", "cfg"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestStaticFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.html", "hoge.css", "hoge.scss", ".html", ".css", ".scss"}
	expectRes := tree.TypeStatic
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".htmlhoge", ".csshoge", ".scsshoge", "html", "css", "scss"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestDocumentFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.md", "hoge.txt", ".md", ".txt"}
	expectRes := tree.TypeDocument
	for _, fileName := range fileNames {
		res, _ := tree.FileClassifier(fileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".mdhoge", ".txthoge", "md", "txt"}
	expectRes = tree.TypeOther
	for _, wrongFileName := range wrongFileNames {
		res, _ := tree.FileClassifier(wrongFileName, "go")
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}
