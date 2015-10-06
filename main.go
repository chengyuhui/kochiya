package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"text/template"

	"github.com/kardianos/osext"
)

func execToString(tmpl *template.Template, data interface{}) (string, error) {
	var doc bytes.Buffer
	err := tmpl.Execute(&doc, data)
	return doc.String(), err
}

func linkTemp(src, targetBase, tmpdir string) (string, error) {
	ext := filepath.Ext(src)
	dst := filepath.Join(tmpdir, fmt.Sprint(targetBase, ext))
	err := os.Symlink(src, dst)
	return dst, err
}

func moveImages(videos *[]video, tmpdir string) error {
	for i, video := range *videos {
		videoTmp, err := linkTemp(video.VideoPath, strconv.Itoa(i), tmpdir)
		if err != nil {
			return err
		}

		imageTmp, err := linkTemp(video.ImagePath, strconv.Itoa(i), tmpdir)
		if err != nil {
			return err
		}

		video.VideoPath = videoTmp
		video.ImagePath = imageTmp

		(*videos)[i] = video
	}
	return nil
}

func scheduleCleanup(doCleanup func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		doCleanup()
	}()
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Printf("Usage: %s <config.json>", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	log.Println("正在尝试引发奇迹...")
	conf := loadConfig(os.Args[1])
	log.Println("配置文件已加载")

	tmpl, err := template.New("template.avs").ParseFiles("template.avs")
	if err != nil {
		log.Fatal(err)
	}

	avs, err := execToString(tmpl, conf)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "kochiya")
	if err != nil {
		log.Fatal(err)
	}
	cleanup := func() {
		log.Println("正在清理...")
		// Dirty Hack that just works.
		exec.Command("taskkill", "/F", "/IM", "x264_64_tMod-8bit-all.exe").Run()
		err := os.RemoveAll(tmpdir)
		fmt.Println(err)
	}
	defer cleanup()
	scheduleCleanup(cleanup)
	log.Printf("已创建临时文件夹：%s", tmpdir)

	avsPath := filepath.Join(tmpdir, "encode.avs")
	err = ioutil.WriteFile(avsPath, []byte(avs), 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AVS文件已写入 %s", avsPath)

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("BIN=%s", folderPath)

	err = moveImages(&conf.Videos, tmpdir)
	if err != nil {
		log.Fatal(err)
	}
	encode(conf, avsPath, folderPath)
}
