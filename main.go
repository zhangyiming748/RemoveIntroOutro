package main

import (
	"fmt"
	"github.com/zhangyiming748/RemoveIntroOutro/mediainfo"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var LogLevel string

func init() {
	SetLogLevel("Debug")
}

func main() {
	var (
		perfix string = "00:00:25.558" // 片头时长
		suffix string = "00:00:00.000" // 片尾时长
	)

	folderPath := "/mnt/e/video/Straplez/title1" // 例如： "C:/Users/username/Documents"
	extension := ".mp4"                          // 例如： ".txt"

	files, err := getFilesWithExtension(folderPath, extension)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		newFile := strings.ReplaceAll(file, ".mp4", "_trim.mp4")
		t, err := mediainfo.GetTime(file)
		if err != nil {
			return
		}
		//fmt.Println(t)
		ss := perfix
		to := mediainfo.TimeSub(t, suffix)
		_, width, height := mediainfo.GetCodec(file)
		crf := mediainfo.GetCrfForVP9(width, height)
		slog.Info("获取到的视频信息", slog.String("宽", width), slog.String("高", height), slog.String("crf", crf))
		cmd := exec.Command("ffmpeg", "-i", file, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-ss", ss, "-to", to, newFile)
		slog.Info("cmd", "cmd", cmd.String())
		err = cmd.Run()
		if err != nil {
			return
		} else {
			os.Remove(file)
		}
	}
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

func SetLogLevel(s string) {
	file := "RIO.log"
	logf, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0770)
	if err != nil {
		panic(err)
	}
	var opt slog.HandlerOptions
	switch s {
	case "Debug", "debug":
		LogLevel = "Debug"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info", "info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: false,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
		LogLevel = "Info"
	case "Warn", "warn":
		LogLevel = "Err"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err", "err":
		LogLevel = "Err"
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Debug("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}
