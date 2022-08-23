#ifndef TYPES_H
#define TYPES_H

typedef unsigned long XID;

typedef XID Window;

// man 3 XScreenSaverQueryInfo
typedef struct XScreenSaverInfo_t {
  Window window;    /* screen saver window */
  int state;      /* ScreenSaver{Off,On,Disabled} */
  int kind;       /* ScreenSaver{Blanked,Internal,External} */
  unsigned long til_or_since;   /* milliseconds */
  unsigned long idle;     /* milliseconds */
  unsigned long eventMask;  /* events */
} XScreenSaverInfo_t;


#endif
