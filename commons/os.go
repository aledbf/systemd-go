package commons

import (
	"io"
	"os"

	"github.com/deis/systemd/logger"
)

func Getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		logger.Log.Debugf("returning default value \"%s\" for key \"%s\"", dfault, name)
		value = dfault
	}
	return value
}

func CopyFile(dst, src string) (int64, error) {
	sf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sf.Close()
	df, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sf)
}
