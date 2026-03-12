package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"
)

// ---------------------------------------------------------------------------
// Paths
// ---------------------------------------------------------------------------

var (
	repoRoot     = findRepoRoot()
	codeNetRoot  = filepath.Join(repoRoot, "data", "RAG", "unprocessed", "Project_CodeNet")
	problemListCSV = filepath.Join(codeNetRoot, "metadata", "problem_list.csv")
	metadataDir    = filepath.Join(codeNetRoot, "metadata")
	dataDir        = filepath.Join(codeNetRoot, "data")
	outputDir      = filepath.Join(repoRoot, "data", "processed", "parallel_corpus", "codeNet")
	outputFile     = filepath.Join(outputDir, "python_go_pairs.jsonl")
)

func findRepoRoot() string {
	// Start from the directory containing this file and go up
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	// Or use current working directory
	if currentDir == "." || currentDir == "" {
		var err error
		currentDir, err = os.Getwd()
		if err != nil {
			return ""
		}
	}

	// Walk up to find repo root (look for common markers)
	for {
		if _, err := os.Stat(filepath.Join(currentDir, ".git")); err == nil {
			return currentDir
		}
		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			break
		}
		currentDir = parent
	}
	return ""
}

// ---------------------------------------------------------------------------
// Data Types
// ---------------------------------------------------------------------------

type Submission struct {
	SubmissionID string
	Language    string
	Status      string
	CodeSize    int64
}

type OutputRecord struct {
	ProblemID         string `json:"problem_id"`
	PythonCode        string `json:"python_code"`
	GoCode            string `json:"go_code"`
	ProblemDescription string `json:"problem_description"`
}

// ---------------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------------

func readAcceptedSubmissions(problemID, language string) ([]Submission, error) {
	metaCSV := filepath.Join(metadataDir, problemID+".csv")
	file, err := os.Open(metaCSV)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Find column indices
	langIdx := -1
	statusIdx := -1
	subIDIdx := -1
	sizeIdx := -1

	for i, h := range header {
		switch h {
		case "language":
			langIdx = i
		case "status":
			statusIdx = i
		case "submission_id":
			subIDIdx = i
		case "code_size":
			sizeIdx = i
		}
	}

	if langIdx == -1 || statusIdx == -1 || subIDIdx == -1 {
		return nil, fmt.Errorf("missing required columns")
	}

	var submissions []Submission
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) <= langIdx || len(record) <= statusIdx || len(record) <= subIDIdx {
			continue
		}

		if record[langIdx] != language || record[statusIdx] != "Accepted" {
			continue
		}

		var codeSize int64
		if sizeIdx != -1 && len(record) > sizeIdx && record[sizeIdx] != "" {
			codeSize, _ = strconv.ParseInt(record[sizeIdx], 10, 64)
		}

		submissions = append(submissions, Submission{
			SubmissionID: record[subIDIdx],
			Language:    record[langIdx],
			Status:      record[statusIdx],
			CodeSize:    codeSize,
		})
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
		return submissions[i].CodeSize < submissions[j].CodeSize
	})

	for _, sub := range submissions {
		filepath := filepath.Join(langDir, sub.SubmissionID+"."+ext)
		data, err := os.ReadFile(filepath)
		if err != nil {
			continue
		}
		return string(data), nil
	}

	return "", nil
}

func readDescription(problemID string) string {
	descPath := filepath.Join(dataDir, problemID, "description.html")
	data, err := os.ReadFile(descPath)
	if err != nil {
		return ""
	}
	return string(data)
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	fmt.Printf("\033[1;36mCodeNet Parallel Pair Extractor\033[0m\n")
	fmt.Printf("Reading problem list from: %s\n", problemListCSV)

	if _, err := os.Stat(problemListCSV); os.IsNotExist(err) {
		fmt.Printf("\033[1;31mERROR:\033[0m problem_list.csv not found at %s\n", problemListCSV)
		os.Exit(1)
	}

	// Read problem list
	file, err := os.Open(problemListCSV)
	if err != nil {
		fmt.Printf("Error opening problem list: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Find id column
	idIdx := -1
	for i, h := range header {
		if h == "id" {
			idIdx = i
			break
		}
	}

	if idIdx == -1 {
		fmt.Printf("Error: 'id' column not found in problem_list.csv\n")
		os.Exit(1)
	}

	var problemIDs []string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		if len(record) > idIdx && record[idIdx] != "" {
			problemIDs = append(problemIDs, record[idIdx])
		}
	}

	fmt.Printf("Total problems: %d\n", len(problemIDs))

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Open output file
	outF, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outF.Close()

	pairsFound := 0
	skippedNoDirs := 0
	skippedNoAccepted := 0

	// Progress tracking
	startTime := time.Now()
	barWidth := 40

	for i, problemID := range problemIDs {
		// Update progress bar
		progress := float64(i+1) / float64(len(problemIDs))
		filled := int(progress * float64(barWidth))
		bar := strings.Repeat("=", filled) + strings.Repeat("-", barWidth-filled)
		elapsed := time.Since(startTime)
		eta := time.Duration(0)
		if progress > 0 {
			eta = time.Duration(float64(elapsed) / progress * (1 - progress))
		}
		fmt.Printf("\r[%s] %d/%d ETA: %v", bar, i+1, len(problemIDs), eta.Round(time.Second))
		fmt.Printf("\033[%dD", 0) // Move cursor back

		pythonDir := filepath.Join(dataDir, problemID, "Python")
		goDir := filepath.Join(dataDir, problemID, "Go")

		// Check if both directories exist
		pythonInfo, err := os.Stat(pythonDir)
		if err != nil || !pythonInfo.IsDir() {
			skippedNoDirs++
			continue
		}

		goInfo, err := os.Stat(goDir)
		if err != nil || !goInfo.IsDir() {
			skippedNoDirs++
			continue
		}

		// Get shortest accepted code for each language
		pythonCode, err := shortestAcceptedCode(problemID, "Python", pythonDir, "py")
		if err != nil || pythonCode == "" {
			skippedNoAccepted++
			continue
		}

		goCode, err := shortestAcceptedCode(problemID, "Go", goDir, "go")
		if err != nil || goCode == "" {
			skippedNoAccepted++
			continue
		}

		description := readDescription(problemID)

		record := OutputRecord{
			ProblemID:           problemID,
			PythonCode:          pythonCode,
			GoCode:              goCode,
			ProblemDescription: description,
		}

		jsonData, err := json.Marshal(record)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			continue
		}

		if _, err := outF.WriteString(string(jsonData) + "\n"); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			continue
		}

		pairsFound++
	}

	// Clear progress line and print summary
	fmt.Printf("\r\033[K") // Clear line
	fmt.Printf("\n\033[1;32mDone!\033[0m\n")
	fmt.Printf("  Pairs extracted   : \033[32m%d\033[0m\n", pairsFound)
	fmt.Printf("  Skipped (no dirs) : \033[33m%d\033[0m\n", skippedNoDirs)
	fmt.Printf("  Skipped (no acc.) : \033[33m%d\033[0m\n", skippedNoAccepted)
	fmt.Printf("  Output file       : \033[36m%s\033[0m\n", outputFile)
}
