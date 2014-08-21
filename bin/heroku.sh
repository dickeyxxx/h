#!/bin/sh

if [ ! -f ~/.hk/hk ]; then
  HK_URL="https://s3.amazonaws.com/dickeyxxx_dev/heroku-client/releases/hk.gz"
  mkdir -p ~/.hk
  cd ~/.hk
  echo "Downloading hk..."
  if [ -z `which wget` ]; then
    curl -s $HK_URL | gunzip > ~/.hk/hk
  else
    wget -qO- $HK_URL | gunzip > ~/.hk/hk
  fi
  chmod +x ~/.hk/hk
fi

exec ~/.hk/hk "$@"
