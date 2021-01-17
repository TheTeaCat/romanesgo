#!/bin/bash
mkdir "samples/multicorn/frames"
for i in $(seq 1 0.04 5)
do
    ./romanesgo -ff=multicorn -i=256 -z=0.5 -ss=4 -w=400 -h=400 -c=$i -fn="samples/multicorn/frames/${i}.png"
done