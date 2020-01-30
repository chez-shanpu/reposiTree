package tree

import (
	"os"
	"path/filepath"
	"strconv"
)

func OutputAbstractRepository(nodes []*Node, outPath string, thr float64) error {
	var dirName string
	for _, n := range nodes {
		cnt := 0
		for true {
			dirName = getDirName(n, strconv.Itoa(cnt), thr)
			if _, err := os.Stat(filepath.Join(outPath, dirName)); err == nil {
				cnt++
				continue
			} else if os.IsNotExist(err) {
				err := os.Mkdir(filepath.Join(outPath, dirName), 0755)
				if err != nil {
					return err
				}
				break
			} else {
				return err
			}
		}
		if n.ChildNodes != nil {
			err := OutputAbstractRepository(n.ChildNodes, filepath.Join(outPath, dirName), thr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getDirName(n *Node, suffix string, thr float64) string {
	name := ""
	nameList := map[int]string{
		TypeOther:    "others",
		TypeSource:   "src",
		TypeBuild:    "build",
		TypeConfig:   "config",
		TypeStatic:   "static",
		TypeDocument: "doc",
		TypeImage:    "image",
	}
	for i, v := range n.Vector {
		if v >= thr {
			if name != "" {
				name += "_"
			}
			name += nameList[i]
		}
	}
	if name == "" {
		name += "empty"
	}
	name += "_" + suffix
	return name
}
