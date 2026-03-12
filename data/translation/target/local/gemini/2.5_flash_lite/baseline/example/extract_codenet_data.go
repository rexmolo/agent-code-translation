package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/Puerkito/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
)

// ---------------------------------------------------------------------------
// Paths
// ---------------------------------------------------------------------------

var ( 
	RepoRoot string
	CodNetRoot string
	ProblemListCSV string
	MetaDataDir string
	DataDir string
	OutputDir string
	OutputFile string
)

func initPaths() {
	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Assuming REPO_ROOT is 4 levels up from the current file's directory
	// This might need adjustment based on the actual project structure.
	RepoRoot = "../../../../"
	CodNetRoot = filepath.Join(RepoRoot, "data", "RAG", "unprocessed", "Project_CodeNet")
	ProblemListCSV = filepath.Join(CodNetRoot, "metadata", "problem_list.csv")
	MetaDataDir = filepath.Join(CodNetRoot, "metadata")
	DataDir = filepath.Join(CodNetRoot, "data")
	OutputDir = filepath.Join(RepoRoot, "data", "processed", "parallel_corpus", "codeNet")
	OutputFile = filepath.Join(OutputDir, "python_go_pairs.jsonl")
}

// ---------------------------------------------------------------------------
// Data Structures
// ---------------------------------------------------------------------------

type Submission struct {
	SubmissionID string `json:"submission_id"`
	Language     string `json:"language"`
	StatusCode   string `json:"status_code"` // Corresponds to 'status' in Python
	CodeSize     string `json:"code_size"`
}

type Problem struct {
	ID          string `json:"id"`
	Description string `json:"problem_description"`
}

type ParallelPair struct {
	ProblemID           string `json:"problem_id"`
	PythonCode          string `json:"python_code"`
	GoCode              string `json:"go_code"`
	ProblemDescription string `json:"problem_description"`
}

// ---------------------------------------------------------------------------
// Helper Functions
// ---------------------------------------------------------------------------

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Basic CSV parsing, assumes no complex quoting or embedded commas
	var data [][]string
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		data = append(data, fields)
	}
	return data, nil
}

func readAcceptedSubmissions(problemID string, language string) ([]Submission, error) {
	metaFilePath := filepath.Join(MetaDataDir, problemID+".csv")
	if _, err := os.Stat(metaFilePath); os.IsNotExist(err) {
		return nil, nil
	}

	csvData, err := readCSV(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading metadata CSV %s: %v", metaFilePath, err)
	}

	if len(csvData) < 2 { // Need at least header and one row
		return nil, nil
	}

	// Infer header indexes
	header := csvData[0]
	var submissionIDIdx, languageIdx, statusCodeIdx, codeSizeIdx int = -1, -1, -1, -1
	for i, col := range header {
		switch col {
		case "submission_id":
			submissionIDIdx = i
		case "language":
			languageIdx = i
		case "status_code":
			statusCodeIdx = i
		case "code_size":
			codeSizeIdx = i
		}
	}

	if submissionIDIdx == -1 || languageIdx == -1 || statusCodeIdx == -1 || codeSizeIdx == -1 {
		return nil, fmt.Errorf("missing required columns in %s", metaFilePath)
	}

	var acceptedSubmissions []Submission
	for _, row := range csvData[1:] {
		if len(row) <= submissionIDIdx || len(row) <= languageIdx || len(row) <= statusCodeIdx || len(row) <= codeSizeIdx {
			continue // Skip malformed rows
		}
		if row[languageIdx] == language && row[statusCodeIdx] == "Accepted" {
			acceptedSubmissions = append(acceptedSubmissions, Submission{
				SubmissionID: row[submissionIDIdx],
				Language:     row[languageIdx],
				StatusCode:   row[statusCodeIdx],
				CodeSize:     row[codeSizeIdx],
			})
		}
	}

	return acceptedSubmissions, nil
}

type SubmissionBySize []Submission

func (a SubmissionBySize) Len() int {
	return len(a)
}

func (a SubmissionBySize) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SubmissionBySize) Less(i, j int) bool {
	// Convert code_size to int for comparison
	sizeI, errI := strconv.Atoi(a[i].CodeSize)
	if errI != nil {
		sizeI = 0 // Treat unparseable as smallest
	}
	sizeJ, errJ := strconv.Atoi(a[j].CodeSize)
	if errJ != nil {
		sizeJ = 0 // Treat unparseable as smallest
	}
	return sizeI < sizeJ
}

