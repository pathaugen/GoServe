

package main


import (
	//"fmt"
	"strings"
)


func convUrlToFilename(input string) (string) {
	/*
	//re, err := regexp.Compile(`img.*?src=\"(.*?)\.(.*?)\"`)
	re, err := regexp.Compile(`img.*?src=\"(.*?)\.(.*?)\"`) // parentfolder/subfolder/file.ext
	if err != nil {
		//return "", err
		return ""
	}
	indexes := re.FindAllStringSubmatchIndex(input, -1)
	
	output := input
	for _, match := range indexes {
		imgStart := match[2]
		imgEnd := match[3]
		newImgName := strings.Replace(input[imgStart:imgEnd], "m", "a", -1)
		output = output[:imgStart] + newImgName + input[imgEnd:]
	}
	return output //, nil
	*/
	
	longFilenameArr := strings.Split(input, ".")
	longFilename := strings.Replace(longFilenameArr[0], "/", "--", -1)
	
	/*
	arr := strings.Split(input, "/") // Array of all parts of input
	filename := arr[len(arr)-1] // Last part of array
	*/
	
	/*
	filearr := strings.Split(filename, ".")
	file := filearr[0]
	ext := filearr[len(arr)-1]
	*/
	
	return longFilename
}

