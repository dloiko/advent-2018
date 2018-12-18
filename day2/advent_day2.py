file = open("advent_day2_input.txt", "r")
f = file.readlines()
count2 = 0
count3 = 0
for line in f:
    for char in line:
        if line.count(char) == 2:
            count2+=1
            print "count2: " + str(count2)
            break
    for char in line:
        if line.count(char) == 3:
            count3+=1
            print "count3: " + str(count3)
            break
checksum = count2 * count3
print str(checksum)
