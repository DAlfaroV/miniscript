#include "runtime.h"

int main() {
  Value x;
  Value y;
  Value name;
  Value isValid;
  Value nothing;
  Value greeting;
  Value quote;
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=42; tmp; });
  y = ({ Value tmp; tmp.type=VAL_FLOAT; tmp.f=3.140000; tmp; });
  name = ({ Value tmp; tmp.type=VAL_STRING; tmp.s="MiniScript Tester"; tmp; });
  isValid = true;
  nothing = nil;
  greeting = ({ Value tmp; tmp.type=VAL_INT; tmp.i=({ Value tmp; tmp.type=VAL_STRING; tmp.s="Hi, "; tmp; }).i + name.i; tmp; });
  quote = ({ Value tmp; tmp.type=VAL_STRING; tmp.s="She said: "Keep coding!""; tmp; });
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="contenido asignado a las variables es:"; tmp; }));
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor x:"; tmp; }));
  value_print(x);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor y:"; tmp; }));
  value_print(y);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor name:"; tmp; }));
  value_print(name);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor greeting:"; tmp; }));
  value_print(greeting);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor quote:"; tmp; }));
  value_print(quote);
  return 0;
}
