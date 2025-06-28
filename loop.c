#include "runtime.h"

int main() {
  Value x;
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=0; tmp; });
  while (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i < ({ Value tmp; tmp.type=VAL_INT; tmp.i=3; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="x es:"; tmp; }));
  value_print(x);
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=x.i + ({ Value tmp; tmp.type=VAL_INT; tmp.i=1; tmp; }).i; tmp; });
  }
  return 0;
}
