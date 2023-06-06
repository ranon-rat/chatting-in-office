GOOS=darwin go build  -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o executables/macOS-CIO main.go 
upx  --best --ultra-brute  executables/macOS-CIO
echo "darwin executable complete"
GOOS=windows go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o executables/windows-CIO main.go 
upx  --best --ultra-brute  executables/windows-CIO

echo "windows executable complete"
GOOS=linux go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o executables/linux-CIO main.go 
upx  --best --ultra-brute  executables/linux-CIO

echo "linux executable complete"
