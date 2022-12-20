import itertools
import random
import re

class Valve:
    def __init__(self, name, flow):
        self.name = name
        self.flow = flow
        self.connections = []

    def __repr__(self):
        return f"Valve({self.name}, {self.low}, {self.high})"


valves = {}
MAX_MINUTE = 26
INITIAL_POSITION = "AA"
regex = re.compile("Valve ([A-Z]+) has flow rate=([0-9]+);.* valve[s]* (.+)")
with open("day16_input.txt", 'r') as f:
    data = [regex.match(line).groups() for line in f.readlines()]
    for line in data:

        connections = []
        for connection in line[2].split(", "):
            connections.append({
                "valve": connection,
                "weight": 1
            })

        valves[line[0]] = {
            "valve": line[0],
            "flow": int(line[1]),
            "connections": connections
        }

def calculateDistancesFrom(graph, initial):
    visited = []
    unvisited = set(graph.keys())
    distances = {initial: 0}

    while len(unvisited) > 0:
        current = min(unvisited, key=lambda x: distances[x] if (x in distances) else 9999)
        for connection in graph[current]["connections"]:
            if connection["valve"] not in distances:
                distances[connection["valve"]] = distances[current] + connection["weight"]
        visited.append(current)
        unvisited.remove(current)
    return distances

for valve in valves.values():
    distances = calculateDistancesFrom(valves, valve["valve"])
    valve["paths"] = {}
    for (node, distance) in distances.items():
        if node == valve["valve"]:
            continue
        if valves[node]["flow"] == 0:
            continue
        valve["paths"][node] = distance


def valueOfOpening(graph, state, node, minute):
    if (minute >= MAX_MINUTE):
        return 0
    if isOpen(state, node):
        return 0

    # 1 is removed here because it takes 1 minute to open.
    remaining = (MAX_MINUTE - minute) - 1
    return graph[node]["flow"] * remaining

def writeOpenToState(state, node):
    nState = state.copy()
    nState[node] = True
    return nState

def isOpen(state, node):
    return node in state and state[node]


class Person:
    def __init__(self, name, position):
        self.name = name
        self.moveHistory = {
            0: position
        }
        self.isMoving = False

    def getPosition(self, minute):
        possible = max([(minute,position) for minute, position in self.moveHistory.items() if minute <= minute], key=lambda mp: mp[0])
        return possible[1]

    def getMovingTo(self, minute):
        if minute in self.moveHistory:
            return self.moveHistory[minute]
        return self.moveHistory[max(self.moveHistory, key=self.moveHistory.get)]


    def setIsMoving(self, isMoving, minute):
        self.isMoving = isMoving
        self.isMovingUntil = minute

    def getIsMoving(self, minute):
        return self.isMoving and minute < self.isMovingUntil

    def setStartMoving(self, destination, startTime, arrivalTime):
        self.setIsMoving(True, arrivalTime)
        self.moveHistory[arrivalTime] = destination

    def setIsOpening(self, minute):
        pass

    def __repr__(self):
        return f"Person({self.name}, {self.moveHistory})"

    def copy(self):
        nPerson = Person(self.name, self.getPosition(0))
        nPerson.moveHistory = self.moveHistory.copy()
        nPerson.isMoving = self.isMoving
        return nPerson


def isOtherPlayerGoingTo(graph, otherPosition, move, minute):
    if otherPosition.getMovingTo(minute) == move:
        return True
    return False


class Path:
    def __init__(self):
        self.posA = INITIAL_POSITION
        self.posB = INITIAL_POSITION

        self.weightA = 0
        self.weightB  = 0
        self.toggle = True
        self.path = []

    def getWeight(self):
        return max(self.weightA, self.weightB)

    def addPath(self, path, weight):
        if self.toggle:
            if(self.weightA + weight > MAX_MINUTE):
                raise Exception("Too heavy")

            self.weightA += weight
            self.path.append((self.weightA, path))
            self.posA = path
        else:
            if(self.weightB + weight > MAX_MINUTE):
                raise Exception("Too heavy")
            self.weightB += weight
            self.path.append((self.weightB, path))
            self.posB = path
        self.toggle = not self.toggle

    def __repr__(self):
        return f"Path{self.path})[{self.getWeight()}-{self.getUniqueKey()}]"

    def getUniqueKey(self):
        return ":".join([s[1] for s in self.path])

    def copy(self):
        nPath = Path()
        nPath.weightA = self.weightA
        nPath.weightB = self.weightB
        nPath.toggle = self.toggle
        nPath.posA = self.posA
        nPath.posB = self.posB
        nPath.path = list(self.path)
        if nPath.path == []:
            raise Exception("Empty path")
        return nPath

    def getCurrent(self):
        if self.toggle:
            return self.posA
        else:
            return self.posB

    def getRemainingTime(self):
        if self.toggle:
            return MAX_MINUTE - self.weightA
        else:
            return MAX_MINUTE - self.weightB

    def hasNode(self, node):
        return node in [path[1] for path in self.path]
