#include "runtime.h"

int main() {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="Hello, World!"; tmp; }));
  return 0;
}
