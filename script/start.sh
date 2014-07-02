#!/bin/sh

rm ./web
go build github.com/zhaoweiguo/go-web-brief/web
./web -v=1 -log_dir="/tmp/go-web-brief/" -stderrthreshold=FATAL



