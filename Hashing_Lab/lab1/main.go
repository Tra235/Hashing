package main

// Lab 1: Made the Password Cracking (MD5)
//

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

// This is the hash we are trying to crack (given in the lab slides).
const targetHash = "6a85dfd77d9cb35770c9dc6728d73d3f"

// This is the wordlist file we will search through.
const wordlistPath = "../data/nord_vpn.txt"

func main() {
	fmt.Println("===== Lab 1: MD5 Password Cracker =====")
	fmt.Println("Target hash:", targetHash)
	fmt.Println("Wordlist file:", wordlistPath)
	fmt.Println()

	// Open the wordlist file for reading.
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Could not open wordlist:", err)
		return
	}
	defer file.Close() // make sure the file gets closed when we're done

	// We'll collect every line we print, so we can save it to a file later.
	var verboseLines []string

	// bufio.Scanner lets us read the file one line at a time.
	scanner := bufio.NewScanner(file)

	attemptNumber := 0
	found := false
	var foundWord string

	for scanner.Scan() {
		word := scanner.Text() // the current line (one password guess)
		if word == "" {
			continue // skip blank lines
		}

		attemptNumber++

		// Hash this word with MD5.
		rawHash := md5.Sum([]byte(word))
		guessHash := hex.EncodeToString(rawHash[:])

		line := fmt.Sprintf("[%d] trying %s -> %s", attemptNumber, word, guessHash)

		// Does our guess match the target hash?
		if guessHash == targetHash {
			line += "  <== MATCH!"
			fmt.Println(line)
			verboseLines = append(verboseLines, line)

			found = true
			foundWord = word
			break // stop the loop, we found it!
		}

		fmt.Println(line)
		verboseLines = append(verboseLines, line)
	}

	fmt.Println()
	if found {
		fmt.Printf("CRACKED! The password is: %s\n", foundWord)
		verboseLines = append(verboseLines, fmt.Sprintf("CRACKED! The password is: %s", foundWord))
	} else {
		fmt.Println("Password not found in the wordlist.")
		verboseLines = append(verboseLines, "Password not found in the wordlist.")
	}

	// Save everything we printed into verbose.txt, as the lab asks for.
	saveVerboseFile("verbose.txt", verboseLines)
}

// saveVerboseFile writes each line of output into a text file.
func saveVerboseFile(path string, lines []string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Could not create verbose file:", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line) // write this line into the file
	}

	fmt.Println("Saved verbose output to", path)
}
