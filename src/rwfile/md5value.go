package rwfile

import (
    "crypto/md5"
    "fmt"
    "io"
    "math"
    "os"
)

const fileChunk = 8192 // 8KB

func countFileMd5(filePath string) string {
    file, err := os.Open(filePath)
    if err != nil {
        return err.Error()
    }
    defer file.Close()

    info, _ := file.Stat()
    fileSize := info.Size()

    blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
    hash := md5.New()

    for i := uint64(0); i < blocks; i++ {
        blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
        buf := make([]byte, blockSize)

        file.Read(buf)
        io.WriteString(hash, string(buf))
    }

    return fmt.Sprintf("%x", hash.Sum(nil))
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println("run with one or more file path")
        fmt.Println("eg: md5sum a.txt b.txt c.iso")
        return
    }
    for i := 1; i < len(os.Args); i++ {
        fmt.Println(countFileMd5(os.Args[i]), os.Args[i])
    }
}