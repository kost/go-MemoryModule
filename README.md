# go-MemoryModule
Go binding example for MemoryModule

Building 64 bit version of Memory Module and Go. Tested with Ubuntu 14.04 LTS.


Install required packages:
```
sudo apt-get install cmake
sudo apt-get install golang-go-linux-386 golang-go-windows-386 golang-go-windows-amd64
sudo apt-get install g++-mingw-w64-x86-64 g++-mingw-w64-i686 gcc-mingw-w64-i686 gcc-mingw-w64-x86-64
```

Clone from this specific branch (until pull request is accepted):
```
cd /source
git clone -b fix-deprecated-wcsnicmp https://github.com/kost/MemoryModule.git
```

Build MemoryModule:
```
cd MemoryModule
mkdir build
cmake ..
make VERBOSE=99
```

Clone this example:

```
cd /source
git clone https://github.com/kost/go-MemoryModule.git
```

Edit loadl.go to include path of MemoryModule

```
cd go-MemoryModule
vim loadl.go
```

Build go loader:

```
export GOOS=windows; export GOARCH=amd64; export CGO_ENABLED=1; export CXX=x86_64-w64-mingw32-g++; export CC=x86_64-w64-mingw32-gcc
CGO_LDFLAGS="-g -lm" GOGCCFLAGS="-m64 -fmessage-length=0" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o loadl.exe -x -v loadl.go
```

Test it:

```
wine loadl.exe 
```

You should get something like:

```
Loading DLL
Can't open DLL file "my.dll"
```

Adjust the source to you needs and good luck!

Note that you should properly reimplement reintrepret_cast for mingw. 


