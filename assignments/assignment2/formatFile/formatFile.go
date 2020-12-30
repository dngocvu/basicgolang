package main


import ( 
	"fmt"
	//"regexp"
	"strings"
	"strconv"
)

// 1. split multi line string
// 2. check how many file in type
// 3. add in substring
// 4. sum size
// 5. printf

func TotalSize(inputTotalSize string)(outputTotalSize string){
	var strList [] string
	for _, line := range strings.Split(strings.TrimSuffix(inputTotalSize, "\n"), "\n") {
		strList = append(strList, line)
	}
	var musicSlice [] string
	var imagesSlice [] string
	var moviesSlice [] string
	var otherSlice [] string
	var musicSize, imagesSize, movieSize, otherSize int = 0, 0, 0, 0
	for i := 0; i < len(strList); i++ {
		if strings.Contains(strList[i], ".mp3")||strings.Contains(strList[i], ".mp3")||strings.Contains(strList[i], ".flac") {
			musicSlice = append(musicSlice, strList[i])
		} else if strings.Contains(strList[i], ".jpg")||strings.Contains(strList[i], ".bmp")||strings.Contains(strList[i], ".gif") {
			imagesSlice = append(imagesSlice, strList[i])
		} else if strings.Contains(strList[i], ".mp4")||strings.Contains(strList[i], ".avi")||strings.Contains(strList[i], ".mkv") {
			moviesSlice = append(moviesSlice, strList[i])
		} else {
			otherSlice = append(otherSlice, strList[i])
		}
	}
	musicSliceSize := make([]int, len(musicSlice))
	imagesSliceSize := make([]int, len(imagesSlice))
	moviesSliceSize := make([]int, len(moviesSlice))
	otherSliceSize := make([]int, len(otherSlice))
	var parseMusicStringSize string = ""
	var parseMusicStringSliceSize [] string
	if len(musicSlice) > 0 {
		for i := 0; i < len(musicSlice); i++ {
			musicSlice[i] = strings.TrimRight(musicSlice[i], "b")
		}
		for i := 0; i < len(musicSlice); i++ {
			parseMusicStringSize = parseMusicStringSize + " " + musicSlice[i]
		}
		parseMusicStringSliceSize = strings.Fields(parseMusicStringSize)
		k_music := 0
		for i := 1; i < len(parseMusicStringSliceSize); i = i + 2 {
			musicSliceSize[k_music],_ = strconv.Atoi(parseMusicStringSliceSize[i])
			k_music += 1
		}
		for i := 0; i < len(musicSliceSize); i++ {
			musicSize = musicSize + musicSliceSize[i]
		}
	} else {
		musicSize = 0
	}
	var parseImageStringSize string = ""
	var parseImageStringSliceSize [] string
	if len(imagesSlice) > 0 {
		for i := 0; i < len(imagesSlice); i++ {
			imagesSlice[i] = strings.TrimRight(imagesSlice[i], "b")
		}
		for i := 0; i < len(imagesSlice); i++ {
			parseImageStringSize = parseImageStringSize + " " + imagesSlice[i]
		}
		parseImageStringSliceSize = strings.Fields(parseImageStringSize)
		k_images := 0
		for i := 1; i < len(parseImageStringSliceSize); i = i + 2 {
			imagesSliceSize[k_images],_ = strconv.Atoi(parseImageStringSliceSize[i])
			k_images += 1
		}
		for i := 0; i < len(imagesSliceSize); i++ {
			imagesSize = imagesSize + imagesSliceSize[i]
		}
	} else {
		imagesSize = 0
	}
	var parseMovieStringSize string = ""
	var parseMovieStringSliceSize [] string
	if len(moviesSlice) > 0 {
		for i := 0; i < len(moviesSlice); i++ {
			moviesSlice[i] = strings.TrimRight(moviesSlice[i], "b")
		}
		for i := 0; i < len(moviesSlice); i++ {
			parseMovieStringSize = parseMovieStringSize + " " + moviesSlice[i]
		}
		parseMovieStringSliceSize = strings.Fields(parseMovieStringSize)
		k_movies := 0
		for i := 1; i < len(parseMovieStringSliceSize); i = i + 2 {
			otherSliceSize[k_movies],_ = strconv.Atoi(parseMovieStringSliceSize[i])
			k_movies += 1
		}
		for i := 0; i < len(otherSliceSize); i++ {
			movieSize = movieSize + otherSliceSize[i]
		}
	} else {
		movieSize = 0
	}
	var parseOtherStringSize string = ""
	var parseOtherStringSliceSize [] string
	if len(otherSlice) > 0 {
		for i := 0; i < len(otherSlice); i++ {
			otherSlice[i] = strings.TrimRight(otherSlice[i], "b")
		}
		for i := 0; i < len(otherSlice); i++ {
			parseOtherStringSize = parseOtherStringSize + " " + otherSlice[i]
		}
		parseOtherStringSliceSize = strings.Fields(parseOtherStringSize)
		k_others := 0
		for i := 1; i < len(parseOtherStringSliceSize); i = i + 2 {
			moviesSliceSize[k_others],_ = strconv.Atoi(parseOtherStringSliceSize[i])
			k_others += 1
		}
		for i := 0; i < len(moviesSliceSize); i++ {
			otherSize = otherSize + moviesSliceSize[i]
		}
	} else {
		otherSize = 0
	}
	outputTotalSize = "music " + strconv.Itoa(musicSize) + "b\n" + "images " + strconv.Itoa(imagesSize) + "b\n" + "movies " + strconv.Itoa(movieSize) + "b\n" + "other " + strconv.Itoa(otherSize) + "b\n"
	return
}

func main () {
	var str string =   
`my.song.mp3 11b
greatSong.flac 1000b
not3.txt 5b
video.mp4 200b
game.exe 100b
mov!e.mkv 10000b`
	
	fmt.Println(TotalSize(str))

}