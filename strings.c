#include "runtime.h"

int main() {
  Value name;
  Value greeting;
  Value quote;
  name = ({ Value tmp; tmp.type=VAL_STRING; tmp.s="MiniScript Tester"; tmp; });
  greeting = ({ Value tmp; tmp.type=VAL_INT; tmp.i=({ Value tmp; tmp.type=VAL_STRING; tmp.s="Hi, "; tmp; }).i + name.i; tmp; });
  quote = ({ Value tmp; tmp.type=VAL_STRING; tmp.s="She said: "Keep coding!""; tmp; });
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="contenido asignado a las variables es:"; tmp; }));
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor name:"; tmp; }));
  value_print(name);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor greeting:"; tmp; }));
  value_print(greeting);
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="valor quote:"; tmp; }));
  value_print(quote);
  return 0;
}
