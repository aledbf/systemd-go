package commons

import (
	"io"
	"os"
)

func Getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		debug("returning default value \"%s\" for key \"%s\"", dfault, name)
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
