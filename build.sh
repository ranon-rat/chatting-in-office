GOOS=linux go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o linux-CIO/main main.go 
upx  --best --ultra-brute  executables/linux-CIO
echo "linux executable complete"
GOOS=darwin go build  -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o macos-CIO/main main.go 
upx  --best --ultra-brute  executables/macOS-CIO
echo "darwin executable complete"
GOOS=windows go build -a -gcflags=all="-l -B -wb=false" -ldflags="-w -s" -o windows-CIO/main main.go 
upx  --best --ultra-brute  executables/windows-CIO
echo "windows executable complete"
