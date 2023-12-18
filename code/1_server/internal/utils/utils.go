package utils

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetAppName() string {
	str, _ := os.Executable()
	return str
}

func ChangeFileExt(fileName, ext string) string {
	goos := runtime.GOOS
	fmt.Println("current os : ", goos)

	switch runtime.GOOS {
	case "darwin":
	case "windows":
		i := strings.LastIndex(fileName, ".")
		if i > 0 {
			ret := fileName[:i]
			return ret + ext
		}
	case "linux":
		return fileName + ext
	}

	return fileName + ext
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func LocalIp() string {
	conn, e := net.Dial("udp", "8.8.8.8:80")
	if e == nil {
		defer conn.Close()
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		return localAddr.IP.String()
	}
	return "0.0.0.0"
}

func GetWorkStation() string {
	w, _ := os.Hostname()
	return w
}

func ExcelDateToDateString(excelDate int) string {
	if excelDate <= 0 {
		return ""
	}
	// Excel日期从1900年1月1日开始计数，但因为Excel的bug，
	// 它会多出来一个日期，需要减去1
	baseDate := time.Date(1899, 12, 31, 0, 0, 0, 0, time.UTC)
	date := baseDate.AddDate(0, 0, excelDate-1)
	return date.Format("2006-01-02")
}
