#include "day_1/day1.h"
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv) {
  if (argc < 3) {
    fprintf(stderr, "Usage: %s <day_number> <part_number> <input_file>\n",
            argv[0]);
    return 1;
  }

  int day = atoi(argv[1]);
  int part = atoi(argv[2]);
  const char *input_file = argv[3];

  switch (day) {
  case 1:
    if (part == 1) {
      part_1(input_file);
    }

    if (part == 2) {
      part_2(input_file);
    }
    break;
  default:
    fprintf(stderr, "Day %d not implemented.\n", day);
    return 1;
  }

  return 0;
}
