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
	"time"
)

// ---------------------------------------------------------------------------
// Paths
// ---------------------------------------------------------------------------

func findRepoRoot() string {
	// Start from the location of the executable and go up 4 levels
	exePath, err := os.Executable()
	if err != nil {
		exePath = ""
	}
	// For development, use the current working directory
	cwd, _ := os.Getwd()
	if exePath == "" {
		exePath = cwd
	}
	
	// Try to find repo root by walking up from current dir
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(cwd, "data", "RAG", "unprocessed", "Project_CodeNet")); err == nil {
			return cwd
		}
		cwd = filepath.Dir(cwd)
	}
	return "."
}

var (
	repoRoot       = findRepoRoot()
	codeNetRoot    = filepath.Join(repoRoot, "data", "RAG", "unprocessed", "Project_CodeNet")
	problemListCSV = filepath.Join(codeNetRoot, "metadata", "problem_list.csv")
	metadataDir    = filepath.Join(codeNetRoot, "metadata")
	dataDir        = filepath.Join(codeNetRoot, "data")
	outputDir      = filepath.Join(repoRoot, "data", "processed", "parallel_corpus", "codeNet")
	outputFile     = filepath.Join(outputDir, "python_go_pairs.jsonl")
)

// ---------------------------------------------------------------------------
// Data Structures
// ---------------------------------------------------------------------------

type Problem struct {
	ID string
}

type Submission struct {
	Language  string
	Status    string
	SubmitID  string
	CodeSize  int64
}

type Result struct {
	ProblemID        string `json:"problem_id"`
	PythonCode       string `json:"python_code"`
	GoCode           string `json:"go_code"`
	ProblemDescHTML  string `json:"problem_description"`
}

// ---------------------------------------------------------------------------
// Helper Functions
// ---------------------------------------------------------------------------

func readProblemList() ([]string, error) {
	file, err := os.Open(problemListCSV)
	if err != nil {
		return nil, fmt.Errorf("failed to open problem_list.csv: %w", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	// Skip header
	_, err = r.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	var problemIDs []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV: %w", err)
		}
		if len(record) > 0 && record[0] != "" {
			problemIDs = append(problemIDs, record[0])
		}
	}
	return problemIDs, nil
}

func readAcceptedSubmissions(problemID, language string) ([]Submission, error) {
	metaCSV := filepath.Join(metadataDir, problemID+".csv")
	file, err := os.Open(metaCSV)
	if err != nil {
		return nil, nil // No file exists, return nil
	}
	defer file.Close()

	r := csv.NewReader(file)
	// Read header to find column indices
	header, err := r.Read()
	if err != nil {
		return nil, nil
	}
	
	// Find column indices
	langIdx := -1
	statusIdx := -1
	subIDIdx := -1
	sizeIdx := -1
	
	for i, col := range header {
		switch col {
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
	
	if langIdx == -1 || statusIdx == -1 || subIDIdx == -1 || sizeIdx == -1 {
		return nil, nil
	}

	var accepted []Submission
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil
		}
		
		if len(record) <= langIdx || len(record) <= statusIdx || 
		   len(record) <= subIDIdx || len(record) <= sizeIdx {
			continue
		}
		
		if record[langIdx] == language && record[statusIdx] == "Accepted" {
			codeSize, _ := strconv.ParseInt(record[sizeIdx], 10, 64)
			accepted = append(accepted, Submission{
				Language: record[langIdx],
				Status:   record[statusIdx],
				SubmitID: record[subIDIdx],
				CodeSize: codeSize,
			})
		}
	}
	return accepted, nil
}

func shortestAcceptedCode(problemID, language, langDir, ext string) string {
	accepted, err := readAcceptedSubmissions(problemID, language)
	if err != nil || accepted == nil || len(accepted) == 0 {
		return ""
	}

	// Sort by code size ascending
	sort.Slice(accepted, func(i, j int) bool {
		return accepted[i].CodeSize < accepted[j].CodeSize
	})

	for _, sub := range accepted {
		filepath := filepath.Join(langDir, sub.SubmitID+"."+ext)
		content, err := os.ReadFile(filepath)
		if err != nil {
			continue
		}
		return string(content)
	}

	return ""
}

