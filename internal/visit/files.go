package visit

import (
	"os"
	"path/filepath"
)

func MatchExtension(filename string, extensions []string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return ""
	}
	for _, extension := range extensions {
		if ext == extension {
			return ext
		}
	}
	return ""
}

type FileFunction func(filename string) error

// FilesInDirectory visits all files in the given directory by calling the visit function. Only
// those regular files that match an extension from extensions (don't forget to include the "." in your
// ".extension"). There is no recursion.
func FilesInDirectory(directory string, extensions []string, visit FileFunction) error {
	m := ExtensionsVisitMatcher{Extensions: extensions}
	return visitMatchedFiles(directory, false, m, visit)
}

// AllFilesInDirectory visits all files in the given directory by calling the visit function. Only
// regular files are visited. There is no recursion.
func AllFilesInDirectory(directory string, visit FileFunction) error {
	return visitMatchedFiles(directory, false, PromiscuousMatcher{}, visit)
}

type Matcher interface {
	ShouldVisit(path string) bool
}

type PromiscuousMatcher struct{}

func (m PromiscuousMatcher) ShouldVisit(path string) bool {
	return true
}

type ExtensionsVisitMatcher struct {
	Extensions []string
}

func (m ExtensionsVisitMatcher) ShouldVisit(path string) bool {
	return MatchExtension(path, m.Extensions) != ""
}

// FilesRecursively visits all files in the given directory and any sub-directories by calling
// the visit function. Only those regular files that match an extension from extensions (don't forget
// to include the "." in your ".extension").
func FilesRecursively(dirPath string, extensions []string, visit FileFunction) error {
	m := ExtensionsVisitMatcher{Extensions: extensions}
	return visitMatchedFiles(dirPath, true, m, visit)
}

// AllFilesRecursively visits all files in the given directory and any sub-directories by calling
// the visit function. Only regular files are visited.
func AllFilesRecursively(dirPath string, visit FileFunction) error {
	return visitMatchedFiles(dirPath, true, PromiscuousMatcher{}, visit)
}

func visitMatchedFiles(dirPath string, recurse bool, m Matcher, visit FileFunction) error {
	dirContents, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, f := range dirContents {
		info, err := f.Info()
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			filePath := filepath.Join(dirPath, f.Name())
			if m.ShouldVisit(filePath) {
				err := visit(filePath)
				if err != nil {
					return err
				}
			}
		} else if info.IsDir() && recurse {
			err = visitMatchedFiles(filepath.Join(dirPath, f.Name()), recurse, m, visit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DirectoryFunction func(directoryName string) error

// DirectoriesInDirectory visits all files in the given directory by calling the visit function. Only
// those regular files that match an extension from extensions (don't forget to include the "." in your
// ".extension"). There is no recursion.
func DirectoriesInDirectory(directory string, visit DirectoryFunction) error {
	dirContents, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, f := range dirContents {
		info, err := f.Info()
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = visit(filepath.Join(directory, f.Name()))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
