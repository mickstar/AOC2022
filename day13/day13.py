with open("day13_input.txt", 'r') as f:
# with open("day13_test_input.txt", 'r') as f:
    lines = f.readlines()

data = []
for line in lines:
    if (line == "\n"):
        continue
    data.append(eval(line))

pairs = []
for i in range(0, len(data), 2):
    pairs.append((data[i], data[i + 1]))
    print("added pair", pairs[-1])


def compare(a, b):
    if b == []:
        if a == []:
            return 0
        return 1

    if type(a) is list and type(b) is not list:
        b = [b]
    elif type(a) is not list and type(b) is list:
        a = [a]

    elif type(a) is not list and type(b) is not list:
        if a < b:
            return -1
        elif a > b:
            return 1
        else:
            return 0

    for i in range(len(a)):
        try:
            cmp = compare(a[i], b[i])
            if cmp == -1:
                return -1
            elif cmp == 1:
                return 1
        except IndexError:
            return 1


def compare2(a, b):
    if (type(a) == type(b) and type(a) is not list):
        if a < b:
            print("a < b",a,b)
            return True
        elif a == b:
            print("continue: a == b",a,b)
            return None
        else:
            print("a > b",a,b)
            return False

    if type(a) is not list:
        a = [a]
    elif type(b) is not list:
        b = [b]

    if a == b:
        return None

    for i in range(len(a)):
        try:
            res = compare2(a[i], b[i])
            if res == None:
                continue
            return res
        except IndexError:
            print("b ran out of elements", a, b)
            return False

    print("True, a ran out of elements", a, b)
    return True


sum = 0
for i, pair in enumerate(pairs):
    print("="*40)
    print("pair 1", pair[0])
    print("pair 2", pair[1])
    index = i + 1
    res = compare2(pair[0], pair[1])
    print(index, res, pair)
    if (res == True):
        sum += index

print("part 1", sum)

# bubble sort LOL
def sortUsingCompare(data):
    for i in range(len(data)):
        for j in range(len(data)):
            if compare2(data[i], data[j]) == True:
                data[i], data[j] = data[j], data[i]

data.append([6])
data.append([2])
sortUsingCompare(data)
mul = 1
for i,line in enumerate(data):
    print(line)
    if line == [2] or line == [6]:
        mul *= (i + 1)

print("part 2", mul)