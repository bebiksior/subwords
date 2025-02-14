# subwords
A simple tool to extract words from a list of subdomains, sort them by frequency, and output them.

## Installation

```bash
go install github.com/bebiksior/subwords@latest
```

## Usage

```bash
# Read from stdin
cat subdomains.txt | subwords

# Read from file
subwords -i subdomains.txt

# Limit output to top N most frequent words
subwords -i subdomains.txt -limit 10
```

The tool splits subdomains by dots (.) and hyphens (-), counts the frequency of each word, and outputs them sorted by frequency.
