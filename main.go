package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/rakyll/statik/fs"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	_ "statik"
	"strings"
)

func pack(file *os.File, header []byte, tail []byte) {
	// setting buffer
	fi, _ := file.Stat()
	data := make([]byte, fi.Size()-76)
	file.ReadAt(data, 44)
	buffer := bytes.NewBuffer(make([]byte, 0, fi.Size()-60))
	buffer.Write(header)
	buffer.Write(data)
	buffer.Write(tail)

	// filling file
	file.Truncate(fi.Size() - 16)
	file.Seek(0, 0)
	file.Write(buffer.Bytes())
}

func encode(file string, rate string, sample string) {
	// getting ffmpeg from vendor
	statik, _ := fs.New()
	ffmpeg, _ := statik.Open("/ffmpeg.exe")
	var buf bytes.Buffer
	io.Copy(&buf, ffmpeg)
	defer ffmpeg.Close()
	tmp, _ := os.CreateTemp("", "ffmpeg-*.exe")
	defer os.Remove(tmp.Name())
	tmp.Write(buf.Bytes())
	os.Chmod(tmp.Name(), 0755)
	tmp.Close()
	// executing
	cmd := exec.Command(tmp.Name(), "-loglevel", "quiet", "-y", "-i", "file:"+sample, "-ar", rate, "-f", "wav", "-c:a", "adpcm_ms", "file:"+file)
	cmd.Run()
}

func getSampleRate(file string) string {
	// getting ffprobe from vendor
	statik, _ := fs.New()
	ffmpeg, _ := statik.Open("/ffprobe.exe")
	var buf bytes.Buffer
	io.Copy(&buf, ffmpeg)
	defer ffmpeg.Close()
	tmp, _ := os.CreateTemp("", "ffprobe-*.exe")
	defer os.Remove(tmp.Name())
	tmp.Write(buf.Bytes())
	os.Chmod(tmp.Name(), 0755)
	tmp.Close()
	// executing
	sampleRate, _ := exec.Command(tmp.Name(), "-loglevel", "quiet", "-v", "error", "-select_streams", "a:0", "-show_entries", "stream=sample_rate", "-of", "default=noprint_wrappers=1:nokey=1", "file:"+file).Output()
	return strings.TrimSpace(string(sampleRate))
}

