package main

// Lab 0: Proof the Hash Program
//
// What this program does, in plain words:
//   1. Ask the user to type two words (we call them input1 and input2).
//   2. Hash each word using 5 different hashing algorithms.
//   3. For each algorithm, print both hashes and say whether they match.
//
// How to run this program:
//   go run main.go
// (run this command from inside the lab0 folder)

import (
	"bufio"         // lets us read a full line of text typed by the user
	"crypto/md5"    // the MD5 hashing algorithm
	"crypto/sha1"   // the SHA1 hashing algorithm
	"crypto/sha256" // the SHA256 hashing algorithm
	"crypto/sha512" // the SHA512 hashing algorithm
	"encoding/hex"  // turns raw hash bytes into a readable hex string
	"fmt"           // for printing text to the screen
	"os"            // lets us read from the keyboard (standard input)
	"strings"       // for cleaning up text (removing extra spaces/newlines)

	"golang.org/x/crypto/sha3" // the SHA3 hashing algorithm (not in the standard library)
)

func main() {
	// bufio.NewReader lets us read a whole line of text at once,
	// instead of reading it one character at a time.
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========Name + Hashing Program========")

	// Ask for the first word.
	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1) // remove the newline/extra spaces

	// Ask for the second word.
	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	// Now hash both words with every algorithm, one at a time.

	fmt.Println("Hash (MD5):")
	compareMD5(input1, input2)

	fmt.Println("Hash (SHA1):")
	compareSHA1(input1, input2)

	fmt.Println("Hash (SHA256):")
	compareSHA256(input1, input2)

	fmt.Println("Hash (SHA512):")
	compareSHA512(input1, input2)

	fmt.Println("Hash (SHA3-256):")
	compareSHA3(input1, input2)
}

// --- Below are 5 small functions, one per algorithm. ---
// Each one does the exact same 3 steps:
//   1. Hash input1  -> outputA
//   2. Hash input2  -> outputB
//   3. Print both, and say Match! or No Match!

func compareMD5(input1, input2 string) {
	// md5.Sum gives us the hash as raw bytes.
	rawA := md5.Sum([]byte(input1))
	rawB := md5.Sum([]byte(input2))

	// hex.EncodeToString turns those raw bytes into a readable string
	// like "5d41402abc4b2a76b9719d911017c592".
	outputA := hex.EncodeToString(rawA[:])
	outputB := hex.EncodeToString(rawB[:])

	printResult(outputA, outputB)
}

func compareSHA1(input1, input2 string) {
	rawA := sha1.Sum([]byte(input1))
	rawB := sha1.Sum([]byte(input2))
	outputA := hex.EncodeToString(rawA[:])
	outputB := hex.EncodeToString(rawB[:])
	printResult(outputA, outputB)
}

func compareSHA256(input1, input2 string) {
	rawA := sha256.Sum256([]byte(input1))
	rawB := sha256.Sum256([]byte(input2))
	outputA := hex.EncodeToString(rawA[:])
	outputB := hex.EncodeToString(rawB[:])
	printResult(outputA, outputB)
}

func compareSHA512(input1, input2 string) {
	rawA := sha512.Sum512([]byte(input1))
	rawB := sha512.Sum512([]byte(input2))
	outputA := hex.EncodeToString(rawA[:])
	outputB := hex.EncodeToString(rawB[:])
	printResult(outputA, outputB)
}

func compareSHA3(input1, input2 string) {
	rawA := sha3.Sum256([]byte(input1))
	rawB := sha3.Sum256([]byte(input2))
	outputA := hex.EncodeToString(rawA[:])
	outputB := hex.EncodeToString(rawB[:])
	printResult(outputA, outputB)
}

// printResult prints two hash outputs and says whether they are equal.
// All 5 functions above call this so we don't repeat the same print code.
func printResult(outputA string, outputB string) {
	fmt.Println("  Output A =", outputA)
	fmt.Println("  Output B =", outputB)

	if outputA == outputB {
		fmt.Println("  => Match!")
	} else {
		fmt.Println("  => No Match!")
	}
}
