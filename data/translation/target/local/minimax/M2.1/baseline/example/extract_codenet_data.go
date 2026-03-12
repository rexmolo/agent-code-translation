package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Paths - will be initialized in main
var (
	repoRoot      string
	codeNetRoot   string
	problemListCSV string
	metadataDir   string
	dataDir       string
	outputDir     string
	outputFile    string
)

// Submission represents a row from the metadata CSV
type Submission struct {
	submissionID string
	language     string
	status       string
	codeSize     int
}

// OutputRecord represents a JSON line output record
type OutputRecord struct {
	ProblemID        string `json:"problem_id"`
	PythonCode       string `json:"python_code"`
	GoCode           string `json:"go_code"`
	ProblemDescription string `json:"problem_description"`
}

func findRepoRoot() string {
	// Start from the current executable directory and go up 4 levels
	// This mimics Path(__file__).resolve().parents[4]
	exePath, err := os.Executable()
	if err != nil {
		exePath = "."
	}

	// Try to find the repo root by looking for common markers
	current := filepath.Dir(exePath)
	if current == "/" || current == "" {
		current = "."
	}

	// Walk up to find repo root (look for data/RAG or go.mod or similar)
	for i := 0; i < 10; i++ {
		dataDir := filepath.Join(current, "data", "RAG")
		if _, err := os.Stat(dataDir); err == nil {
			return current
		}
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	// Fallback: assume we're running from repo root
	return "."
}

func readAcceptedSubmissions(problemID, language string) ([]Submission, error) {
	metaCSV := filepath.Join(metadataDir, fmt.Sprintf("%s.csv", problemID))

	file, err := os.Open(metaCSV)
	if err != nil {
		return nil, nil // No file exists
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, nil
	}

	// Find column indices
	submissionIDIdx := -1
	languageIdx := -1
	statusIdx := -1
	codeSizeIdx := -1

	for i, h := range header {
		switch h {
		case "submission_id":
			submissionIDIdx = i
		case "language":
			languageIdx = i
		case "status":
			statusIdx = i
		case "code_size":
			codeSizeIdx = i
		}
	}

	if submissionIDIdx == -1 || languageIdx == -1 || statusIdx == -1 || codeSizeIdx == -1 {
		return nil, nil
	}

	var submissions []Submission
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil
		}

		if len(record) <= submissionIDIdx || len(record) <= languageIdx || 
		   len(record) <= statusIdx || len(record) <= codeSizeIdx {
			continue
		}

		if record[languageIdx] == language && record[statusIdx] == "Accepted" {
			codeSize, _ := strconv.Atoi(record[codeSizeIdx])
			submissions = append(submissions, Submission{
				submissionID: record[submissionIDIdx],
				language:     record[languageIdx],
				status:       record[statusIdx],
				codeSize:     codeSize,
			})
		}
	}

	return submissions, nil
}

func shortestAcceptedCode(problemID, language, langDir, ext string) (string, error) {
	submissions, err := readAcceptedSubmissions(problemID, language)
	if err != nil {
		return "", err
	}
	if len(submissions) == 0 {
		return "", nil
	}

	// Sort by code_size ascending
	sort.Slice(submissions, func(i, j int) bool {
		return submissions[i].codeSize < submissions[j].codeSize
	})

	for _, sub := range submissions {
		filepath := filepath.Join(langDir, fmt.Sprintf("%s.%s", sub.submissionID, ext))
		content, err := os.ReadFile(filepath)
		if err != nil {
			continue
		}
		return string(content), nil
	}

	return "", nil
}

