// +build ignore
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "runtime.h"

void value_print(Value v) {
    switch (v.type) {
        case VAL_INT:
            printf("%d\n", v.i);
            break;
        case VAL_FLOAT:
            printf("%f\n", v.f);
            break;
        case VAL_STRING:
            printf("%s\n", v.s);
            break;
        case VAL_BOOL:
            printf("%s\n", v.b ? "true" : "false");
            break;
        case VAL_NIL:
            printf("nil\n");
            break;
    }
}

Value value_concat(Value a, Value b) {
    if (a.type != VAL_STRING || b.type != VAL_STRING) {
        fprintf(stderr, "Error: concatenación solo soportada con cadenas.\n");
        exit(1);
    }
    size_t total = strlen(a.s) + strlen(b.s) + 1;
    char* result = malloc(total);
    if (!result) {
        fprintf(stderr, "Error: sin memoria para concatenación.\n");
        exit(1);
    }
    strcpy(result, a.s);
    strcat(result, b.s);
    return (Value){ .type=VAL_STRING, .s=result };
}

