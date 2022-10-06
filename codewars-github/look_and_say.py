from itertools import groupby

def look_and_say(data='1', maxlen=5):
    L = []
    for i in range(maxlen):
        data = "".join(str(len(list(g)))+str(n) for n, g in groupby(data))
        L.append(data)
    return L