package rwfile

import (
    "fmt"
    "os"
    "path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
    if info == nil {
        // 文件名称超过限定长度等其他问题也会导致info == nil
        // 如果此时return err 就会显示找不到路径，并停止查找。
        println("can't find:(" + path + ")")
        return nil
    }
    if info.IsDir() {
        println("This is folder:(" + path + ")")
        return nil
    } else {
        println("This is file:(" + path + ")")
        return nil
    }
}

func showFileList(root string) {
    err := filepath.Walk(root, walkFunc)
    if err != nil {
        fmt.Printf("filepath.Walk() error: %v\n", err)
    }
    return
}

func main() {
    var path string
    fmt.Println("请输入想要遍历的路径：")
    fmt.Scanf("%s", &path)
    fmt.Println("=====")
    // path = "D:\\精通windows server 2008 R2.pdf"
    showFileList(path)
}