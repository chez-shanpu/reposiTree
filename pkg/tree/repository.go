package tree

import (
	"os"
	"path/filepath"
	"strconv"
)

func OutputAbstractRepository(nodes []*Node, outPath string) error {
	var dirName string
	for _, n := range nodes {
		cnt := 0
		for true {
			dirName = getDirName(n, strconv.Itoa(cnt))
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
			err := OutputAbstractRepository(n.ChildNodes, filepath.Join(outPath, dirName))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getDirName(n *Node, suffix string) string {
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
		if v > 0 {
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
