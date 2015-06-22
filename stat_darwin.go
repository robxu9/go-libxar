// struct stat on darwin

// +build darwin

package xar

// #include <stdlib.h>
// #include <sys/stat.h>
import "C"

import "golang.org/x/sys/unix"

func translateStatStruct(stat *unix.Stat_t) *C.struct_stat {
	return &C.struct_stat{
		st_dev:     C.dev_t(stat.Dev),
		st_ino:     C.__darwin_ino64_t(stat.Ino),
		st_mode:    C.mode_t(stat.Mode),
		st_nlink:   C.nlink_t(stat.Nlink),
		st_uid:     C.uid_t(stat.Uid),
		st_gid:     C.gid_t(stat.Gid),
		st_rdev:    C.dev_t(stat.Rdev),
		st_size:    C.off_t(stat.Size),
		st_blksize: C.blksize_t(stat.Blksize),
		st_blocks:  C.blkcnt_t(stat.Blocks),
		st_atimespec: C.struct_timespec{
			tv_sec:  C.__darwin_time_t(stat.Atimespec.Sec),
			tv_nsec: C.long(stat.Atimespec.Nsec),
		},
		st_mtimespec: C.struct_timespec{
			tv_sec:  C.__darwin_time_t(stat.Mtimespec.Sec),
			tv_nsec: C.long(stat.Mtimespec.Nsec),
		},
		st_ctimespec: C.struct_timespec{
			tv_sec:  C.__darwin_time_t(stat.Ctimespec.Sec),
			tv_nsec: C.long(stat.Ctimespec.Nsec),
		},
	}
}
