package main

// Lab 2: Made the Password Cracking (SHA1)
//


import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

const targetHash = "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
const wordlistPath = "../data/nord_vpn.txt"

func main() {
	fmt.Println("===== Lab 2: SHA1 Password Cracker =====")
	fmt.Println("Target hash:", targetHash)
	fmt.Println("Wordlist file:", wordlistPath)
	fmt.Println()

	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Could not open wordlist:", err)
		return
	}
	defer file.Close()

	var verboseLines []string
	scanner := bufio.NewScanner(file)

	attemptNumber := 0
	found := false
	var foundWord string

	for scanner.Scan() {
		word := scanner.Text()
		if word == "" {
			continue
		}

		attemptNumber++

		rawHash := sha1.Sum([]byte(word))
		guessHash := hex.EncodeToString(rawHash[:])

		line := fmt.Sprintf("[%d] trying %s -> %s", attemptNumber, word, guessHash)

		if guessHash == targetHash {
			line += "  <== MATCH!"
			fmt.Println(line)
			verboseLines = append(verboseLines, line)

			found = true
			foundWord = word
			break
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

	saveVerboseFile("verbose.txt", verboseLines)
}

func saveVerboseFile(path string, lines []string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Could not create verbose file:", err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		fmt.Fprintln(file, line)
	}

	fmt.Println("Saved verbose output to", path)
}
