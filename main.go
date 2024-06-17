package main

import (
	"fmt"
	"github.com/zhangyiming748/RemoveIntroOutro/constant"
	"github.com/zhangyiming748/RemoveIntroOutro/mediainfo"
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var LogLevel string

func init() {
	setLog()
	go NumsOfGoroutine()
}

func main() {
	cut := new(constant.Timing)
	cut.SetPrefix("00:02:17.574")
	cut.SetSuffix("00:03:29.168")

	folderPath := "F:\\alist\\电视剧\\甄嬛传\\work" // 例如： "C:/Users/username/Documents"
	extension := ".mp4"                       // 例如： ".txt"

	files, err := getFilesWithExtension(folderPath, extension)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var wg sync.WaitGroup
	ch := make(chan struct{}, constant.CPUS)
	log.Printf("协程缓冲区:%d\n", constant.CPUS)
	for _, file := range files {
		if strings.Contains(file, "trim") {
			log.Printf("跳过已经截取的视频:%s\n", file)
			continue
		}
		ch <- struct{}{}
		go func() {
			wg.Add(1)
			defer func() {
				wg.Done()
				<-ch
			}()
			newFile := strings.ReplaceAll(file, ".mp4", "_trim.mp4")
			t, err := mediainfo.GetTime(file)
			if err != nil {
				return
			}
			//fmt.Println(cut)
			ss := cut.GetPrefix()
			to := mediainfo.TimeSub(t, cut.Suffix)
			_, width, height := mediainfo.GetCodec(file)
			crf := mediainfo.GetCrfForVP9(width, height)
			log.Printf("获取到的视频\t宽:%v\t高:%v\tcrf:%v\n", width, height, crf)
			cmd := exec.Command("ffmpeg", "-i", file, "-ss", ss, "-to", to, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-ac", "1", newFile)
			log.Printf("命令原文%v\n", cmd.String())
			err = cmd.Run()
			if err != nil {
				return
			} else {
				//os.Remove(file)
			}

		}()
	}
	wg.Wait()
}

func getFilesWithExtension(folderPath string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func setLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   "mylog.log",
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
func NumsOfGoroutine() {
	for {
		fmt.Printf("\r当前程序运行时协程个数:%d", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
