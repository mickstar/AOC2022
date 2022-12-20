import itertools
import re

class Valve:
    def __init__(self, name, flow):
        self.name = name
        self.flow = flow
        self.connections = []

    def __repr__(self):
        return f"Valve({self.name}, {self.low}, {self.high})"


valves = {}
MAX_MINUTE = 30
INITIAL_POSITION = "AA"
regex = re.compile("Valve ([A-Z]+) has flow rate=([0-9]+);.* valve[s]* (.+)")
with open("test_input.txt", 'r') as f:
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

def calculateBestmove(graph, state, position, minute):
    if minute >= MAX_MINUTE:
        return "",0

    toOpen = valueOfOpening(graph, state, position, minute)
    if toOpen > 0:
        nState = writeOpenToState(state, position)
        toOpen += calculateBestmove(graph, nState, position, minute + 1)[1]

    moves = {
        position: toOpen
    }
    for (move, distance) in graph[position]["paths"].items():
        if isOpen(state, move):
            continue
        res = valueOfOpening(graph, state, move, minute + distance)
        nState = writeOpenToState(state, move)
        res += calculateBestmove(graph, nState, move, minute + distance + 1)[1]
        moves[move] = res

    best = max(moves, key=moves.get)
    return (best, moves[best])


def test():
    state = {'DD': True}
    m = calculateBestmove(valves, state, "DD", 2)
    assert m[0] == "BB"

flow = 0
total = 0
minute = 1
state = {}
current = INITIAL_POSITION
lMinute = minute
while minute < MAX_MINUTE:
    lMinute = minute
    print(state)
    print(f"Minute {minute} current={current} flow={flow} total={total}")
    (move, score) = calculateBestmove(valves, state, current, minute)
    print(f"Determined best move {move} score={score}")
    distance = 0
    if move == current:
        distance = 0
    else:
        distance = valves[current]["paths"][move]
        total += flow * distance
    if not isOpen(state, move):
        flow += valves[move]["flow"]
        state = writeOpenToState(state, move)
    minute += distance
    minute += 1
    total += flow
    current = move



print("Total", total, minute)