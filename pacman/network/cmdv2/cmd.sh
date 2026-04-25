#! /bin/bash
btcd --configfile ./pacman/network/btcd.conf 
bazel run //pacman/network/cmdv2