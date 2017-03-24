package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	//parseExcelToFile()
	parseData()
	//parseData2()
}

func parseData2() {
	userFile2 := "gaode_code_2.txt"
	fileIn, err := os.Open(userFile2)
	if err != nil {
		panic(err)
	}
	defer fileIn.Close()

	resultMap := make(map[string][]string)
	var cityName string

	rd := bufio.NewReader(fileIn)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if strings.HasPrefix(line, "#############") {
			continue
		}

		fieldArray := strings.Split(line, "|")
		addressName := fieldArray[0]
		//addressCode := fieldArray[1]

		if strings.Contains(line, "-") {
			cityName = strings.Split(addressName, "-")[0]
		} else {
			cityName = addressName
		}

		if _, ok := resultMap[cityName]; ok {

		} else {

		}
	}

	for key, value := range resultMap {
		fmt.Printf("%s\n", key)
		for _, v := range value {
			fmt.Printf("%s\n", v)
		}
	}
}

func parseData() {
	userFile2 := "gaode_code_2.txt"
	fileOut, err := os.Create(userFile2)
	defer fileOut.Close()
	if err != nil {
		fmt.Println(userFile2, err)
		return
	}

	userFile := "gaode_code.txt"
	directlyCityDict := map[string]int{"北京市": 1, "上海市": 1, "重庆市": 1, "天津市": 1}
	var cityCode string
	var cityName string
	var newLine string

	cityDict := make(map[string]string)

	fileIn, err := os.Open(userFile)
	if err != nil {
		panic(err)
	}
	defer fileIn.Close()

	rd := bufio.NewReader(fileIn)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if strings.HasPrefix(line, "#############") {
			fileOut.WriteString(line)
			continue
		}

		newLine = line
		fieldArray := strings.Split(line, "|")
		addressName := fieldArray[0]
		addressCode := fieldArray[1]
		//addressNameMengNiu := fieldArray[2]

		i, err := strconv.ParseInt(addressCode, 10, 32)
		if err != nil {
			continue
		} else {
			println(i)
		}

		// 找到城市编码
		if _, ok := directlyCityDict[addressName]; ok {
			cityCode = addressCode
			cityName = addressName
			cityDict[cityCode] = cityName
		}

		if len(addressCode) == 6 {
			cityCode = addressCode
			cityName = addressName
			cityDict[cityCode] = cityName
		}

		// 同一个城市下面的三级地区
		if strings.HasPrefix(addressCode, cityCode) && len(addressCode) > 6 {
			newLine = cityName + "-" + line
		}

		// 同一个城市下面的三级地区
		if len(addressCode) > 6 {
			//fmt.Printf("%s-%s\n", addressName, addressCode)
			prefixCode := Substr2(addressCode, 0, 6)
			if _, ok := cityDict[prefixCode]; ok {
				newLine = cityDict[prefixCode] + "-" + line
			}
		}

		//fmt.Printf("%s\n", newLine)

		fileOut.WriteString(newLine)

		/*if _, ok := cityMap[cityName]; ok {
			cityMap[fieldArray[0]] += 1
		} else {
			cityMap[fieldArray[0]] = 1
		}*/
	}

	/*for key, value := range cityMap {
		if value > 1{
			fmt.Printf("%s->%d\n", key, value)

		}
		//fmt.Printf("%s->%d\n", key, value)
	}*/

}

func parseExcelToFile() {
	excelFileName := "/Users/xiaolezheng/tmp/20170321_CE.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	GROUP_SIZE := 3
	if err != nil {
		fmt.Println("Hello, 世界")
		log.Print(err)
		return
	}

	userFile := "gaode_code.txt"
	fileOut, err := os.Create(userFile)
	defer fileOut.Close()
	if err != nil {
		fmt.Println(userFile, err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		rowIndex := 0
		var province string
		for _, row := range sheet.Rows {
			rowIndex++
			if rowIndex == 1 {
				continue
			}

			cellIndex := 1
			var groupStr string
			for _, cell := range row.Cells {
				text, _ := cell.String()
				if rowIndex == 2 && cellIndex == 1 {
					province = text
					segmentLine := "############################################" + province + "############################################"
					fileOut.WriteString(segmentLine + "\r\n")
				}

				groupStr += text
				if cellIndex%GROUP_SIZE == 0 {
					/*if strings.HasSuffix(groupStr, "|") {
						groupStr += city
					}*/
					groupStr = strings.Replace(groupStr, " ", "", -1)
					if groupStr != "||" && !strings.HasPrefix(groupStr, "|") {
						fmt.Printf("%s\n", groupStr)
						fileOut.WriteString(groupStr + "\r\n")
					}
					groupStr = ""
				} else {
					groupStr += "|"
				}

				cellIndex++
			}

		}
	}
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
