# Remote File Monitor

A simple web service written in Go to monitor a remote file.

It defaults to running on port 8090, but you can change this by setting the `RFM_PORT` environment variable. For example: `RFM_PORT=9993 remote-file-monitor`

## /size

Reports the size in bytes of a remote file. Takes a URL-encoded query parameter "url" as its sole argument.