func readDescription(problemID string) string {
	descPath := filepath.Join(dataDir, problemID, "description.html")
	content, err := os.ReadFile(descPath)
	if err != nil {
		return ""
	}
	return string(content)
}

func ensureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// ---------------------------------------------------------------------------
// Progress Display
// ---------------------------------------------------------------------------

type ProgressBar struct {
	total    int
	current  int
	startTime time.Time
}

func NewProgressBar(total int) *ProgressBar {
	return &ProgressBar{
		total:    total,
		current:  0,
		startTime: time.Now(),
	}
}

func (p *ProgressBar) Advance() {
	p.current++
}

func (p *ProgressBar) Print() {
	elapsed := time.Since(p.startTime)
	percent := float64(p.current) / float64(p.total) * 100
	
	// Estimate remaining time
	var remaining time.Duration
	if p.current > 0 {
		avgTime := elapsed / time.Duration(p.current)
		remaining = avgTime * time.Duration(p.total-p.current)
	}
	
	// Clear line and print progress
	fmt.Printf("\r[%d/%d] %.1f%% - ETA: %v", p.current, p.total, percent, remaining)
	if p.current == p.total {
		fmt.Println()
	}
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	fmt.Printf("\033[1;36mCodeNet Parallel Pair Extractor\033[0m\n")
	fmt.Printf("Reading problem list from: \033[32m%s\033[0m\n", problemListCSV)

	if _, err := os.Stat(problemListCSV); os.IsNotExist(err) {
		fmt.Printf("\033[31mERROR:\033[0m problem_list.csv not found at %s\n", problemListCSV)
		os.Exit(1)
	}

	problemIDs, err := readProblemList()
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m Failed to read problem list: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total problems: \033[33m%d\033[0m\n", len(problemIDs))

	if err := ensureDir(outputDir); err != nil {
		fmt.Printf("\033[31mERROR:\033[0m Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	outF, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m Failed to create output file: %v\n", err)
		os.Exit(1)
	}
	defer outF.Close()

	pairsFound := 0
	skippedNoDirs := 0
	skippedNoAccepted := 0

	progress := NewProgressBar(len(problemIDs))
	fmt.Printf("Processing problems...")
	progress.Print()

	for i, problemID := range problemIDs {
		pythonDir := filepath.Join(dataDir, problemID, "Python")
		goDir := filepath.Join(dataDir, problemID, "Go")

		pythonDirInfo, err := os.Stat(pythonDir)
		if err != nil || !pythonDirInfo.IsDir() {
			skippedNoDirs++
			progress.Advance()
			progress.Print()
			_ = i // suppress unused variable warning
			continue
		}

		goDirInfo, err := os.Stat(goDir)
		if err != nil || !goDirInfo.IsDir() {
			skippedNoDirs++
			progress.Advance()
			progress.Print()
			continue
		}

		pythonCode := shortestAcceptedCode(problemID, "Python", pythonDir, "py")
		goCode := shortestAcceptedCode(problemID, "Go", goDir, "go")

		if pythonCode == "" || goCode == "" {
			skippedNoAccepted++
			progress.Advance()
			progress.Print()
			continue
		}

		description := readDescription(problemID)

		record := Result{
			ProblemID:        problemID,
			PythonCode:       pythonCode,
			GoCode:           goCode,
			ProblemDescHTML:  description,
		}

		jsonData, err := json.Marshal(record)
		if err != nil {
			fmt.Printf("\n\033[31mERROR:\033[0m Failed to marshal JSON: %v\n", err)
			continue
		}

		outF.Write(jsonData)
		outF.WriteString("\n")

		pairsFound++
		progress.Advance()
		progress.Print()
	}

	fmt.Printf("\n\033[1;32mDone!\033[0m\n")
	fmt.Printf("  Pairs extracted   : \033[32m%d\033[0m\n", pairsFound)
	fmt.Printf("  Skipped (no dirs) : \033[33m%d\033[0m\n", skippedNoDirs)
	fmt.Printf("  Skipped (no acc.) : \033[33m%d\033[0m\n", skippedNoAccepted)
	fmt.Printf("  Output file       : \033[36m%s\033[0m\n", outputFile)
}
