numbers = []

for i in range(1, 100):
    msg = ""
    if i % 3 == 0:
        msg += "Fizz"
    if i % 5 == 0:
        msg += "Buzz"
    if msg == "":
        msg = i
    numbers.append((msg, i))

print(numbers)