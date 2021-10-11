package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

func main() {
	ignoreList := strings.Split(os.Getenv("JUMPDIR_IGNORE"), ",")

	targetDir := os.Args[1]

	found := false

	if err := godirwalk.Walk(os.Getenv("JUMPDIR_ROOT"), &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			for _, entry := range ignoreList {
				if strings.Contains(osPathname, entry) {
					return godirwalk.SkipThis
				}
			}

			if de.ModeType().IsDir() && strings.Contains(osPathname, targetDir) {
				// only use for direct git repos
				if _, err := os.Stat(osPathname + "/.git"); !os.IsNotExist(err) {
					// /.git exists
					dirs := strings.Split(osPathname, "/")
					if dirs[len(dirs)-1] == targetDir {
						if !found {
							found = true
							fmt.Printf("%s", osPathname)
						}
					}
				}
			}

			return nil
		},
		Unsorted: true,
	}); err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
