# Go-Webview-demo

Cross build with debug console:

```
$ GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build
```

Cross build without debug console:

**x86_64**

```
$ GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags "-H windowsgui"
```
**x86**
```
$ GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc    CXX=i686-w64-mingw32-g++ go build -ldflags "-H windowsgui"
```

Finally add dll files from dll/x64 folder.

![2021-04-01_21h24_15](https://user-images.githubusercontent.com/743622/113338478-d4ee4e00-9331-11eb-864c-9a306e0356e0.png)
