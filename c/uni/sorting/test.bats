data=$(echo "23 342 234 234 23 4 543 5 35 345 34 563 645 456 5 5 6 56 456 45645676 567 56 56 567 5675675 67567")
result="4 5 5 5 6 23 23 34 35 56 56 56 234 234 342 345 456 456 543 563 567 567 645 67567 5675675 45645676"

@test "Test bubble sort" {
    out=$(echo ${data} | ./sorting -b)
    out=$(echo $out)

    expected=$(echo "Sort with bubble sort algorithm. ${result}")
    [[ "$out" == "$expected" ]]
}

@test "Test insertion sort" {
    out=$(echo ${data} | ./sorting -i)
    out=$(echo $out)

    expected=$(echo "Sort with insertion sort algorithm. ${result}")
    [[ "$out" == "$expected" ]]
}
