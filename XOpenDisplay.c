#include <X11/Xlib.h>

int HasXorg() {
  Display *dpy = XOpenDisplay(NULL);
  if(!dpy) {
    return 1;
  }

  XCloseDisplay(dpy);
  return 0;
}
