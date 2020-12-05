#!/bin/sh
app=test_git_action
go build -o $app
mkdir ./output
tar -cvf output/$app.tgz -C . $app