func readDescription(problemID string) string {
	descPath := filepath.Join(dataDir, problemID, "description.html")
	content, err := os.ReadFile(descPath)
	if err != nil {
		return ""
	}
	return string(content)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func main() {
	// Initialize paths
	repoRoot = findRepoRoot()
	codeNetRoot = filepath.Join(repoRoot, "data", "RAG", "unprocessed", "Project_CodeNet")
	problemListCSV = filepath.Join(codeNetRoot, "metadata", "problem_list.csv")
	metadataDir = filepath.Join(codeNetRoot, "metadata")
	dataDir = filepath.Join(codeNetRoot, "data")
	outputDir = filepath.Join(repoRoot, "data", "processed", "parallel_corpus", "codeNet")
	outputFile = filepath.Join(outputDir, "python_go_pairs.jsonl")

	fmt.Println("\033[1;36mCodeNet Parallel Pair Extractor\033[0m")
	fmt.Printf("Reading problem list from: \033[32m%s\033[0m\n", problemListCSV)

	// Check if problem_list.csv exists
	if _, err := os.Stat(problemListCSV); os.IsNotExist(err) {
		fmt.Printf("\033[31mERROR:\033[0m problem_list.csv not found at %s\n", problemListCSV)
		os.Exit(1)
	}

	// Read problem list
	problemListFile, err := os.Open(problemListCSV)
	if err != nil {
		fmt.Printf("Error opening problem list: %v\n", err)
		os.Exit(1)
	}
	defer problemListFile.Close()

	reader := csv.NewReader(problemListFile)
	header, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Find problem_id column (may be 'id' in the CSV)
	idIdx := -1
	for i, h := range header {
		if h == "id" || h == "problem_id" {
			idIdx = i
			break
		}
	}

	if idIdx == -1 {
		fmt.Printf("Could not find 'id' column in problem_list.csv\n")
		os.Exit(1)
	}

	var problemIDs []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if len(record) > idIdx && strings.TrimSpace(record[idIdx]) != "" {
			problemIDs = append(problemIDs, record[idIdx])
		}
	}

	fmt.Printf("Total problems: \033[33m%d\033[0m\n", len(problemIDs))

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Open output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	pairsFound := 0
	skippedNoDirs := 0
	skippedNoAccepted := 0

	// Progress tracking
	fmt.Println()
	for i, problemID := range problemIDs {
		// Simple progress indicator
		if i%10 == 0 {
			percent := float64(i) / float64(len(problemIDs)) * 100
			fmt.Printf("\rProcessing problems: %.1f%% (%d/%d)", percent, i, len(problemIDs))
		}

		pythonDir := filepath.Join(dataDir, problemID, "Python")
		goDir := filepath.Join(dataDir, problemID, "Go")

		if !dirExists(pythonDir) || !dirExists(goDir) {
			skippedNoDirs++
			continue
		}

		pythonCode, err := shortestAcceptedCode(problemID, "Python", pythonDir, "py")
		if err != nil {
			skippedNoAccepted++
			continue
		}
		if pythonCode == "" {
			skippedNoAccepted++
			continue
		}

		goCode, err := shortestAcceptedCode(problemID, "Go", goDir, "go")
		if err != nil {
			skippedNoAccepted++
			continue
		}
		if goCode == "" {
			skippedNoAccepted++
			continue
		}

		description := readDescription(problemID)

		record := OutputRecord{
			ProblemID:            problemID,
			PythonCode:           pythonCode,
			GoCode:               goCode,
			ProblemDescription:   description,
		}

		jsonData, err := json.Marshal(record)
		if err != nil {
			continue
		}

		outFile.WriteString(string(jsonData) + "\n")
		pairsFound++
	}

	fmt.Printf("\rProcessing problems: 100.0%% (%d/%d)\n", len(problemIDs), len(problemIDs))

	fmt.Println()
	fmt.Println("\033[1;32mDone!\033[0m")
	fmt.Printf("  Pairs extracted   : \033[32m%d\033[0m\n", pairsFound)
	fmt.Printf("  Skipped (no dirs) : \033[33m%d\033[0m\n", skippedNoDirs)
	fmt.Printf("  Skipped (no acc.) : \033[33m%d\033[0m\n", skippedNoAccepted)
	fmt.Printf("  Output file       : \033[36m%s\033[0m\n", outputFile)
}
