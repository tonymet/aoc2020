//go:build unix

package shared

// func MmapFile(filename string) ([]byte, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer file.Close()
// 	fi, err := file.Stat()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get file info: %w", err)
// 	}
// 	size := int(fi.Size())
// 	data, err := syscall.Mmap(
// 		int(file.Fd()),
// 		0, // offset (start from the beginning of the file)
// 		size,
// 		syscall.PROT_READ,
// 		syscall.MAP_SHARED,
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to mmap file: %w", err)
// 	}
// 	return data, nil
// }
