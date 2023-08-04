@test "Without atomic" {
    ./atomic -atomic=false &
    serverPID=$!

    ab -c 1000 -n 150000 localhost:8000/count
    count=$(curl -qo- localhost:8000/getCount)
    kill $serverPID

    [ $count -ne 150000 ]
}

@test "Witht atomic" {
    ./atomic &
    serverPID=$!

    ab -c 1000 -n 150000 localhost:8000/count
    count=$(curl -qo- localhost:8000/getCount)
    kill $serverPID

    [ $count -eq 150000 ]
}