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

# Include statistics in output
subwords -i subdomains.txt -stats
```

The tool splits subdomains by dots (.), hyphens (-), and underscores (_), counts the frequency of each word, and outputs them sorted by frequency.