func main() {
	custPath := ""
	origPath := ""
	outPath := ""
	requirements := false
	reader := bufio.NewReader(os.Stdin)
	// frontend
	if len(os.Args) > 3 {
		flag.StringVar(&custPath, "c", "", "Custom audio file")
		flag.StringVar(&origPath, "i", "", "Original audio file")
		flag.StringVar(&outPath, "o", "", "Output audio file")

		if flag.CommandLine.Parse(os.Args[1:]) != nil {
			requirements = true
		}
	} else if len(os.Args) == 1 {
		fmt.Println(string([]byte{0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x5f, 0x5f, 0x20, 0x20, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x20, 0x5c, 0x20, 0x2f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x5f, 0x5f, 0x20, 0x5c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x5c, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x20, 0x5c, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x20, 0x5c, 0x20, 0x2f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x20, 0x20, 0x5c, 0x2f, 0x20, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x7c, 0x20, 0x7c, 0x5f, 0x5f, 0x29, 0x20, 0x7c, 0x20, 0x28, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x29, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x20, 0x5c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x5f, 0x5f, 0x29, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x5c, 0x20, 0x20, 0x2f, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x2f, 0x20, 0x5c, 0x5f, 0x5f, 0x5f, 0x20, 0x5c, 0x20, 0x20, 0x2f, 0x20, 0x2f, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x2f, 0x5c, 0x20, 0x5c, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x2f, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x5c, 0x2f, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x29, 0x20, 0x7c, 0x2f, 0x20, 0x2f, 0x5f, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x5c, 0x7c, 0x20, 0x7c, 0x5f, 0x5f, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x2f, 0x7c, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x20, 0x2f, 0x5f, 0x2f, 0x20, 0x20, 0x20, 0x20, 0x5c, 0x5f, 0x5c, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x2f, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5c, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x20, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x5c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x28, 0x5f, 0x29, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x20, 0x5c, 0x20, 0x20, 0x5f, 0x20, 0x20, 0x20, 0x5f, 0x20, 0x20, 0x5f, 0x5f, 0x7c, 0x20, 0x7c, 0x5f, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x5f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x20, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x5f, 0x20, 0x5f, 0x5f, 0x7c, 0x20, 0x7c, 0x5f, 0x20, 0x5f, 0x5f, 0x5f, 0x20, 0x5f, 0x20, 0x5f, 0x5f, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x2f, 0x5c, 0x20, 0x5c, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x2f, 0x20, 0x5f, 0x60, 0x20, 0x7c, 0x20, 0x7c, 0x2f, 0x20, 0x5f, 0x20, 0x5c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x20, 0x2f, 0x20, 0x5f, 0x20, 0x5c, 0x7c, 0x20, 0x27, 0x5f, 0x20, 0x5c, 0x20, 0x5c, 0x20, 0x2f, 0x20, 0x2f, 0x20, 0x5f, 0x20, 0x5c, 0x20, 0x27, 0x5f, 0x5f, 0x7c, 0x20, 0x5f, 0x5f, 0x2f, 0x20, 0x5f, 0x20, 0x5c, 0x20, 0x27, 0x5f, 0x5f, 0x7c, 0x0a, 0x20, 0x20, 0x2f, 0x20, 0x5f, 0x5f, 0x5f, 0x5f, 0x20, 0x5c, 0x20, 0x7c, 0x5f, 0x7c, 0x20, 0x7c, 0x20, 0x28, 0x5f, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x28, 0x5f, 0x29, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x5f, 0x5f, 0x5f, 0x7c, 0x20, 0x28, 0x5f, 0x29, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x7c, 0x20, 0x5c, 0x20, 0x56, 0x20, 0x2f, 0x20, 0x20, 0x5f, 0x5f, 0x2f, 0x20, 0x7c, 0x20, 0x20, 0x7c, 0x20, 0x7c, 0x7c, 0x20, 0x20, 0x5f, 0x5f, 0x2f, 0x20, 0x7c, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x2f, 0x5f, 0x2f, 0x20, 0x20, 0x20, 0x20, 0x5c, 0x5f, 0x5c, 0x5f, 0x5f, 0x2c, 0x5f, 0x7c, 0x5c, 0x5f, 0x5f, 0x2c, 0x5f, 0x7c, 0x5f, 0x7c, 0x5c, 0x5f, 0x5f, 0x5f, 0x2f, 0x20, 0x20, 0x20, 0x5c, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5c, 0x5f, 0x5f, 0x5f, 0x2f, 0x7c, 0x5f, 0x7c, 0x20, 0x7c, 0x5f, 0x7c, 0x5c, 0x5f, 0x2f, 0x20, 0x5c, 0x5f, 0x5f, 0x5f, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x5c, 0x5f, 0x5f, 0x5c, 0x5f, 0x5f, 0x5f, 0x7c, 0x5f, 0x7c, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20}))
		fmt.Print("1. Drag your custom audio file here and press Enter.\n")
		custPath, _ = reader.ReadString('\n')
		fmt.Print("2. Drag the original audio file here (ex: sound_011.wav) and press Enter.\n")
		origPath, _ = reader.ReadString('\n')
		outPath = "output.wav"
		requirements = true
	}
	if len(custPath) != 0 && len(origPath) != 0 && len(outPath) != 0 {
		requirements = true
	}

	if requirements {
		// validation
		custPath, _ = filepath.Abs(strings.TrimSpace(custPath))
		origPath, _ = filepath.Abs(strings.TrimSpace(origPath))
		outPath, _ = filepath.Abs(strings.TrimSpace(outPath))
		custPath = filepath.ToSlash(custPath)
		origPath = filepath.ToSlash(origPath)
		outPath = filepath.ToSlash(outPath)

		// open files
		origFile, _ := os.Open(origPath)

		// populating bytes
		headBytes := make([]byte, 44)
		tailBytes := make([]byte, 16)
		origFile.Seek(0, 0)
		origFile.Read(headBytes)
		info, _ := origFile.Stat()
		origFile.ReadAt(tailBytes, (info.Size())-16)

		// making output file
		outFile, _ := os.Create(outPath)

		// ADPCM compressing
		encode(outPath, getSampleRate(origPath), custPath)

		// PS2 Encrypting
		pack(outFile, headBytes, tailBytes)

		fmt.Print("File generated successfully. Press Enter to exit.")
		reader.ReadString('\n')
	} else {
		fmt.Println("Usage: ps2adpcm -c custom_audio -i original_audio -o output_file")
		flag.PrintDefaults()
		os.Exit(1)
	}

}
