package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0o755)
}

func WriteFileAtomic(path string, data []byte) error {
	if err := EnsureDir(filepath.Dir(path)); err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func SupplierCachePath(baseDir, supplierId string) string {
    return filepath.Join(baseDir, "suppliers", fmt.Sprintf("%s.toon", supplierId))
}


