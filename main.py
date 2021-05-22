
import random
import math


def randMoney(money, people):
    if people == 1:
        return money, 0
    
    min = 0.01
    max = (money / people * 2.0)
    nexMoney = random.random() * max
    if nexMoney <= min:
        nexMoney = min

    nexMoney = math.floor(nexMoney * 100)/100.0

    money = math.floor((money - nexMoney) * 100)/100.0
    if money <= 0:
        money = 0.01
        nexMoney -= 0.01

    return nexMoney, money

def randMoneyV2(money, people):
    if people == 1:
        return money, 0
    
    min = 0.01
    max = money - ((people-1) * min)

    nexMoney = random.random() * max
    if nexMoney <= min:
        nexMoney = min

    nexMoney = math.floor(nexMoney * 100)/100.0

    money = math.floor((money - nexMoney) * 100)/100.0
    if money <= 0:
        money = 0.01
        nexMoney -= 0.01

    return nexMoney, money


def show(n, remainMoney):
    array = []
    for i in range(n):
        # money, remainMoney = randMoney(remainMoney, n-i)
        money, remainMoney = randMoneyV2(remainMoney, n-i)
        array.append(money)
    return array
              

if __name__ == '__main__':
    for i in range(100):
        r = show(6, 100)
        r.sort()
        print(r)
        # print(sum(r))

            
