# get version number as argument

VERSION=$1

# if no version number is provided, exit
if [ -z "$VERSION" ]
then
    echo "No version number provided. Exiting."
    exit 1
fi

# releases folder
RELEASES=./releases

# version folder
VERSION_FOLDER=$RELEASES/$VERSION

# create version folder
echo "Creating $VERSION_FOLDER"
mkdir -p $VERSION_FOLDER

# create darwin amd64
echo "Creating darwin_amd64"
GOOS=darwin GOARCH=amd64 go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/darwin_amd64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer

# create darwin arm64
echo "Creating darwin_arm64"
GOOS=darwin GOARCH=arm64 go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/darwin_arm64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer

# create linux amd64
echo "Creating linux_amd64"
GOOS=linux GOARCH=amd64 go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/linux_amd64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer

# create linux arm64
echo "Creating linux_arm64"
GOOS=linux GOARCH=arm64 go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/linux_arm64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer

# create linux 386
echo "Creating linux_386"
GOOS=linux GOARCH=386 go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/linux_386_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer

# create linux arm
echo "Creating linux arm"
GOOS=linux GOARCH=arm go build -o $VERSION_FOLDER/simpleHTTPServer
zip -j $VERSION_FOLDER/linux_arm_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer
rm $VERSION_FOLDER/simpleHTTPServer




# create windows amd64
echo "Creating windows_amd64"
GOOS=windows GOARCH=amd64 go build -o $VERSION_FOLDER/simpleHTTPServer.exe
zip -j $VERSION_FOLDER/windows_amd64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer.exe
rm $VERSION_FOLDER/simpleHTTPServer.exe

# create windows 386
echo "Creating windows_386"
GOOS=windows GOARCH=386 go build -o $VERSION_FOLDER/simpleHTTPServer.exe
zip -j $VERSION_FOLDER/windows_386_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer.exe
rm $VERSION_FOLDER/simpleHTTPServer.exe

# create windows arm
echo "Creating windows_arm"
GOOS=windows GOARCH=arm go build -o $VERSION_FOLDER/simpleHTTPServer.exe
zip -j $VERSION_FOLDER/windows_arm_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer.exe
rm $VERSION_FOLDER/simpleHTTPServer.exe

# create windows arm64
echo "Creating windows_arm64"
GOOS=windows GOARCH=arm64 go build -o $VERSION_FOLDER/simpleHTTPServer.exe
zip -j $VERSION_FOLDER/windows_arm64_$VERSION.zip $VERSION_FOLDER/simpleHTTPServer.exe
rm $VERSION_FOLDER/simpleHTTPServer.exe

echo "Done."