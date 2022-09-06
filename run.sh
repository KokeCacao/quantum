rm out.txt
x=10000
while [ $x -gt 0 ]; do
    /usr/local/go/bin/go build -ldflags "-s -w" && ./quantum -q < program.txt >> out.txt;
    x=$(($x-1))
done
