package main

import "os"

func fileName(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", nil
	}
	defer f.Close()
	return f.Name(), nil
}
