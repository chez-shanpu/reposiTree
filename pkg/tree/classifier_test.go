package tree_test

import (
	"github.com/chez-shanpu/reposiTree/pkg/tree"
	"testing"
)

func TestSourceFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.go", ".go"}
	expectRes := 0
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".gohoge", "go", ".go.piyo"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestShellscriptFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.sh", ".sh"}
	expectRes := 1
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".shhoge", "sh", ".sh.piyo"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestMakeFileClassifier(t *testing.T) {
	fileNames := []string{"Makefile"}
	expectRes := 2
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{"hogeMakefile", "Makefilehoge", "makefile"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestDockerFileClassifier(t *testing.T) {
	fileNames := []string{"Dockerfile"}
	expectRes := 3
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{"hogeDockerfile", "Dockerfilehoge", "dockerfile"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestConfigFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.env", "hoge.cfg", ".env", ".cfg"}
	expectRes := 4
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".envhoge", ".cfghoge", "env", "cfg"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestStaticFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.html", "hoge.css", "hoge.scss", ".html", ".css", ".scss"}
	expectRes := 5
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".htmlhoge", ".csshoge", ".scsshoge", "html", "css", "scss"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestDocumentFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.md", "hoge.txt", ".md", ".txt"}
	expectRes := 6
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".mdhoge", ".txthoge", "md", "txt"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}

func TestImageFileClassifier(t *testing.T) {
	fileNames := []string{"hoge.jpeg", "hoge.jpg", "hoge.png", "hoge.svc", ".jpeg", ".jpg", ".png", ".svc"}
	expectRes := 7
	for _, fileName := range fileNames {
		res := tree.FileClassifier(fileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, fileName)
		}
	}

	wrongFileNames := []string{".jpeghoge", ".jpghoge", ".pnghoge", ".svchoge", "jpeg", "jpg", "png", "svc"}
	expectRes = 8
	for _, wrongFileName := range wrongFileNames {
		res := tree.FileClassifier(wrongFileName)
		if res != expectRes {
			t.Errorf("Return: %v Expected: %v FileName: %v", res, expectRes, wrongFileName)
		}
	}
}
