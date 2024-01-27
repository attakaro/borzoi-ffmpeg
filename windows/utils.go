package windows

import (
	"path/filepath"
)

func getOutputFilePath(filePath, suffix, format string) string {
	dir := filepath.Dir(filePath)
	base := filepath.Base(filePath)
	ext := filepath.Ext(filePath)

	if format != "" {
		newExt := "." + format
		newBase := base[:len(base)-len(ext)] + suffix + newExt
		newPath := filepath.Join(dir, newBase)
		return newPath
	}

	newBase := base[:len(base)-len(ext)] + suffix + ext
	newPath := filepath.Join(dir, newBase)

	return newPath
}

func getFileNameWithExt(filePath string) string {
	base := filepath.Base(filePath)
	return base
}
