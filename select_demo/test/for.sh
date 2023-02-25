rm -f for.txt
touch for.txt
for i in `seq 10000`
do
  go build -race main.go
  ./main >> for.txt
  rm -f main
done
