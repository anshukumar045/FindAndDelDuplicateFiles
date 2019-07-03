/*
******************************************************************************************
 The Script is to find duplicates and delete the duplicate files starting with "par" and ending with ".gz"
 It will retain the latest file
 filename:= par_AAAA_bbbb_cccc_ddddd.gz
******************************************************************************************
*/

package main

import (
        "os"
        "strings"
        "log"
        "sync"
        "fmt"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
	log.Fatal(`            No arguments provided
							Enter value as [odate]in range or single value <partition> as single value or range
							Example 'script del [190306,190307-190309,190310] [0000,0001-0003]'
							Example 'script del [190307] [0000-0003]'
							Example 'script del [190307] [0000]'
							Example 'script del [190307]'
							Example 'script dup [190306,190307-190309,190310] [0000,0001-0003]'
							Example 'script dup [190307] [0000-0003]'
							Example 'script dup [190307] [0000]'
							Example 'script dup [190307]'
							To Test 'script dup test /xxxx/yyyy/zzzz/src/test'
							To test 'script del test /xxxx/yyyy/zzzz/src/test'
							Default for partition is 0000
							Note:- * 'del' will only delete duplicate files starting with "par" and ending with ".gz" and print deleted files names on the screen
									* 'dup' will print duplicate files name on the screen                   `)
	}
	var wg sync.WaitGroup
	if strings.ToUpper(arguments[1]) == "DEL"{
		if len(arguments) < 3 {
				log.Fatal("Please provide complete list of arguments")
		}
		if strings.ToUpper(arguments[2]) == "TEST" {
				testpath := arguments[3]
				fmt.Println()
				fmt.Println("testpath", testpath)
				fmt.Println("***************************************************")
				ListFiles(testpath, true)
				fmt.Println("***************************************************")
		} else {
			odates, parts := getArgs(arguments)
			pathlst := GenPath(odates, parts)
			fmt.Println("***************************************************")
			for _ ,pth := range pathlst {
				go func() {
					defer wg.Done()
					ListFiles(pth,true)
				}()
				wg.Add(1)
				wg.Wait()
			}
			fmt.Println("***************************************************")
		}
		log.Println("Duplicate Files Deletion Completed")
	} else if strings.ToUpper(arguments[1]) == "DUP"{
		if len(arguments) < 3 {
				log.Fatal("Please provide complete list of arguments")
		}
		if strings.ToUpper(arguments[2]) == "TEST" {
			testpath := arguments[3]
			fmt.Println()
			fmt.Println("testpath", testpath)
			fmt.Println("***************************************************")
			ListFiles(testpath,false)
			fmt.Println("***************************************************")
		} else {
			odates, parts := getArgs(arguments)
			pathlst := GenPath(odates, parts)
			fmt.Println("***************************************************")
			for _ ,pth := range pathlst {
					go func() {
							defer wg.Done()
							ListFiles(pth,false)
					}()
					wg.Add(1)
					wg.Wait()
			}
			fmt.Println("***************************************************")
		}
		log.Println("Duplicate Search Completed")
	}

}

func getArgs(args []string)([]string, []string) {
	var odates, parts []string
	var err error
	if len(args) < 4 {
		parts = append(parts,"0000")
	} else {
		partition := args[3]
		parts, err = PrseArgs(partition,true)
    if err != nil {
      log.Fatal("Wrong format for partition", err)
    }
	}
	odate := args[2]
	odates, err = PrseArgs(odate,false)
	if err != nil {
					log.Fatal("Wrong format for odates", err)
	}
	return odates,parts
}
