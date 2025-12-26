package main

import (
	"strings"
)
func cleanInput(text string) []string{
	text = strings.ToLower(text)
	updatedSlice := []string{}
	n := len(text)
	i:=0
	for i<n{
		if text[i]==' ' {
			i++
			continue
		}
		newtext := ""
		for i<n && text[i]!=' '{
			newtext+=string(text[i])
			i++
		}
		updatedSlice=append(updatedSlice, newtext)
	}
	return updatedSlice
}