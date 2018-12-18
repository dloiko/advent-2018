sum=0
while IFS='' read -r line || [[ -n "$line" ]]; do
  sum+=$line
done < "$1"

echo $sum
