#!/bin/bash
touch a
touch '!'
touch '\'
touch '"'
mkdir '`'
cp '!' '`'

#export MOVE_A="yes"

if [ "$MOVE_A" = "yes" ]; then
    mv a '`'
elif [ "$MOVE_A" = "no" ]; then
    rm a
else 
    echo "Variable is unset or has an unexpected value"
fi
