s = input()
a = []
for i in s:
    a.append(i)
a.sort(reverse=True)


tmp_1, tmp_2 = "", ""
for i in range(0, len(a) - 1, 2):
    x, y = a[i], a[i + 1]
    if tmp_1 > tmp_2:
        if x > y:
            x, y = y, x
    elif tmp_1 < tmp_2:
        if x < y:
            x, y = y, x
    tmp_1 += x
    tmp_2 += y

if len(a) % 2 != 0:
    if int(tmp_1) > int(tmp_2):
        tmp_2 += a[len(a) - 1]
    else:
        tmp_1 += a[len(a) - 1]
print(int(tmp_1) * int(tmp_2))
