package base

import (
	"io"
	"os"
)

// 创建文件夹
func creatDir(path string){
	os.Mkdir(path,777)
}

//删除文件夹
func removeDir(path string){
	os.Remove(path)
}

/**
    检查文件是否存在
 */
func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建文件
func creatFile(path string){
	os.Create(path)
}

//文件写入
func write(path string,str string){
	file,_:=os.OpenFile(path,os.O_APPEND,0777)
	io.WriteString(file,str)
}

//删除文件
func removeFile(path string){
	os.Remove(path)
}

func CreatAndWrite(path string,str string){
	a,_ := fileExists(path)
	if !a {
		creatFile(path)
	}

	write(path,str)
}