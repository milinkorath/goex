package main

import "fmt"

type FileMode uint32

// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file (not backed up)
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // Unix permission bits
)

func main() {

	fmt.Println("ModeDir", ModeDir)
	fmt.Println("ModeAppend", ModeAppend)
	fmt.Println("ModeExclusive", ModeExclusive)
	fmt.Println("ModeTemporary", ModeTemporary)
	fmt.Println("ModeSymlink", ModeSymlink)
	fmt.Println("ModeDevice", ModeDevice)
	fmt.Println("ModeNamedPipe", ModeNamedPipe)
	fmt.Println("ModeSocket", ModeSocket)
	fmt.Println("ModeSetuid", ModeSetuid)
	fmt.Println("ModeSetgid", ModeSetgid)
	fmt.Println("ModeSticky", ModeSticky)
	fmt.Println("ModeCharDevice", ModeCharDevice)

}
