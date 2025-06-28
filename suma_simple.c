#include "runtime.h"

int main() {
  Value x;
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=({ Value tmp; tmp.type=VAL_INT; tmp.i=3; tmp; }).i + ({ Value tmp; tmp.type=VAL_INT; tmp.i=4; tmp; }).i; tmp; });
  value_print(x);
  return 0;
}
