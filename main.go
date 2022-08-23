//        ________  ________  ________  ________  ___       ________
//        |\   ___ \|\   __  \|\   __  \|\   __  \|\  \     |\   __  \
//        \ \  \_|\ \ \  \|\  \ \  \|\  \ \  \|\  \ \  \    \ \  \|\  \
//         \ \  \ \\ \ \  \\\  \ \  \\\  \ \   ____\ \  \    \ \   __  \
//          \ \  \_\\ \ \  \\\  \ \  \\\  \ \  \___|\ \  \____\ \  \ \  \
//           \ \_______\ \_______\ \_______\ \__\    \ \_______\ \__\ \__\
//            \|_______|\|_______|\|_______|\|__|     \|_______|\|__|\|__|
//
// 	A Unique URL Identifier for use in bug bounties and penetration testing

package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func parseQueryString(splitQuery []string) []string {
	var queryParams []string
	if len(splitQuery) > 1 {
		i := 0
		for i < len(splitQuery) {
			queryParams = append(queryParams, strings.TrimSpace(splitQuery[i])+"=TEST")
			i = i + 2
		}
	} else {
		queryParams = append(queryParams, strings.Join(splitQuery, ""))
	}

	return queryParams
}

func isStaticContent(u *url.URL) bool {
	static_exts := []string{
		"js", "css", "png", "jpg", "jpeg", "svg",
		"ico", "webp", "ttf", "otf", "woff", "gif",
		"pdf", "bmp", "eot", "mp3", "woff2", "mp4", "avi"}

	for _, ext := range static_exts {
		if strings.Contains(u.Path, ext) {
			return true
		}
	}
	return false
}

func isUserContent(u *url.URL) bool {
	pathParts := strings.Split(u.Path, "/")
	return strings.Contains(pathParts[len(pathParts)-1], "-")
}

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	seenQueryStrings := make([]string, 0)

	for scanner.Scan() {
		txt := scanner.Text()
		u, err := url.Parse(txt)
		if err == nil {
			queryString := u.RawQuery
			splitQuery := strings.Split(queryString, "=")

			var queryParams []string
			queryParams = parseQueryString(splitQuery)
			if !isStaticContent(u) && !isUserContent(u) {
				testQueryString := u.Path + strings.Join(queryParams, "&")
				if !Contains(seenQueryStrings, testQueryString) {
					fmt.Println(txt)
					seenQueryStrings = append(seenQueryStrings, testQueryString)
				}
			}
		}
	}
}
