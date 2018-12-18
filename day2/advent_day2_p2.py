file = open("advent_day2_input.txt", "r")
f = file.readlines()

for line1 in f:
    for line2 in f:
        distance = sum([1 for x, y in zip(line1, line2) if x.lower() != y.lower()])
        if distance == 1:
            print "line1: " + line1 + " , line2: " + line2




#for line in f:
#    for char in line:
#        if line.count(char) == 2:
#            count2+=1
#            print "count2: " + str(count2)
#            break
#    for char in line:
#        if line.count(char) == 3:
#            count3+=1
#            print "count3: " + str(count3)
#            break
#checksum = count2 * count3
#print str(checksum)
