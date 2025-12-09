//go:build windows

package shared

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows" // Use windows package for constants and specific functions
)

// MmapFileData structure to hold the mapped memory and necessary handles for cleanup.
type MmapFileData struct {
	Data       []byte
	file       *os.File
	mmapHandle windows.Handle
}

// mmapOpenFile maps the entire contents of a file into memory on Windows.
// It returns the MMapFile struct for access and later cleanup.
func MmapFile(filePath string) (*MmapFileData, error) {
	// 1. Open the file
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	// 2. Get file size
	stat, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("error getting file stats: %w", err)
	}
	fileSize := stat.Size()
	if fileSize == 0 {
		f.Close()
		return nil, fmt.Errorf("cannot map empty file")
	}

	// Get the OS file handle (syscall.Handle is the underlying Windows HANDLE type)
	handle := syscall.Handle(f.Fd())

	// 3. Create the file mapping object (CreateFileMapping)
	// We map the entire file (high=0, low=fileSize).
	// MEM_COMMIT is part of PAGE_READWRITE
	mmapHandle, err := windows.CreateFileMapping(
		windows.Handle(handle),
		nil,
		windows.PAGE_READWRITE,
		uint32(fileSize>>32),
		uint32(fileSize&0xFFFFFFFF),
		nil,
	)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("CreateFileMapping failed: %w", err)
	}

	// 4. Map the view of the file mapping object (MapViewOfFile)
	// Access: FILE_MAP_WRITE allows both read and write
	addr, err := windows.MapViewOfFile(mmapHandle, windows.FILE_MAP_WRITE, 0, 0, uintptr(fileSize))
	if err != nil {
		windows.CloseHandle(mmapHandle)
		f.Close()
		return nil, fmt.Errorf("MapViewOfFile failed: %w", err)
	}

	// 5. Convert the memory address to a Go []byte slice
	// Use unsafe pointers to treat the mapped memory region as a Go slice.
	var data []byte
	sliceHeader := (*[3]uintptr)(unsafe.Pointer(&data))
	sliceHeader[0] = addr
	sliceHeader[1] = uintptr(fileSize)
	sliceHeader[2] = uintptr(fileSize)

	return &MmapFileData{
		Data:       data,
		file:       f,
		mmapHandle: mmapHandle,
	}, nil
}
