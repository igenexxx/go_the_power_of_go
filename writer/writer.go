package writer

import "os"

func WriteToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0600)
}
