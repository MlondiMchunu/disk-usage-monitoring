package main

import (
	"fmt"
	"os"
	"syscall"
)

func getDiskUsage(path string){
	var stat syscall.Staffs_t
	err := syscall.Staffs_t(path, &stat)
	if err != nil{
		fmt.Println("Error fetching disk usage:",err)
		return
	}
	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	percentUsed := float64(used)/ float64(total) * 100
	fmt.Printf("Disk usage of %s:\n",path)
	fmt.printf("Total: %d GB\n",total/1e9)
	fmt.printf("Free: %d GB\n", free/1e9)
}
func main(){
	path := "/"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err){
		fmt.Printf("Error: '%s' Path doesn't exist.\n",path)
		return
	}else if err != nil{
		fmt.Printf("Error occued while accessing path %s: %v \n",path,err)
		return 
	}
	getDiskUsage(path)
}