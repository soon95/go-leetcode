package no71

import "strings"

func simplifyPath(path string) string {

	split := strings.Split(path, "/")

	nodeList := make([]string, 0)

	for _, s := range split {
		if len(s) == 0 {
			continue
		} else if s == "." {
			continue
		} else if s == ".." {
			if len(nodeList) != 0 {
				nodeList = nodeList[:len(nodeList)-1]
			}
		} else {
			nodeList = append(nodeList, s)
		}
	}

	return "/" + strings.Join(nodeList, "/")

}
