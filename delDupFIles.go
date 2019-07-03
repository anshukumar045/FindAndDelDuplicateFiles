package main

import (
        "strings"
        "os"
        "log"
        "path/filepath"
)

// remove file
/********************************************/
func removeFile(fname string) {
	err := os.Remove(fname)
	if err != nil {
		log.Println("file:= ", fname, " Error in del:- ", err)
	}
}

// Get the last modification time of file
/********************************************/
func getFileModTime(fname string) int64 {
	filestat, err := os.Stat(fname)
	if err != nil {
		log.Println("ERROR to find creation time:- ",err, " file :- ", fname)
	}
return filestat.ModTime().UnixNano()
}

// Walk the dir to get the files
/********************************************/
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
	return nil
	})
return files, err
}

// list all files in a path
/********************************************/
func ListFiles(corpath string, todel bool) {
	var files, final []string
	files, err := FilePathWalkDir(corpath)
		if err != nil {
			log.Fatal(err)
		}
	for _ , fl := range files {
		tmp:= strings.Split(fl, "_")
		tmpp := strings.Split(tmp[0], "/")
		if tmpp[len(tmpp)-1] == "par" {
			if strings.HasSuffix(fl, ".gz")  {
					final = append(final,fl)
			}
		}
	}
		findDupFiles(final,todel)
}
	   
// Generate path based on odate and partition
/********************************************/
func GenPath(odt []string, par []string) []string {
	paths := []string{}
	for _, i := range odt {
		for _, j := range par {
			path := "/home/kanshu/Projects/src/test/cor/" + i + "/sab/" + j
			paths = append(paths, path)
		}
	}
	return paths
}
