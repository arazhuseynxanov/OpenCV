Installing
Install the GoCV package:

![gocvlogo](https://github.com/arazhuseynxanov/OpenCV/assets/103748582/0254b3f5-3004-409b-bbdd-695e9c56f7ce)


go get -u -d gocv.io/x/gocv
You can install OpenCV 4.7.0 using Homebrew.

If you already have an earlier version of OpenCV installed, you should update:

brew upgrade opencv
If this is your first time installing OpenCV 4.7.0:

brew install opencv
pkgconfig Installation
pkg-config is used to determine the correct flags for compiling and linking OpenCV. You can install it by using Homebrew:

brew install pkgconfig
How to build/run code
Once you have installed OpenCV, you should be able to build or run any of the command examples:

go run ./cmd/version/main.go
The version program should output the following:
gocv version: 0.34.0
opencv lib version: 4.8.0
Cache builds
If you are running a version of Go older than v1.10 and not modifying GoCV source, precompile the GoCV package to significantly decrease your build times:

go install gocv.io/x/gocv
Custom Environment
By default, pkg-config is used to determine the correct flags for compiling and linking OpenCV. This behavior can be disabled by supplying -tags customenv when building/running your application. When building with this tag you will need to supply the CGO environment variables yourself.

go run -tags customenv ./cmd/version/main.go
