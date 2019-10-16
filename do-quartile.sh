#!/bin/sh

# params : liste, datafield, parameters to quartile

if [ $# -lt 2 ]
then
 echo "Usage: $0 datafile field [quartile parameters]"
 exit 1
fi

file=$1
shift
datafield=$1
shift

awk -v field=$datafield '{ print $field }' $file | quartile $* > /tmp/qR
rg -f /tmp/qR $file