func shortestAcceptedCode(problemID string, language string, langDir string, ext string) (string, error) {
	accepted, err := readAcceptedSubmissions(problemID, language)
	if err != nil {
		return "", fmt.Errorf("failed to read accepted submissions for %s %s: %w", problemID, language, err)
	}

	if len(accepted) == 0 {
		return "", nil
	}

	ssort.Sort(SubmissionBySize(accepted))

	for _, sub := range accepted {
		filePath := filepath.Join(langDir, sub.SubmissionID+"."+ext)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			continue
		}

		code, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue // Ignore errors reading file, try next
		}
		return string(code), nil
	}

	return "", nil
}

func readDescription(problemID string) (string, error) {
	descPath := filepath.Join(DataDir, problemID, "description.html")
	if _, err := os.Stat(descPath); os.IsNotExist(err) {
		return "", nil
	}

	content, err := ioutil.ReadFile(descPath)
	if err != nil {
		return "", fmt.Errorf("error reading description file %s: %w", descPath, err)
	}
	return string(content), nil
}

// ---------------------------------------------------------------------------
// Main Execution
// ---------------------------------------------------------------------------

func main() {
	initPaths()

	fmt.Println("\033[1m\033[36mCodeNet Parallel Pair Extractor\033[0m")
	fmt.Printf("Reading problem list from: \033[32m%s\033[0m\n", ProblemListCSV)

	if _, err := os.Stat(ProblemListCSV); os.IsNotExist(err) {
		fmt.Printf("\033[31mERROR:\033[0m problem_list.csv not found at %s\n", ProblemListCSV)
		os.Exit(1)
	}

	problemListData, err := readCSV(ProblemListCSV)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m failed to read problem list: %v\n", err)
		os.Exit(1)
	}

	var problemIDs []string
	// Assuming 'id' is the first column in problem_list.csv
	if len(problemListData) > 1 {
		for _, row := range problemListData[1:] {
			if len(row) > 0 && row[0] != "" {
				problemIDs = append(problemIDs, row[0])
			}
		}
	}

	fmt.Printf("Total problems: \033[33m%d\033[0m\n", len(problemIDs))

	err = os.MkdirAll(OutputDir, 0755)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m failed to create output directory %s: %v\n", OutputDir, err)
		os.Exit(1)
	}

	pairsFound := 0
	skippedNoDirs := 0
	skippedNoAccepted := 0

	bar := progressbar.NewOptions(len(problemIDs),
		progressbar.OptionSetDescription("Processing problems"),
		progressbar.OptionSetWriter(os.Stdout),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(100),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionOnCompletion(func() {
			fmt.Println("\nProcessing complete.")
		}),
	)

	outFile, err := os.OpenFile(OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("\033[31mERROR:\033[0m failed to open output file %s: %v\n", OutputFile, err)
		os.Exit(1)
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)

	for _, problemID := range problemIDs {
		pythonDir := filepath.Join(DataDir, problemID, "Python")
		goDir := filepath.Join(DataDir, problemID, "Go")

		pythonDirInfo, errPython := os.Stat(pythonDir)
		goDirInfo, errGo := os.Stat(goDir)

		if os.IsNotExist(errPython) || !pythonDirInfo.IsDir() || os.IsNotExist(errGo) || !goDirInfo.IsDir() {
			skippedNoDirs++
			bar.Increment()
			continue
		}

		pythonCode, err := shortestAcceptedCode(problemID, "Python", pythonDir, "py")
		if err != nil {
			fmt.Printf("\nWarning: error getting Python code for %s: %v\n", problemID, err)
			pythonCode = ""
		}

		goCode, err := shortestAcceptedCode(problemID, "Go", goDir, "go")
		if err != nil {
			fmt.Printf("\nWarning: error getting Go code for %s: %v\n", problemID, err)
			goCode = ""
		}

		if pythonCode == "" || goCode == "" {
			skippedNoAccepted++
			bar.Increment()
			continue
		}

		description, err := readDescription(problemID)
		if err != nil {
			fmt.Printf("\nWarning: error reading description for %s: %v\n", problemID, err)
			description = ""
		}

		record := ParallelPair{
			ProblemID:          problemID,
			PythonCode:         pythonCode,
			GoCode:             goCode,
			ProblemDescription: description,
		}

		err = encoder.Encode(record)
		if err != nil {
			fmt.Printf("\nWarning: error encoding JSON for problem %s: %v\n", problemID, err)
		}

		pairsFound++
		bar.Increment()
	}

	fmt.Println("\n\033[1m\033[32mDone![/033[0m")
	fmt.Printf("  Pairs extracted   : \033[32m%d\033[0m\n", pairsFound)
	fmt.Printf("  Skipped (no dirs) : \033[33m%d\033[0m\n", skippedNoDirs)
	fmt.Printf("  Skipped (no acc.) : \033[33m%d\033[0m\n", skippedNoAccepted)
	fmt.Printf("  Output file       : \033[36m%s\033[0m\n", OutputFile)
}
