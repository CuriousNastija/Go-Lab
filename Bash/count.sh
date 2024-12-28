#!/bin/bash
count=$(find . | wc -l); printf "\t\vTotal files * 5: %d\v\n" $((count * 5))