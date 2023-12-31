package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"bufio"
	"strings"
	"io/ioutil"
)

type Arguments struct {
	port string
	rootDir string
	transfer bool
	ssl bool
	crt string
	key string
	uploadPassword string
	getPassword string
}

func args() Arguments {
	f := Arguments{}
	flag.StringVar(&f.port, "p", "9000", "Port to listen on")
	flag.StringVar(&f.rootDir, "t", ".", "Root directory to serve files from")
	flag.BoolVar(&f.transfer, "u", false, "Accept transfer files from clients")
	flag.BoolVar(&f.ssl, "s", false, "Use SSL")
	flag.StringVar(&f.crt, "c", "", "SSL certificate file (.crt)")
	flag.StringVar(&f.key, "k", "", "SSL key file (.key)")
	flag.StringVar(&f.uploadPassword, "upw", "", "Password for uploading files")
	flag.StringVar(&f.getPassword, "pw", "", "Password for downloading files")
	flag.Parse()
	return f
}

func listAllFiles(rootDir string) []string {
	var files []string
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		// remove rootDir from path if path starts with rootDir
		if len(path) >= len(rootDir) && path[:len(rootDir)] == rootDir {
			path = path[len(rootDir):]
		}
		files = append(files, path)
		return nil
	})
	return files
}



func main() {
	arguments := args()

	// if rootDir does not end with a slash, add one
	if arguments.rootDir[len(arguments.rootDir)-1] != '/' {
		arguments.rootDir += "/"
	}

	fmt.Printf("Serving files from: %s\n", arguments.rootDir)
	fmt.Printf("Listening on port: %s\n", arguments.port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Download the file on: /id/<id_of_the_file>\n")
		allFiles := listAllFiles(arguments.rootDir)
		for id, file := range allFiles {
			fmt.Fprintf(w, "%d\t%s\n", id, file)
		}
	})

	r.HandleFunc("/id/{id}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
	
		allFiles := listAllFiles(arguments.rootDir)
		
		idInt, _ := strconv.Atoi(id)

		if idInt < len(allFiles) {
			relativeFilePath := allFiles[idInt]
			absoluteFilePath := arguments.rootDir + relativeFilePath

			// read file bytes
			file, err := os.Open(absoluteFilePath)
			if err != nil {
				fmt.Printf("File could not be read: %s\t%s\n", err, absoluteFilePath)
				http.Error(w, "File could not be read", http.StatusNotFound)
			}

			// write file bytes to response
			fileReader := bufio.NewReader(file)
			fileReader.WriteTo(w)
			file.Close()

		} else {
			// return with a bad request
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}
	})

	r.HandleFunc("/upload/{filename}", func (w http.ResponseWriter, r *http.Request) {
		if !arguments.transfer {
			http.Error(w, "File transfer is not enabled on the server!", http.StatusForbidden)
			return
		}
		filename := mux.Vars(r)["filename"]
		httpBody := r.Body
		body, bodyErr := ioutil.ReadAll(httpBody)

		if bodyErr != nil {
			http.Error(w, "Could not read HTTP body", http.StatusBadRequest)
		}


		if filename == "" {
			http.Error(w, "Blank filename", http.StatusBadRequest)
			return
		}

		if len(body) == 0 {
			http.Error(w, "Blank HTTP body", http.StatusBadRequest)
			return
		}

		// remove null characters from filename
		newFilename := ""
		for _, c := range filename {
			if c != 0 {
				newFilename += string(c)
			}
		}
		filename = newFilename

		filename = strings.Replace(filename, "../", "", -1)

		filename = strings.Replace(filename, "/", "", -1)

		// arguments.rootDir ends with a slash
		absoluteFilePath := arguments.rootDir + filename

		// write to file
		file, err := os.Create(absoluteFilePath)

		if err != nil {
			fmt.Printf("File could not be created: %s\t%s\n", err, absoluteFilePath)
			http.Error(w, "File could not be created", http.StatusNotFound)
		}

		fmt.Fprintf(file, "%s", body)
		// body.WriteTo(file)
		file.Close()

		fmt.Fprintf(w, "File uploaded successfully: %s\n", filename)

	}).Methods("PUT")

	port := ":" + arguments.port

	http.ListenAndServe(port, r)
}
