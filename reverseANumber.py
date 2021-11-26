def reverse_number(n):
    is_negative = False
    if n < 0: 
        is_negative = True
        n *= -1
    rn = 0
    while (n != 0):
        rn = (rn*10) + n%10
        n //= 10
    if is_negative: rn *= -1
    return rn