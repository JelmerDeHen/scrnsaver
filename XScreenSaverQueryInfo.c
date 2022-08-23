#include <X11/extensions/scrnsaver.h>
#include <X11/Xlib.h>
#include "types.h"

int GetXScreenSaverQueryInfo(struct XScreenSaverInfo_t *info) {

  Display *dpy = XOpenDisplay(NULL);

  if(!dpy) {
    return 1;
  }

  XScreenSaverInfo *xscreensaverinfo = XScreenSaverAllocInfo();
  int errno = XScreenSaverQueryInfo(dpy, DefaultRootWindow(dpy), xscreensaverinfo);
  if(errno == 0) {
    return 2;
  }

  info->window = xscreensaverinfo->window;
  info->state = xscreensaverinfo->state;
  info->kind = xscreensaverinfo->kind;
  info->til_or_since = xscreensaverinfo->til_or_since;
  info->idle = xscreensaverinfo->idle;
  info->eventMask = xscreensaverinfo->eventMask;

  XFree(xscreensaverinfo);
  XCloseDisplay(dpy);

  return 0;
}
