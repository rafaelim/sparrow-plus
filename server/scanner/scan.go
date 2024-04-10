package scanner

import (
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"sync"
)

type Scanner struct {
	DirsToScan []string
}

func NewScanner(dirsToScan []string) *Scanner {
	return &Scanner{
		DirsToScan: dirsToScan,
	}
}

func ScanDirectory(wg *sync.WaitGroup, syncChan chan<- string, rootDir string) {
	defer wg.Done()

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			syncChan <- path
		}
		return nil
	})

	if err != nil {
		slog.Error("Failed to scan directory", "dirPath", rootDir)
		return
	}
	slog.Info("Scan completed for the directory", "dirPath", rootDir)
}

func (s *Scanner) scan() []string {
	wg := &sync.WaitGroup{}
	syncChannel := make(chan string)

	for _, dir := range s.DirsToScan {
		wg.Add(1)
		go ScanDirectory(wg, syncChannel, dir)

	}

	go func() {
		wg.Wait()
		close(syncChannel)
	}()
	var videoPaths []string
	for path := range syncChannel {
		videoPaths = append(videoPaths, path)
	}
	return videoPaths
}

func (s *Scanner) Run() {
	videoPaths := s.scan()
	fmt.Println(videoPaths)
}
