package main

import (
	"errors"

	"io/ioutil"
	"strings"
)

// A new definition of File (in contrast of what is defined in os.File)
// FIXME: Change to a better type name because of type occlusion (?).
type File struct {
	Filename string
	Content  []byte
}

// Loads all the files under the directory <directory> with the extension ext
func Load(directory string, ext string) ([]File, error) {
	f, err := ioutil.ReadDir(directory)

	if err != nil {
		return nil, err
	}

	i, j := 0, 0
	files := make([]File, 0)
	extensions := strings.Split(ext, ",")

	for i < len(f) {
		if !f[i].IsDir() {
			var ok bool = false

			for k := range extensions {
				if strings.HasPrefix(extensions[k], ".") == false {
					err := errors.New("[!] Invalid extension " + extensions[k])
					return nil, err
				}

				ok = ok || strings.HasSuffix(f[i].Name(), extensions[k])
			}

			if ok {
				var file File
				file.Filename = f[i].Name()
				file.Content, err = ioutil.ReadFile(directory + f[i].Name())

				if err != nil {
					return nil, err
				}

				files = append(files, file)
			}
			j++
		}
		i++
	}

	return files, nil
}
