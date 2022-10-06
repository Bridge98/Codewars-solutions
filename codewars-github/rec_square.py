def sqInRect(lng, wdth):
    
    # your code
    if wdth == lng: 
        return None
    
    res = []
    while wdth != lng:
        if wdth > lng:
            wdth -= lng
            res.append(lng)
        elif lng > wdth:
            lng -= wdth
            res.append(wdth)
    res.append(wdth)
    
    return res