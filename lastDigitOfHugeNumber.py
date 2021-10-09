def last_digit(lst):
    if lst == []:
        return 1
    digit = 1
    for num in lst[::-1]:
        d = digit
        if d > 4:
            d = d%4+4
        digit = num ** d
    return digit % 10