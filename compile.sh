#/bin/sh

set -e

#Linux
go build -o ./bin/

# Windows
fyne-cross windows -arch=amd64 --app-id br.com.mugomes.micheckhash --icon icon/micheckhash.png

exit 0
