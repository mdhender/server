#!/bin/bash
############################################################################
# list any process running on the port we're expecting to use
# netstat -tulpn | grep 3000

############################################################################
# usage
[ -n "${1}" ] || {
  echo "usage: $0 ./your_executable your optional args"
  exit 2
}

############################################################################
osname="$( uname )"
root=$PWD
# unset ftp_proxy FTP_PROXY http_proxy HTTP_PROXY https_proxy HTTPS_PROXY no_proxy NO_PROXY rsync_proxy RSYNC_PROXY

############################################################################
#
while true; do
  sleep 1 # give file system changes a chance to flush
  echo "############################################################################"
  echo "#"
  PID=
  go build
  case "$?" in
    0) $@ &
       PID=$! ;;
    *) echo
       echo "error: build failed"
       echo ;;
  esac
  case "${osname}" in
    Darwin)
      echo "watch: watching . excluding .git waiting 1 second"
      fswatch -l 1 --directories --recursive --exclude .git --one-event . ;;
    Linux)
      echo "watch: watching . excluding .git waiting 1 second"
      inotifywait -r -e modify . @.git
      sleep 1 ;;
    *)
      echo "error: unknown os '${oname}'"
      exit 2 ;;
  esac
  [ -n "${PID}" ] && {
    kill $PID
  }
  echo "#"
done
