from sympy import symbols, solve, parse_expr
from sympy.parsing.mathematica import parse_mathematica

filename = "day21_input.txt"

def isNumber(s):
    try:
        int(s)
        return True
    except ValueError:
        return False

monkies = {}
data = []
with open(filename, 'r') as f:
    lines = [line.strip() for line in f.readlines()]
    for line in lines:
        a,b = line.split(": ")
        data.append(
            {
                "name": a,
                "value": int(b) if isNumber(b) else None,
                "operation": b if not isNumber(b) else None
            }
        )

    for d in data:
        monkies[d["name"]] = d
print(data)

# we trace every instruction and see what happens


evaluated = {}
def evaluateExpression(monkey):
    if monkey["value"] is not None:
        return monkey["value"]
    assert monkey["operation"] is not None

    parts = monkey["operation"].split(" ")
    assert len(parts) == 3
    if parts[0] not in evaluated:
        evaluated[parts[0]] = evaluateExpression(monkies[parts[0]])
    if parts[2] not in evaluated:
        evaluated[parts[2]] = evaluateExpression(monkies[parts[2]])
    exp = f"{evaluated[parts[0]]} {parts[1]} {evaluated[parts[2]]}"
    return eval(exp)

# need a symbolic number for humn
def eval2(monkey):
    if monkey["value"] is not None:
        return str ( monkey["value"] )

    if monkey["name"] == "humn":
        return "humn"

    assert monkey["operation"] is not None

    parts = monkey["operation"].split(" ")
    assert len(parts) == 3
    p1 = eval2(monkies[parts[0]])
    p2 = eval2(monkies[parts[2]])

    return f"({p1} {parts[1]} {p2})"

def part1():
    print("part 1", evaluateExpression(monkies["root"]))

def part2():
    monkies["root"] = {
        "name": "root",
        "value": None,
        "operation": monkies["root"]["operation"].replace("+", "==")
    }

    # humn is me
    monkies["humn"] = {
        "name": "humn",
        # "value": 301,
        "value": None,
        "operation": None
    }
    # print (evaluateExpression(monkies["root"]))
    print (eval2(monkies["root"]))

    exp = eval2(monkies["root"])
    exp = exp.replace("humn", "x")
    print("exp", exp)
    expr = parse_mathematica(exp)
    print("expr", expr)
    sol = solve(expr)
    print("part 2", sol)


part1()
part2()