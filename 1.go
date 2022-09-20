package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

type mail struct {
	NAME string
	BODY string
	TIME int
}

func CreateFile(name string) error {
	var message string
	fmt.Scanf("%s", &message)
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func createJson(n string) error {
	var name string
	var body string
	var time int
	fmt.Scan(&name)
	fmt.Scan(&body)
	fmt.Scan(&time)
	m := mail{name, body, time}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(n)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(b)
	return nil
}

func createXml(n string) error {
	var name string
	var body string
	var time int
	fmt.Scan(&name)
	fmt.Scan(&body)
	fmt.Scan(&time)
	m := mail{name, body, time}
	v, _ := xml.MarshalIndent(m, "", " ")
	file, err := os.Create(n)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(v)
	return nil
}

func createZip() error {
	var n string
	fmt.Scan(&n)
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open(n)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create("final.txt")
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}
	zipWriter.Close()
	return nil
}

func unzipSource(source string) error {
	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := unzipFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File) error {
	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func main() {

	// 1

	disk_mem, _ := disk.Usage("\\")
	partitions, _ := disk.Partitions(false)
	for _, partition := range partitions {
		fmt.Println(partition.Mountpoint)
	}
	//hostInf,_ := host.Info()
	//platform := hostInf.Platform
	disk_mem_total := disk_mem.Total / 1024 / 1024
	fmt.Println("disk memory", disk_mem_total)
	fmt.Println(runtime.GOOS)

	//2

	CreateFile("3.txt")
	fdata, err := ioutil.ReadFile("3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fdata) + "\n")
	err = os.Remove("3.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 3

	createJson("3.json")
	var m mail
	tdata, err := ioutil.ReadFile("3.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(tdata, &m)
	fmt.Println(string(tdata) + "\n")
	err = os.Remove("3.json")
	if err != nil {
		fmt.Println(err)
	}

	// 4
	createXml("3.xml")
	g := &mail{}
	sdata, err := ioutil.ReadFile("3.xml")
	if err != nil {
		fmt.Println(err)
	}
	xml.Unmarshal([]byte(sdata), g)
	fmt.Println(string(sdata) + "\n")
	err = os.Remove("3.xml")
	if err != nil {
		fmt.Println(err)
	}

	// 5
	createZip()
	err = unzipSource("archive.zip")
	if err != nil {
		fmt.Println(err)
	}
	readFile, err := ioutil.ReadFile("final.txt")
	if err != nil {
		return
	}
	fmt.Println(string(readFile))
	os.Remove("archive.zip")
	os.Remove("final.txt")
}
