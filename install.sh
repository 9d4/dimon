#!/bin/bash
[[ "$UID" -eq 0 ]] || exec sudo "$0" "$@"

# constants
DIMONEXEC="/usr/bin/dimon"
DIMONSERVICE="./dimon.service"
BIN_DIR="/usr/bin/"
LIB_DIR="/usr/lib/dimon/"
VAR_DIR="/var/lib/dimon/"
SYSTEM_SERVICE_DIR="/etc/systemd/system"

# make lib dir
rm -rf $LIB_DIR && mkdir $LIB_DIR
rm -rf $VAR_DIR && mkdir $VAR_DIR

# copy local bin to $BIN_DIR
cp ./dimon $BIN_DIR

# change permission
chmod +x $DIMONEXEC

# copy service
cp -f $DIMONSERVICE $SYSTEM_SERVICE_DIR

# reload daemon
systemctl daemon-reload

# start service
systemctl start dimon.service