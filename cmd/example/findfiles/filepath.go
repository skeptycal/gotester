package main

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const (
	modeDir     os.FileMode = 0755
	modeNormal  os.FileMode = 0644
	modePrivate os.FileMode = 0700
)

// NewFilePath creates a new FilePath object with the given
// filename and options.
//
// The FilePath object encapsulates common functionality into
// a single object and abstracts away some of the underlying
// file system operations.
//
// Most values are cached and the underlying file system
// implementation is maintained on a JIT basis as needed.
// Continue to use
//  defer FP.Close()
// as necessary to release system resources.
//
// The root directory is automatically flagged as read-only.
func NewFilePath(path string, isDir, isPrivate, isReadOnly bool) (FilePath, error) {
	a, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fp := &filePath{
		input: path,
		abs:   a,
		isDir: isDir,
	}
	fp.SetPrivate(isPrivate)
	fp.SetReadOnly(isReadOnly)

	return fp, nil
}

// FilePath encapsulates common functionality into
// a single object and abstracts away some of the underlying
// file system operations.
//
// Most values are cached and the underlying file system
// implementation is maintained on a JIT basis as needed.
// Continue to use
//  defer FP.Close()
// as necessary to release system resources.
//
// The root directory is automatically flagged as read-only.
type FilePath interface {

	// Abs returns the absolute path
	Abs() string
	// Base returns the base name of the file without the directory
	Base() string
	// Dir returns the directory path of the file
	Dir() string

	// Exists returns whether the file exists
	Exists() bool
	// IsDir returns whether the file is a directory
	IsDir() bool
	// IsPrivate returns whether the file is private
	IsPrivate() bool
	// IsReadOnly returns whether the file is read-only
	IsReadOnly() bool

	// Mkdir creates a new directory if it does not exist
	Mkdir() error
	// Touch creates a new file if it does not exist
	Touch() error
	// SetPrivate sets the file privacy flag
	SetPrivate(isPrivate bool)
	// SetReadOnly sets the read-only flag
	SetReadOnly(isReadOnly bool)

	// Implements the common *os.File interface
	io.ReadWriteCloser
}

type filePath struct {
	input      string
	abs        string
	base       string
	dir        string
	isDir      bool
	isReadOnly bool
	isPrivate  bool
	exists     bool
	rw         *os.File
}

// Abs returns an absolute representation of path. If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. The absolute path name for a given file is not guaranteed to be unique. Abs calls Clean on the result.
func (f *filePath) Abs() string {
	var err error

	if f.abs == "" {
		f.abs, err = filepath.Abs(f.base)
		if err != nil {
			log.Fatal(err)
		}
	}

	return f.abs
}

// Base returns the last element of path. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns ".". If the path consists entirely of separators, Base returns a single separator.
func (f *filePath) Base() string {
	if f.base == "" {
		f.base = Base(f.Abs())
	}
	return f.base
}

func (f *filePath) Dir() string {
	if f.dir == "" {
		f.dir = Dir(f.Abs())
	}
	return f.dir
}

func (f *filePath) IsDir() bool {
	return f.isDir
}

func (f *filePath) IsReadOnly() bool {
	return f.isReadOnly
}

func (f *filePath) IsPrivate() bool {
	return f.isPrivate
}

func (f *filePath) SetPrivate(isPrivate bool) {
	f.isPrivate = isPrivate
}

func (f *filePath) SetReadOnly(isReadOnly bool) {

	// just in case ...
	if f.Abs() == "/" || f.Abs() == "~" {
		f.isReadOnly = true
		return
	}
	f.isReadOnly = isReadOnly
}

func (f *filePath) Stat() (fs.FileInfo, error) {
	return f.File().Stat()
}

func (f *filePath) Size() int {
	fi, err := f.File().Stat()
	if err != nil {
		// is this really necessary?
		log.Fatal(err)
		return 0
	}
	return int(fi.Size())
}

// File returns the underlying file descriptor. If the file
// is not open, it is opened with mode os.O_RDWR
func (f *filePath) File() *os.File {
	if f.rw == nil {
		f.Open()
	}
	return f.rw
}

// Read reads up to len(b) bytes from the File. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.
func (f *filePath) Read(p []byte) (int, error) {
	return f.File().Read(p)
}

// Write writes len(b) bytes to the File. It returns the number of bytes written and an error, if any. Write returns a non-nil error when n != len(b).
func (f *filePath) Write(p []byte) (int, error) {
	return f.File().Write(p)
}

// Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an error. Close will return an error if it has already been called.
func (f *filePath) Close() error {
	if f.rw == nil {
		return os.ErrClosed
	}
	return f.rw.Close()
}

// Open opens the underlying file with filePath flag IsReadOnly. If the file does not exist, it is created with mode determined by IsPrivate and IsDir. If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.
func (f *filePath) Open() (err error) {
	if f.rw == nil {
		f.rw, err = os.OpenFile(f.Abs(), os.O_RDWR|os.O_CREATE, f.modePerm())
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *filePath) Chmod(mode os.FileMode) error {
	return f.File().Chmod(mode)
}

func (f *filePath) Touch() error {
	if !f.Exists() {
		fi, err := os.Create(f.Abs())
		if err != nil {
			return err
		}
		f.rw = fi
		return nil
	}
	return os.ErrExist
}

func (f *filePath) Mkdir() error {
	return os.MkdirAll(f.Abs(), f.modePerm())
}

func (f *filePath) Exists() bool {
	_, err := os.Stat(f.Abs())
	return err == nil
}

func (f *filePath) modePerm() os.FileMode {
	if f.isDir {
		return modeDir
	}
	if f.isPrivate {
		return modePrivate
	}
	return modeNormal
}
