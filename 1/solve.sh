#!/usr/local/bin/bash

lines=$(wc -l < input)

left_array=()
right_array=()
left_array_sorted=()
right_array_sorted=()
while read -r line; do
  left_array+=("$(echo "$line" | tr -s ' ' | cut -d ' ' -f 1)")
  right_array+=("$(echo "$line" | tr -s ' ' | cut -d ' ' -f 2)")
done < input

IFS=$'\n'
while read -r num; do
  left_array_sorted+=("$num")
done <<< "$(sort -n <<< "${left_array[*]}")"

while read -r num; do
  right_array_sorted+=("$num")
done <<< "$(sort -n <<< "${right_array[*]}")"
unset IFS

total_distance=0
part_2_sum=0
index=0
while [[ index -lt $lines ]]; do
  left_num=${left_array_sorted[$index]}
  right_num=${right_array_sorted[$index]}
  if [[ $left_num -gt $right_num ]]; then
    distance=$(( left_num - right_num ))
  else
    distance=$(( right_num - left_num ))
  fi
  (( total_distance+=distance ))

  #### Part 2
  count=$(printf "%s\n" "${right_array_sorted[@]}" | grep -c "$left_num")
  (( part_2_sum+= left_num * count ))

  (( index+=1 ))
done

echo "Total distance: $total_distance"
echo "Part 2 sum: $part_2_sum"