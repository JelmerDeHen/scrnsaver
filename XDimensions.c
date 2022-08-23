#include <X11/Xlib.h>
#include <err.h>
#include <stdio.h>
#include <stdlib.h>

int GetDisplayWidth() {
  int width, snum;

  Display *dpy = XOpenDisplay(NULL);
  if(!dpy) {
    return 0;
  }

  snum = DefaultScreen(dpy);
  width = DisplayWidth(dpy, snum);

  XCloseDisplay(dpy);

  return width;
}

int GetDisplayHeight() {
  int height, snum;

  Display *dpy = XOpenDisplay(NULL);
  if(!dpy) {
    return 0;
  }

  snum = DefaultScreen(dpy);
  height = DisplayHeight(dpy, snum);

  XCloseDisplay(dpy);

  return height;
}
