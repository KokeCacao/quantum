x=10000
while [ $x -gt 0 ]; do
    /usr/local/go/bin/go run main.go util.go -q < program.txt >> out.txt;
    x=$(($x-1))
done

# dangerous, race condition
# parallel -j 1000 sh -c "/usr/local/go/bin/go run main.go util.go -q < program.txt >> out.txt" -- $(seq 1 1000)
