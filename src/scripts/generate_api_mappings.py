#!/usr/bin/env python3
"""Generate additional API mappings and append to api_mappings.jsonl.

Usage:
    uv run python src/scripts/generate_api_mappings.py
"""

import json

from rich.console import Console

from src.config import API_MAPPINGS_FILE

console = Console()

# Curated Python-to-Go API mappings for categories not yet covered
NEW_MAPPINGS = [
    # --- type_conversion ---
    {"category": "type_conversion", "python_api": "int()", "go_api": "strconv.Atoi()", "description": "Converts string to integer."},
    {"category": "type_conversion", "python_api": "int(x, base)", "go_api": "strconv.ParseInt(s, base, 64)", "description": "Converts string to integer with a given base."},
    {"category": "type_conversion", "python_api": "str()", "go_api": "strconv.Itoa()", "description": "Converts integer to string."},
    {"category": "type_conversion", "python_api": "float()", "go_api": "strconv.ParseFloat(s, 64)", "description": "Converts string to float64."},
    {"category": "type_conversion", "python_api": "str(float_val)", "go_api": "strconv.FormatFloat(f, 'f', -1, 64)", "description": "Converts float to string."},
    {"category": "type_conversion", "python_api": "bool()", "go_api": "strconv.ParseBool()", "description": "Converts string to boolean."},
    {"category": "type_conversion", "python_api": "hex()", "go_api": "fmt.Sprintf(\"%x\", n)", "description": "Converts integer to hexadecimal string."},
    {"category": "type_conversion", "python_api": "bin()", "go_api": "fmt.Sprintf(\"%b\", n)", "description": "Converts integer to binary string."},
    {"category": "type_conversion", "python_api": "oct()", "go_api": "fmt.Sprintf(\"%o\", n)", "description": "Converts integer to octal string."},
    {"category": "type_conversion", "python_api": "ord()", "go_api": "int(rune)", "description": "Returns the Unicode code point of a character."},
    {"category": "type_conversion", "python_api": "chr()", "go_api": "string(rune(n))", "description": "Returns the character for a Unicode code point."},
    {"category": "type_conversion", "python_api": "abs()", "go_api": "math.Abs()", "description": "Returns the absolute value (float64 in Go; for int use manual check)."},

    # --- io ---
    {"category": "io", "python_api": "input()", "go_api": "fmt.Scan() / bufio.NewScanner(os.Stdin)", "description": "Reads input from stdin."},
    {"category": "io", "python_api": "print()", "go_api": "fmt.Println()", "description": "Prints a line to stdout with newline."},
    {"category": "io", "python_api": "print(f'...')", "go_api": "fmt.Printf()", "description": "Prints formatted string to stdout."},
    {"category": "io", "python_api": "print(x, end='')", "go_api": "fmt.Print()", "description": "Prints without trailing newline."},
    {"category": "io", "python_api": "sys.stdin", "go_api": "bufio.NewScanner(os.Stdin)", "description": "Line-by-line stdin reading."},
    {"category": "io", "python_api": "sys.stdout.write()", "go_api": "fmt.Print() / os.Stdout.WriteString()", "description": "Writes string to stdout without newline."},
    {"category": "io", "python_api": "sys.stderr.write()", "go_api": "fmt.Fprint(os.Stderr, ...)", "description": "Writes to stderr."},
    {"category": "io", "python_api": "sys.exit()", "go_api": "os.Exit()", "description": "Exits the program with a status code."},
    {"category": "io", "python_api": "sys.argv", "go_api": "os.Args", "description": "Command-line arguments."},

    # --- sorting ---
    {"category": "sorting", "python_api": "sorted()", "go_api": "sort.Slice()", "description": "Sorts a slice using a custom less function."},
    {"category": "sorting", "python_api": "list.sort()", "go_api": "sort.Ints() / sort.Strings()", "description": "In-place sort for typed slices."},
    {"category": "sorting", "python_api": "sorted(key=func)", "go_api": "sort.Slice(s, func(i, j int) bool { ... })", "description": "Sort with custom comparison via closure."},
    {"category": "sorting", "python_api": "sorted(reverse=True)", "go_api": "sort.Sort(sort.Reverse(sort.IntSlice(s)))", "description": "Sorts in descending order."},
    {"category": "sorting", "python_api": "sorted(s, key=lambda x: x[1])", "go_api": "sort.Slice(s, func(i, j int) bool { return s[i][1] < s[j][1] })", "description": "Sort by specific field/index."},
    {"category": "sorting", "python_api": "bisect.insort()", "go_api": "sort.SearchInts() + slice insert", "description": "Insert into sorted slice maintaining order."},
    {"category": "sorting", "python_api": "bisect.bisect_left()", "go_api": "sort.SearchInts()", "description": "Binary search for insertion point."},
    {"category": "sorting", "python_api": "min()", "go_api": "slices.Min() / manual loop", "description": "Returns minimum element."},
    {"category": "sorting", "python_api": "max()", "go_api": "slices.Max() / manual loop", "description": "Returns maximum element."},

    # --- string_formatting ---
    {"category": "string_formatting", "python_api": "f'...'", "go_api": "fmt.Sprintf()", "description": "F-string to formatted string."},
    {"category": "string_formatting", "python_api": "str.format()", "go_api": "fmt.Sprintf()", "description": "String format method."},
    {"category": "string_formatting", "python_api": "'%d' % val", "go_api": "fmt.Sprintf(\"%d\", val)", "description": "Percent-style formatting."},
    {"category": "string_formatting", "python_api": "repr()", "go_api": "fmt.Sprintf(\"%#v\", v)", "description": "Printable representation of an object."},
    {"category": "string_formatting", "python_api": "str.zfill()", "go_api": "fmt.Sprintf(\"%0Nd\", val)", "description": "Zero-pads a numeric string."},
    {"category": "string_formatting", "python_api": "str.ljust()", "go_api": "fmt.Sprintf(\"%-Ns\", s)", "description": "Left-justify a string in a field."},
    {"category": "string_formatting", "python_api": "str.rjust()", "go_api": "fmt.Sprintf(\"%Ns\", s)", "description": "Right-justify a string in a field."},

    # --- data_structures ---
    {"category": "data_structures", "python_api": "list", "go_api": "[]T (slice)", "description": "Dynamic array / slice."},
    {"category": "data_structures", "python_api": "dict", "go_api": "map[K]V", "description": "Hash map."},
    {"category": "data_structures", "python_api": "set", "go_api": "map[T]struct{}", "description": "Set using map with empty struct values."},
    {"category": "data_structures", "python_api": "tuple", "go_api": "struct or multiple return values", "description": "Fixed-size heterogeneous collection."},
    {"category": "data_structures", "python_api": "list comprehension", "go_api": "for loop + append()", "description": "[expr for x in items] becomes a for loop with append."},
    {"category": "data_structures", "python_api": "dict comprehension", "go_api": "for loop + map assignment", "description": "{k: v for ...} becomes a for loop with m[k] = v."},
    {"category": "data_structures", "python_api": "enumerate()", "go_api": "for i, v := range slice", "description": "Iterate with index."},
    {"category": "data_structures", "python_api": "zip()", "go_api": "manual index loop", "description": "Pair elements from two sequences (manual in Go)."},
    {"category": "data_structures", "python_api": "range(n)", "go_api": "for i := 0; i < n; i++", "description": "Integer range loop."},
    {"category": "data_structures", "python_api": "range(a, b)", "go_api": "for i := a; i < b; i++", "description": "Integer range with start and stop."},
    {"category": "data_structures", "python_api": "range(a, b, step)", "go_api": "for i := a; i < b; i += step", "description": "Integer range with step."},
    {"category": "data_structures", "python_api": "len()", "go_api": "len()", "description": "Returns the length of a slice, map, string, or channel."},
    {"category": "data_structures", "python_api": "in (membership)", "go_api": "_, ok := m[key]", "description": "Check key existence in map."},
    {"category": "data_structures", "python_api": "del dict[key]", "go_api": "delete(m, key)", "description": "Delete a key from a map."},
    {"category": "data_structures", "python_api": "dict.get(key, default)", "go_api": "v, ok := m[key]; if !ok { v = default }", "description": "Get with default value."},
    {"category": "data_structures", "python_api": "dict.keys()", "go_api": "for k := range m", "description": "Iterate over map keys."},
    {"category": "data_structures", "python_api": "dict.values()", "go_api": "for _, v := range m", "description": "Iterate over map values."},
    {"category": "data_structures", "python_api": "dict.items()", "go_api": "for k, v := range m", "description": "Iterate over map key-value pairs."},
    {"category": "data_structures", "python_api": "slice[::-1]", "go_api": "slices.Reverse() / manual loop", "description": "Reverse a slice."},
    {"category": "data_structures", "python_api": "slice[a:b]", "go_api": "s[a:b]", "description": "Sub-slice / slicing."},
    {"category": "data_structures", "python_api": "''.join(list)", "go_api": "strings.Join(s, \"\")", "description": "Join slice of strings."},

    # --- error_handling ---
    {"category": "error_handling", "python_api": "try/except", "go_api": "if err != nil { ... }", "description": "Error handling via explicit error checks."},
    {"category": "error_handling", "python_api": "raise ValueError('msg')", "go_api": "errors.New(\"msg\") / fmt.Errorf(\"msg\")", "description": "Create and return a new error."},
    {"category": "error_handling", "python_api": "raise Exception('msg')", "go_api": "fmt.Errorf(\"msg\") / panic(\"msg\")", "description": "Raise/create an error; panic for unrecoverable."},
    {"category": "error_handling", "python_api": "with open(f) as fp:", "go_api": "f, err := os.Open(f); defer f.Close()", "description": "Context manager pattern using defer."},
    {"category": "error_handling", "python_api": "finally:", "go_api": "defer func() { ... }()", "description": "Cleanup logic runs on function exit."},
    {"category": "error_handling", "python_api": "assert condition", "go_api": "if !condition { panic(\"assertion failed\") }", "description": "Debug assertion (panic in Go)."},
    {"category": "error_handling", "python_api": "except Exception as e:", "go_api": "if err != nil { // use err }", "description": "Capture error for inspection."},
    {"category": "error_handling", "python_api": "try/except/else", "go_api": "if err != nil { ... } else { ... }", "description": "Execute code only if no error occurred."},

    # --- functional ---
    {"category": "functional", "python_api": "map(func, iterable)", "go_api": "for loop with transformation", "description": "Apply function to each element (manual in Go)."},
    {"category": "functional", "python_api": "filter(func, iterable)", "go_api": "for loop with condition + append", "description": "Filter elements by predicate (manual in Go)."},
    {"category": "functional", "python_api": "reduce(func, iterable)", "go_api": "for loop with accumulator", "description": "Reduce sequence to single value (manual in Go)."},
    {"category": "functional", "python_api": "lambda x: expr", "go_api": "func(x T) T { return expr }", "description": "Anonymous function."},
    {"category": "functional", "python_api": "any(iterable)", "go_api": "for loop with early return true", "description": "Check if any element is true."},
    {"category": "functional", "python_api": "all(iterable)", "go_api": "for loop with early return false", "description": "Check if all elements are true."},
    {"category": "functional", "python_api": "sum(iterable)", "go_api": "for loop accumulator", "description": "Sum all elements (manual loop in Go)."},
    {"category": "functional", "python_api": "reversed(seq)", "go_api": "iterate in reverse: for i := len(s)-1; i >= 0; i--", "description": "Reverse iteration."},

    # --- itertools ---
    {"category": "itertools", "python_api": "itertools.product()", "go_api": "nested for loops", "description": "Cartesian product via nested loops."},
    {"category": "itertools", "python_api": "itertools.permutations()", "go_api": "recursive backtracking", "description": "Generate permutations via recursion."},
    {"category": "itertools", "python_api": "itertools.combinations()", "go_api": "recursive backtracking", "description": "Generate combinations via recursion."},
    {"category": "itertools", "python_api": "itertools.chain()", "go_api": "append(s1, s2...)", "description": "Concatenate multiple iterables into one."},
    {"category": "itertools", "python_api": "itertools.groupby()", "go_api": "manual loop with key tracking", "description": "Group consecutive elements by key."},
    {"category": "itertools", "python_api": "itertools.count()", "go_api": "for i := start; ; i++", "description": "Infinite counter."},
    {"category": "itertools", "python_api": "itertools.repeat()", "go_api": "for loop with fixed value", "description": "Repeat a value n times."},
    {"category": "itertools", "python_api": "itertools.accumulate()", "go_api": "for loop with running sum/accumulator", "description": "Running accumulation (prefix sums)."},
    {"category": "itertools", "python_api": "itertools.zip_longest()", "go_api": "manual loop with length check", "description": "Zip with padding for unequal lengths."},

    # --- concurrency (expand) ---
    {"category": "concurrency", "python_api": "threading.Thread()", "go_api": "go func() { ... }()", "description": "Launch a concurrent goroutine."},
    {"category": "concurrency", "python_api": "threading.Lock()", "go_api": "sync.Mutex{}", "description": "Mutual exclusion lock."},
    {"category": "concurrency", "python_api": "queue.Queue()", "go_api": "chan T", "description": "Thread-safe queue via channels."},
    {"category": "concurrency", "python_api": "asyncio.gather()", "go_api": "sync.WaitGroup + goroutines", "description": "Run multiple tasks concurrently and wait."},
    {"category": "concurrency", "python_api": "asyncio.sleep()", "go_api": "time.Sleep()", "description": "Pause execution for a duration."},
    {"category": "concurrency", "python_api": "concurrent.futures.ThreadPoolExecutor()", "go_api": "goroutine pool pattern", "description": "Worker pool for parallel tasks."},
]


def main():
    console.print("[bold]Generating additional API mappings...[/bold]")

    # Load existing mappings to check for duplicates
    existing_apis = set()
    if API_MAPPINGS_FILE.exists():
        with open(API_MAPPINGS_FILE, encoding="utf-8") as f:
            for line in f:
                line = line.strip()
                if line:
                    r = json.loads(line)
                    existing_apis.add(r["python_api"])

    console.print(f"  Existing mappings: {len(existing_apis)}")

    # Filter out duplicates
    new_entries = [m for m in NEW_MAPPINGS if m["python_api"] not in existing_apis]
    console.print(f"  New mappings to add: {len(new_entries)}")

    if not new_entries:
        console.print("  [yellow]No new mappings to add.[/yellow]")
        return

    # Append to file
    with open(API_MAPPINGS_FILE, "a", encoding="utf-8") as f:
        for entry in new_entries:
            f.write(json.dumps(entry, ensure_ascii=False) + "\n")

    total = len(existing_apis) + len(new_entries)
    console.print(f"  [green]Done![/green] Total mappings: {total}")


if __name__ == "__main__":
    main()
