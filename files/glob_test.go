package files

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestGlobalVariables(t *testing.T) {
	t.Run("test global variables initialization", func(t *testing.T) {
		pattern := "**/**.go"

		initGlobalVariables(pattern)
		if tokenizedPattern[0] != "**" && tokenizedPattern[1] != "**.go" {
			log.Fatal("pattern **/**.go should result in an array with two elements [0] == ** and [1] == **.go")
		}
		if patternLen != 2 {
			log.Fatal("pattern **/**.go should result in an array with a len of 2")
		}
		if len(fp) != 0 {
			log.Fatal("fp array must be empty as he is filled in GlobFiles function")
		}
		if lookupPattern != ".go" {
			log.Fatal("as last pattern contains two stars, we take every files that contains .go extension")
		}
	})
}

func TestSimplePatterns(t *testing.T) {
	t.Run("*.go pattern", func(t *testing.T) {
		// glob_test.go glob.go keep_or_remove_file_test.go keep_or_remove_file.go watcher.go watcher_test.go
		pattern := "*.go"
		files, err := GlobFiles(pattern)
		fmt.Println("Files", files)
		if err != nil {
			log.Fatal(err)
		}
		if len(files) != 6 {
			t.Errorf("A .go file hasn't been globbed, check pattern")
		}
	})

	t.Run(".git pattern", func(t *testing.T) {
		// .gitignore
		pattern := ".git"
		files, err := GlobFiles(pattern)
		if err != nil {
			log.Fatal(err)
		}
		for i := range files {
			if !(strings.Contains(files[i], ".git")) {
				t.Errorf("As we ignore directories, it should only return one file : .gitignore")
			}
		}
	})
}

func TestDoubleStarPatterns(t *testing.T) {
	t.Run("3 level nested double star pattern", func(t *testing.T) {
		// _samples/a/a.go _samples/a/a2.go _samples/b/b.go
		// glob.go keep_or_remove_file_test.go keep_or_remove_file.go watcher.go files_test.go watcher_test.go
		pattern := "**/**/*.go"

		files, err := GlobFiles(pattern)
		if err != nil {
			log.Fatal(err)
		}

		if len(files) != 9 {
			t.Errorf("A .go file hasn't been globbed when pattern is 3 level nested")
		}
	})

	t.Run("5 level nested double star pattern", func(t *testing.T) {
		// _samples/a/a.go _samples/a/a2.go _samples/b/b.go
		// _samples/a/aa/aa.go _samples/a/aa/aa2.go _samples/b/ba/ba.go _samples/b/ba/ba2.go
		// _samples/a/aa/aaa/aaa.go _samples/a/aa/aaa/aaa2.go
		// glob.go keep_or_remove_file_test.go keep_or_remove_file.go watcher.go files_test.go watcher_test.go
		pattern := "**/**/**/**/*.go"

		files, err := GlobFiles(pattern)
		if err != nil {
			log.Fatal(err)
		}

		if len(files) != 15 {
			t.Errorf("A .go file hasn't been globbed when pattern is 5 level nested")
		}
	})
}

func TestMultiplePatternsWithWildcardPattern(t *testing.T) {
	t.Run("*.go and *.yml pattern", func(t *testing.T) {
		var filesCount int
		// _samples/b/b.yml
		// glob.go keep_or_remove_file_test.go keep_or_remove_file.go watcher.go files_test.go watcher_test.go
		expression := "*.go **/**/*.yml"

		patterns := strings.Split(expression, " ")
		for i := range patterns {
			files, err := GlobFiles(patterns[i])
			if err != nil {
				log.Fatal(err)
			}

			filesCount += len(files)
		}

		if filesCount != 7 {
			t.Errorf("A file matching *.go or *.yml hasn't been globbed")
		}
	})
}

func TestMultiplePatternsWithoutWildcardPattern(t *testing.T) {
	t.Run("*.go and go.yml pattern", func(t *testing.T) {
		var filesCount int
		// glob.go keep_or_remove_file_test.go keep_or_remove_file.go watcher.go files_test.go watcher_test.go
		expression := "*.go **/**/go.yml"

		patterns := strings.Split(expression, " ")
		for i := range patterns {
			files, err := GlobFiles(patterns[i])
			if err != nil {
				log.Fatal(err)
			}

			filesCount += len(files)
		}
		if filesCount != 6 {
			t.Errorf("A file matching *.go or *.yml hasn't been globbed")
		}
	})
}
