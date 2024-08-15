package writer

import "os"

func WriteToFile(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0600); err != nil {
		return err
	}

	return os.Chmod(filename, 0600)
}
