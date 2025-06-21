// Prueba de función y retorno.
function factorial(n)
    if n <= 1
        return 1
    end if
    return n * factorial(n - 1)
end function

result = factorial(5)
print "factorial de 5 es = "
print result