package xsftp

import (
	"path/filepath"
)

//const maxPacket = 1 << 15
//
//func NewSftpClient(h  model.TMachine) (*sftp.Client, error) {
//	conn, err := ssh.NewSshClient(h)
//	if err != nil {
//		return nil, err
//	}
//	return sftp.NewClient(conn, sftp.MaxPacket(maxPacket))
//}
func toUnixPath(path string) string {
	return filepath.Clean(path)
}
