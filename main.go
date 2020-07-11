package main

import (
	"fmt"
	"io/ioutil"
	"log"
	_ "os"
	"os/exec"
)
func main(){

	fileInfoList ,err:=ioutil.ReadDir("D:\\go-work\\src\\hugo_hk_blog\\content\\post")
	if err!=nil{
		log.Fatal(err)
	}
	for i:=range fileInfoList{
		fmt.Printf("%-2d--%s\n",i,fileInfoList[i].Name())
	}
	num:=0
	name:="未命名"
	fmt.Scanln(&num,&name)
	nameFile:="post/"+fileInfoList[num].Name()+"/"+name+".md"
	cmd:=exec.Command("hugo","new",nameFile)
	err1:=cmd.Run()
	if err1!=nil{
		fmt.Println(err1)
	}


	nameFile2:="content/"+nameFile
	cmd2:=exec.Command("cmd","/c","start",nameFile2)
	fmt.Println("test")
	err2:=cmd2.Run()
	if err2!=nil{
		fmt.Println(err2)
	}

}