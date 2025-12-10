from pulp import *

with open("input.txt") as f:
    lines = f.readlines()
    summ = 0
    for line in lines:
        parts = line.strip().split(" ")
        seqs = [[int(x) for x in seq[1:-1].split(",")] for seq in parts[1:-1]]
        jolts = [int(x) for x in parts[-1][1:-1].split(",")]

        max_jolt = max(jolts)

        vars = [LpVariable(f"x{i}", 0, max_jolt, cat="Integer") for i, seq in enumerate(seqs)]
        prob = LpProblem("triv", LpMinimize)

        prob += lpSum(vars)

        for i, jolt in enumerate(jolts):
            v = list(filter(lambda x: i in x[1], enumerate(seqs)))
            exp = vars[v[0][0]]
            for j, va in v[1:]:
                exp += vars[j]
            prob += exp == jolt
        status = prob.solve(PULP_CBC_CMD(msg=0))
        values = list(map(value, vars))

        summ += sum(values)
    print(summ)
