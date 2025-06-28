// +build ignore
#include <stdio.h>
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
