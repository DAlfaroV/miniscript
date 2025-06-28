#ifndef RUNTIME_H
#define RUNTIME_H

typedef enum {
    VAL_INT,
    VAL_FLOAT,
    VAL_STRING,
    VAL_BOOL,
    VAL_NIL
} ValueType;

typedef struct {
    ValueType type;
    union {
        int i;
        double f;
        const char* s;
        int b;
    };
} Value;

void value_print(Value v);

#endif
