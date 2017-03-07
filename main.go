package main

import (
	"os"
	"bufio"
	"io"
	"strings"
	"io/ioutil"
	"fmt"
)



func main() {
	fmt.Println("file process!")
	var arrayList []string
	skillfolder := "./"
	// 获取所有文件
	files, _ := ioutil.ReadDir(skillfolder)
	for _,file := range files {
		if file.IsDir() {
			continue
		} else {
			if strings.HasPrefix(file.Name(), "INFO-20170") {
				//fmt.Println(file.Name())
				arrayList=append(arrayList,file.Name())
			}
			//fmt.Println(file.Name())
		}
	}
	f, err := os.OpenFile("./showPVProcess-guangdong", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Sprintf("fileCreateAndWrite os openFile error! fileName: %s err: %s \n", "./showPVProcess-guangdong", err.Error())
		return
	}
	for _, value := range arrayList {
		var (
			readAll     = false
			readByte    []byte
			line        []byte
			err         error
			content string
			contentArray []string
		)
		fs, err := os.Open(fmt.Sprintf("./%s",value))
		if err != nil {
			fmt.Printf("readFileByLine open error! filePath: %s err: %s \n", fmt.Sprintf("./%s",value), err.Error())
			return
		}
		buf := bufio.NewReader(fs)
		for err != io.EOF {
			if err != nil {
				fmt.Sprintf("readFileByLine read error! err: %s \n", err.Error())
			}
			if readAll {
				readByte, readAll, err = buf.ReadLine()
				line = append(line, readByte...)
			} else {
				readByte, readAll, err = buf.ReadLine()
				line = append(line, readByte...)
				if len(strings.TrimSpace(string(line))) == 0 {
					continue
				}
				content=string(line)
				line = line[:0]
				if strings.Index(content, "show success")>0 {
					contentArray=strings.Split(content,"referer: ")
					if len(contentArray) > 1 {
						_, err = f.Write([]byte(fmt.Sprintf("%s\n",contentArray[1])))
						if err != nil {
							fmt.Sprintf("fileCreateAndWrite write error! content: %v fileName: %s err: %s \n", content, fmt.Sprintf("./%s",value), err.Error())
							//return
						}
					}
				}
			}
		}
		fs.Close()
	}
	f.Close()
	fmt.Println("successfully!")
}

/**
写文件
创建人:邵炜
创建时间:2016年9月7日16:31:39
输入参数:文件内容 写入文件的路劲(包含文件名)
输出参数:错误对象
*/
func fileCreateAndWrite(content *[]byte, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		//glog.Error("fileCreateAndWrite os openFile error! fileName: %s err: %s \n", fileName, err.Error())
		return err
	}
	defer f.Close()
	_, err = f.Write(*content)
	if err != nil {
		//go glog.Error("fileCreateAndWrite write error! content: %v fileName: %s err: %s \n", *content, fileName, err.Error())
		return err
	}
	//glog.Info("fileCreateAndWrite run success! fileName: %s content: %s  \n", fileName, string(*content))
	return nil
}

/**
文件读取逐行进行读取
创建人:邵炜
创建时间:2016年9月20日10:23:41
输入参数: 文件路劲
输出参数: 字符串数组(数组每一项对应文件的每一行) 错误对象
*/
func readFileByLine(filePath string) (*[]string, error) {
	var (
		readAll     = false
		readByte    []byte
		line        []byte
		err         error
		contentLine []string
	)
	fs, err := os.Open(filePath)
	if err != nil {
		//fmt.Println("readFileByLine open error! filePath: %s err: %s \n", filePath, err.Error())
		return nil, err
	}
	defer fs.Close()
	buf := bufio.NewReader(fs)
	for err != io.EOF {
		if err != nil {
			//glog.Error("readFileByLine read error! err: %s \n", err.Error())
		}
		if readAll {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
		} else {
			readByte, readAll, err = buf.ReadLine()
			line = append(line, readByte...)
			if len(strings.TrimSpace(string(line))) == 0 {
				continue
			}
			contentLine = append(contentLine, string(line))
			line = line[:0]
		}
	}
	//glog.Info("readFileByLine run success! filePath: %s fileContent: %v \n", filePath, contentLine)
	return &contentLine, nil
}
