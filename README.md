# simpleHTTPServer
A simple HTTP Server inspired by python's simpleHTTPServer but implemented in Go and compiled to a single executable + more features.

This project is under development but is already in a usable state!

The project was written in go1.21 and uses gorilla/mux under the hood.

# Features

This is the roadmap for the project, you can already use the features if they are checked:

- [x] Interface for cURL and wget
- [x] Choosing a web root directory in the CLI interface
- [x] Accepting upload requests with `curl -T <file>`
- [ ] Interface for browsers implemented in svelte (probably)
- [ ] Zip compressing for multiple files
- [ ] Tar compressing for multiple files (for machines that do not have zip installed)
- [ ] TLS support
- [ ] Basic authentication for downloading files
- [ ] Basic authentication for uploading files
- [ ] Server side GUI interface for easier use

# Install

## Releases

Check https://github.com/hacker-szabo/simpleHTTPServer/releases and download the .zip file for your operating system and architecture.

## Using go installer

go1.21 is recommended

```
go install github.com/hacker-szabo/simpleHTTPServer@latest
```

# Build

go1.21 is recommended

```
git clone https://github.com/hacker-szabo/simpleHTTPServer.git
cd simpleHTTPServer
go build simpleHTTPServer.go
```

# Usage


## Server side

Obtain the executable and optionally set it up in a PATH folder.

Just simply run it and it will serve the current directory on port 9000

```
simpleHTTPServer
```

You can change the port using `-p`:

```
simpleHTTPServer -p 8000
```

You can change the webroot with the `-t` parameter:

```
simpleHTTPServer -t ~/tmp/web
```

## Client side

So far only CLI is supported, use any HTTP client, like cURL:

```
curl http://server_address:9000
```

It will return a file list of downloadable files with `id` numbers like so:

```
Download the file on: /id/<id_of_the_file>
0	
1	index.html
2	output
3	picture.jpg
4	test.bin
```

To download the file you can use either cURL or wget with the URL `/id/<id_number_of_the_file>`:

```
wget http://server_address:9000/id/3 -O picture.jpg
```

In the future, browser support will be available with a nice UI.

## Upload a file to simpleHTTPServer

Navigate to the webroot and run the software:

```
simpleHTTPServer -u
```

Where the `-u` means that you enable the clients to send files to upload with the PUT method.

OR provide the webroot by the `-t` flag:

```
simpleHTTPServer -u -t <your_webroot>
```

On client side you need to send the file with the following command:

```
curl <yourserver>:<yourport>/upload/ -T <yourfile>
```

For example:

```
curl 127.0.0.1:9000/upload/ -T test.txt
```

You can also give a different filename and the server will write it into that file or overwrite the existing file:

```
curl 127.0.0.1:9000/upload/FILENAME_ON_SERVER.TXT -T test.txt
```