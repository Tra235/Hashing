# Hashing Lab - Beginner Version

A simpler version of the hashing labs: one plain, heavily-commented file
per lab, no custom packages, no build tags. Just read top to bottom.

## Structure
- `lab0/main.go` - Lab 0: compares hashes of two inputs
- `lab1/main.go` - Lab 1: cracks an MD5 hash
- `lab2/main.go` - Lab 2: cracks a SHA1 hash
- `lab3/main.go` - Lab 3: cracks a SHA512 hash
- `data/nord_vpn.txt` - the password wordlist used by labs 1-3
- `third_party/` - a local copy of the SHA3 library (needed because Go's
  standard library doesn't include SHA3 until Go 1.24)

## How to run each lab
Each lab is its own folder with its own `main.go`. To run one, `cd` into
that folder and use `go run .`:

```bash
cd lab0
go run .
```

```bash
cd lab1
go run .
```

Same pattern for lab2 and lab3. Each one only prints its own output, so
nothing gets mixed together in your terminal.

## Expected results
| Lab | Algorithm | Cracked password |
|---|---|---|
| 1 | MD5 | Google101 |
| 2 | SHA1 | Princess1 |
| 3 | SHA512 | Gutter11 |
