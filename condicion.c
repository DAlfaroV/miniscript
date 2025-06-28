#include "runtime.h"

int main() {
  Value x;
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=0; tmp; });
  if (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i == ({ Value tmp; tmp.type=VAL_INT; tmp.i=3; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="pregunta 1: verdadero"; tmp; }));
  }
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=3; tmp; });
  if (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i == ({ Value tmp; tmp.type=VAL_INT; tmp.i=3; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="pregunta 2: verdadero"; tmp; }));
  }
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=10; tmp; });
  if (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i < ({ Value tmp; tmp.type=VAL_INT; tmp.i=20; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="pregunta 3: verdadero"; tmp; }));
  }
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=15; tmp; });
  if (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i >= ({ Value tmp; tmp.type=VAL_INT; tmp.i=12; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="pregunta 4: verdadero"; tmp; }));
  }
  x = ({ Value tmp; tmp.type=VAL_INT; tmp.i=11; tmp; });
  if (({ Value tmp; tmp.type=VAL_BOOL; tmp.b=(x.i <= ({ Value tmp; tmp.type=VAL_INT; tmp.i=23; tmp; }).i); tmp; }).b) {
  value_print(({ Value tmp; tmp.type=VAL_STRING; tmp.s="pregunta 5: verdadero"; tmp; }));
  }
  return 0;
}
