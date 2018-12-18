file = open("advent.freq", "r")
f = file.readlines()
seen = []
sum = 0
while 1==1:
    for line in f:
        #print seen
        if sum in seen:
            print "seen again: " + str(sum)
            exit()
        else:
            seen = seen + [sum]
        #print line
        if "-" in line:
            sum-=int(line[1:])
        else:
            sum+=int(line[1:])
        #print sum
