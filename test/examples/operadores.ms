// Prueba de estructuras de control: if, while, for y operadores lógicos.
count = 0
maxCount = 5

while count < maxCount
    print "Iteración: " + count
    count = count + 1
end while

for i = 1 to 3
    if i % 2 == 0
        print "Número par: " + i
    else
        print "Número impar: " + i
    end if
end for

// Prueba de función y retorno.
function factorial(n)
    if n <= 1
        return 1
    end if
    return n * factorial(n - 1)
end function

result = factorial(5)
print "5! = " + result