#include "runtime.h"

int main() {
  Value isValid;
  Value nothing;
  isValid = true;
  nothing = nil;
  if (isValid.b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="es valido"; tmp; }));
  value_print(nil);
  }
  return 0;
}
