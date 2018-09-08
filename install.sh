#!/bin/bash

export VERSION=v0.1.0

system=''
machine=''
case `uname` in
  Darwin) system='darwin' ;;
  Linux) system='linux' ;;
esac
case `uname -m` in
  x86_64) machine='amd64' ;;
  x86) machine='amd64' ;;
  amd64) machine='amd64' ;;
  i386) machine='386' ;;
esac
TARGET="${system}_${machine}"

EXECUTABLE_URL="https://github.com/coolopsio/coolops/releases/download/${VERSION}/coolops_${TARGET}"

echo "Downloading executable file ${EXECUTABLE_URL}..."
curl --fail -s -o coolops -L $EXECUTABLE_URL

echo "Installing the executable..."
mkdir -p /usr/local/bin
install -m 0755 coolops /usr/local/bin/coolops

echo "Cleaning up..."
rm coolops

echo "Instalation successfuly completed"
