package helper

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func TestInput(day int, name string) string {
	path := fmt.Sprintf("day%02d/%s", day, name)
	if bytes, err := os.ReadFile(path); err == nil {
		return string(bytes)
	}
	return ""
}

func Input(day int) string {
	name := fmt.Sprintf("day%02d/input.txt", day)

	if bytes, err := os.ReadFile(name); err == nil {
		return string(bytes)
	}

	sessionBytes, err := os.ReadFile("session.txt")
	if err != nil {
		panic(err)
	}
	sessionCookie := string(sessionBytes)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed to download file: %s", resp.Status))
	}

	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		panic(err)
	}

	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
