#!/bin/sh
app=test_git_action
go build -o $app
tar -cvf output/$app.tgz -C $app