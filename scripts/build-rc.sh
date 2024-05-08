set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

cd ..
go build -o ravencheck.exe bin/ravencheck/ravencheck.go