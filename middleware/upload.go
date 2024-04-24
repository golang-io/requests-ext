package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// ServeUpload serve upload handler.
// curl http://127.0.0.1:8080/_upload H 'Content-Type: multipart/form-data' -F '/abc/123=@xxx.txt' -F '456=@abc/xyz.txt'
// upload files is $perfix/abc/123/xxx.txt and $perfix/456/xyz.txt
func ServeUpload(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reader, err := r.MultipartReader()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		save := func(prefix, dir, file string, r io.Reader) error {
			if file == "" { // this is FormData
				data, err := io.ReadAll(r)
				fmt.Printf("FormData=[%s]\n", string(data))
				return err
			}
			paths := path.Join(prefix, dir)
			if err := os.MkdirAll(paths, 0755); err != nil {
				return err
			}
			dst, err := os.Create(path.Join(paths, file))
			if err != nil {
				return err
			}
			defer dst.Close()
			_, err = io.Copy(dst, r)
			return err
		}

		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			//fmt.Printf("FileName=[%s], FormName=[%s]\n", part.FileName(), part.FormName())
			if err := save(prefix, part.FormName(), part.FileName(), part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		_, _ = fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}
