# Let's start with Installing GO in linux system

I use Linux as my ultimate coding platform. And I think hopefully all the folks watching this also use Linux as a their workspace.  

Lets get started with Go then, shall we .

## Step 1 – Install Go on Ubuntu

--------------
Login to your Ubuntu system and upgrade to apply latest security updates there.

``` shell
sudo apt-get update
sudo apt-get -y upgrade
```

Now download the Go language binary archive file using following link. To find and download latest version available or 32 bit version go to official [download page.](https://golang.org/dl/)  

``` shell
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
```

Now go to directory where you downloaded Go and extract it and install it into the desired location. Here I am install packages in `/usr/local` directory. You can put it into any other location you want.

``` shell
sudo tar -xvf go1.12.6.linux-amd64.tar.gz
sudo mv go /usr/local
```

## Step 2 – Setup Go Environment

--------------
Now you need to setup Go language environment variables for your project. Commonly you need to set 3 environment variables as **GOROOT**, **GOPATH** and **PATH**.  
  
**GOROOT** is the location where Go package is installed on your system.

```bash
$ export GOROOT=/usr/local/go
```

**GOPATH** is the location of your work directory. For example my project directory is `~/Projects/Proj1` .

```shell
$ export GOPATH=$HOME/Projects/Proj1
```

Now set the **PATH** variable to access go binary system wide.

```shell
$ export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

All the above environment will be set for your current session only. To make it permanent add above commands in `~/.profile` file or `~/.bashrc file`.  
  
## Step 3 – Verify Installation

--------------

First, use the following command to check the Go version.
  
```bash
$ go version

go version go1.12.6 linux/amd64
```
Now also verify all configured environment variables using following command.
  
```bash
$ go env

GOARCH="amd64"
GOBIN=""
GOCACHE="/home/azazel/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="BA1660E71660A65B/Go"
GOPROXY=""
GORACE=""
GOROOT="/usr/local/go"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build498866625=/tmp/go-build -gno-record-gcc-switches"
```

> Congratulations you are good to go for next chapter.
