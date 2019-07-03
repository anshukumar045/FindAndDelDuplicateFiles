package main

import (
        "strings"
        "fmt"
)

// filter the duplicate and delete
/********************************************/
// get the count of data of deleted files
//find the number of dup files

type fileStat struct {
	Name string
	cnt  int
	prt  string
}

func findDupFiles(in []string, todel bool){
	filemap := make(map[string]*fileStat)
	for idx:= 0; idx < len(in); idx++ {
		insplit := strings.Split(in[idx], "_")
		if filemap[insplit[4]] == nil {
				filemap[insplit[4]] = &fileStat{ Name : in[idx],
												cnt: 0,
												prt: strings.Split(insplit[len(insplit)-1],".")[0],
												}
		} else {
			for k,v := range filemap {
				if insplit[4] == k {
					modtime := getFileModTime(v.Name)
					modtimenextfile := getFileModTime(in[idx])
					if modtimenextfile - modtime != 0 {
						if todel == true {
								if modtimenextfile > modtime {
									removeFile(v.Name)
									v.Name = in[idx]
									v.cnt = v.cnt + 1
								} else {
									removeFile(in[idx])
									v.cnt = v.cnt+1
								}
						} else {
							v.Name = in[idx]
							v.cnt = v.cnt + 1
						}
					}
				}
			}
		}
	}
	for k,v := range filemap {
			fmt.Println("hr=",k, " : part=", v.prt, " : dupfiles=", v.cnt)
	}
}