paths = []
# we are going to generate all possible paths and brute force it.
currentA = INITIAL_POSITION
currentB = INITIAL_POSITION

# start at initial
# add to some path obj, inc by 1
# visit all child nodes

# p = Path()
# p.addPath(INITIAL_POSITION,0)
# p.addPath(INITIAL_POSITION,0)
# paths.append(p)
# def populatePaths(graph, paths):
#     lastUniquePathCount = -1
#     uniquePaths = set()
#     while True:
#         print(len(paths), len(uniquePaths))
#         newPaths = []
#         if lastUniquePathCount == len(uniquePaths):
#             print(uniquePaths)
#             break
#         lastUniquePathCount = len(uniquePaths)
#         uniquePaths = set()
#         for path in paths:
#             node = path.getCurrent()
#             for (connection, distance) in graph[node]["paths"].items():
#                 if path.getRemainingTime() <= distance:
#                     if path.getUniqueKey() in uniquePaths:
#                         continue
#                     newPaths.append(path)
#                     uniquePaths.add(path.getUniqueKey())
#                     continue
#                 elif path.hasNode(connection):
#                     continue
#
#                 newPath = path.copy()
#                 newPath.addPath(connection, distance+1)
#                 if newPath.getUniqueKey() in uniquePaths:
#                     continue
#                 uniquePaths.add(newPath.getUniqueKey())
#                 newPaths.append(newPath)
#         paths = newPaths
#     return paths

# paths = populatePaths(valves, paths)
# print(len(paths))

def dedup(paths):
    unique = set()
    dedupped = []
    for path in paths:
        if path.getUniqueKey() in unique:
            continue

        unique.add(path.getUniqueKey())
        dedupped.append(path)

    print("dedup done reduced", len(dedupped),"original", len(paths))
    return dedupped


maxScore = 0
checkedPaths = 0

# This heuristics function is specific to my input,
# this program likely won't work for other inputs.
def heuristic(path):
    remaining = path.getRemainingTime()
    pathValue = evaluatePath(path)

    # at minute 5 we have 19 flow from UV
    if remaining <= 20:
        if pathValue < (19 * 19):
            return False

    if remaining < 16:
        if pathValue < 800:
            return False

    if remaining < 12:
        if pathValue < 1200:
            return False

    if remaining < 6:
        if pathValue < 1400:
            return False

    return True


def pp2(graph, inprogress, completed):
    global maxScore
    global checkedPaths

    # print("inprogress", inprogress)
    if inprogress == []:
        return completed

    nInprogress = []
    i = 0
    for path in inprogress:
        if(not heuristic(path)):
            # skip path because it's too weak for our heuristic
            continue

        i += 1
        if i % 1000 == 0:
            percent = (float(i) / len(inprogress)) * 100
            percentS = "%.2f" % percent
            print(path, f"{path.getUniqueKey()}")
            print(f"{percentS}% {i}/{len(inprogress)} {len(nInprogress)}/{len(inprogress)} {len(completed)}")
        didWrite = False
        for newPath, distance in valves[path.getCurrent()]["paths"].items():
            if path.hasNode(newPath):
                continue
            if path.getRemainingTime() < distance + 1:
                continue
            didWrite = True
            nPath = path.copy()
            nPath.addPath(newPath, distance+1)
            nInprogress.append(nPath)

        if not didWrite:
            path.toggle = not path.toggle
            for newPath, distance in valves[path.getCurrent()]["paths"].items():
                if path.hasNode(newPath):
                    continue
                if distance + 1 >= path.getRemainingTime():
                    continue
                didWrite = True
                nPath = path.copy()
                nPath.addPath(newPath, distance + 1)
                nInprogress.append(nPath)

        if not didWrite:
            s = evaluatePath(path)
            checkedPaths += 1
            if s > maxScore:
                print("new max score", s)
                maxScore = s
            if s == 0:
                print("zero score", path)
            print(checkedPaths, "evaluated path", s, path, "max", maxScore)

    new = pp2(graph, nInprogress, completed)
    # combine new and completed
    return completed + new


def evaluatePath(path):
    path = path.copy()
    flow = 0
    total = 0
    minute = 0
    while minute <= MAX_MINUTE:
        total += flow
        # print("minute", minute, "flow", flow, "total", total)
        if minute == 0:
            minute += 1
            continue
        for (pMin, node) in list(path.path):
            if pMin <= minute:
                flow += valves[node]["flow"]
                path.path.remove((pMin, node))

        minute += 1
    return total

mp = Path()
mp.addPath(INITIAL_POSITION,0)
mp.addPath(INITIAL_POSITION,0)
paths = pp2(valves, [mp], [])
print(len(paths))

# path = Path()
# path.addPath("AA", 0)
# path.addPath("AA", 0)
#
# # ele
# path.addPath("DD", 2)
# # me
# path.addPath("JJ", 3)
#
# path.addPath("HH", 5)
# path.addPath("BB", 4)
#
# path.addPath("EE", 4)
# path.addPath("CC", 2)
#
# print(path)
# print (evaluatePath(path))